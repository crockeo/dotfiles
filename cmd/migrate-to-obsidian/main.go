package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/crockeo/dotfiles/cmd/migrate-to-obsidian/notion"
	"github.com/crockeo/dotfiles/cmd/migrate-to-obsidian/org"
	"github.com/crockeo/dotfiles/cmd/migrate-to-obsidian/things"
)

func main() {
	app := &cli.App{
		Name:  "migrate-to-obsidian",
		Usage: "Migrate notes from other formats to obsidian",
		Commands: []*cli.Command{
			{
				Name:   "migrate-org",
				Usage:  "Migrate org-mode files to obsidian",
				Action: org.Main,
			},
			{
				Name:   "migrate-things",
				Usage:  "Migrate Things.app database to obsidian",
				Action: things.Main,
			},
			{
				Name:   "migrate-notion",
				Usage:  "Migrate Notion.so database to obsidian",
				Action: notion.Main,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
