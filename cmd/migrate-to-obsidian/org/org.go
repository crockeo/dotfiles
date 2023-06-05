package org

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/crockeo/dotfiles/cmd/migrate-to-obsidian/util"
)

func Main(ctx *cli.Context) error {
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
		match, ok := util.Match(ORG_ROAM_NODE_RE, targetName)
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
