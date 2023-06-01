package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

func match(re *regexp.Regexp, s string) (map[string]string, bool) {
	match := re.FindStringSubmatch(s)
	if match == nil {
		return nil, false
	}
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i == 0 {
			continue
		}
		result[name] = match[i]
	}
	return result, true
}

func run(args ...string) error {
	cmd := exec.Cmd{}
	cmd.Path = args[0]
	cmd.Args = args[1:]

	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"failed to execute %s: %w\nstdout: %s\nstderr: %s",
			cmd.Path,
			err,
			cmd.Stdout,
			cmd.Stderr,
		)
	}
	return nil
}

var ORG_ROAM_NODE_RE = regexp.MustCompile(`(\d+-)?(?P<name>.*)\.org`)

func migrateOrg(ctx *cli.Context) error {
	orgFolder := ctx.Args().Get(0)
	if orgFolder == "" {
		return fmt.Errorf("org folder is required")
	}

	destFolder := ctx.Args().Get(1)
	if destFolder == "" {
		return fmt.Errorf("dest folder is required")
	}

	var err error
	orgFolder, err = filepath.Abs(orgFolder)
	if err != nil {
		return err
	}
	destFolder, err = filepath.Abs(destFolder)
	if err != nil {
		return err
	}

	pandocPath, err := exec.LookPath("pandoc")
	if err != nil {
		return err
	}

	err = filepath.WalkDir(orgFolder, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) != ".org" {
			return nil
		}

		relPath, err := filepath.Rel(orgFolder, path)
		if err != nil {
			return err
		}

		targetDir := filepath.Join(destFolder, filepath.Dir(relPath))
		targetName := filepath.Base(relPath)
		match, ok := match(ORG_ROAM_NODE_RE, targetName)
		if ok {
			targetName = match["name"] + ".md"
		}
		targetName = strings.Replace(targetName, "_", " ", -1)
		targetPath := filepath.Join(targetDir, targetName)

		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return err
		}

		return run(pandocPath, "-s", path, "-o", targetPath)
	})
	if err != nil {
		return err
	}

	return nil
}

func getThingsDBPath() (string, error) {
	// https://culturedcode.com/things/support/articles/2982272/
	rootDir := os.Getenv("HOME")
	if rootDir == "" {
		return "", fmt.Errorf("HOME is not set")
	}

	rootDir = filepath.Join(rootDir, "Library", "Group Containers", "JLMPQHK86H.com.culturedcode.ThingsMac")
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return "", err
	}

	var databaseDir string
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "ThingsData-") {
			databaseDir = filepath.Join(rootDir, entry.Name())
			break
		}
	}

	if databaseDir == "" {
		return "", fmt.Errorf("Missing ThingsData-XXXXX directory")
	}

	return filepath.Join(databaseDir, "main.sqlite"), nil
}

func migrateThings(ctx *cli.Context) error {
	thingsDBPath, err := getThingsDBPath()
	if err != nil {
		return err
	}

	fmt.Println(thingsDBPath)

	return nil
}

func main() {
	app := &cli.App{
		Name:  "migrate-to-obsidian",
		Usage: "Migrate notes from other formats to obsidian",
		Commands: []*cli.Command{
			{
				Name:   "migrate-org",
				Usage:  "Migrate org-mode files to obsidian",
				Action: migrateOrg,
			},
			{
				Name:   "migrate-things",
				Usage:  "Migrate Things.app database to obsidian",
				Action: migrateThings,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
