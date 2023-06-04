package things

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/crockeo/dotfiles/cmd/migrate-to-obsidian/util"
)

func Main(ctx *cli.Context) error {
	destFolder := ctx.Args().Get(0)
	if destFolder == "" {
		return fmt.Errorf("dest folder is required")
	}

	destFolder, err := filepath.Abs(destFolder)
	if err != nil {
		return err
	}

	thingsDBPath, err := getThingsDBPath()
	if err != nil {
		return err
	}

	conn, err := sql.Open("sqlite3", thingsDBPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	areas, err := GetAreas(conn)
	if err != nil {
		return err
	}

	tasks, err := GetTasks(conn)
	if err != nil {
		return err
	}
	isGroup := map[string]struct{}{}
	for _, task := range tasks {
		if task.Project != nil {
			isGroup[*task.Project] = struct{}{}
		}
		if task.Heading != nil {
			isGroup[*task.Heading] = struct{}{}
		}
	}

	for _, task := range tasks {
		if !task.IsActive() {
			continue
		}
		hierarchy := task.Hierarchy(areas, tasks)

		if _, ok := isGroup[task.Uuid]; ok {
			continue
		}

		targetPath := filepath.Join(destFolder, hierarchy.Path(), util.EscapePath(task.Title)+".md")
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return err
		}

		contents, err := task.Render()
		if err != nil {
			return err
		}
		if err := os.WriteFile(targetPath, []byte(contents), 0644); err != nil {
			return err
		}
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

	return filepath.Join(databaseDir, "Things Database.thingsdatabase", "main.sqlite"), nil
}
