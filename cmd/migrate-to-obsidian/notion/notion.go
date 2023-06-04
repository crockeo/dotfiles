package notion

import (
	"archive/zip"
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func Main(ctx *cli.Context) error {
	notionExportFile := ctx.Args().Get(0)
	if notionExportFile == "" {
		return fmt.Errorf("notion export file required (.zip Settings > Export content > Export all workspace content)")
	}

	destFolder := ctx.Args().Get(1)
	if destFolder == "" {
		return fmt.Errorf("dest folder is required")
	}

	var err error
	notionExportFile, err = filepath.Abs(notionExportFile)
	if err != nil {
		return err
	}
	destFolder, err = filepath.Abs(destFolder)
	if err != nil {
		return err
	}

	reader, err := zip.OpenReader(notionExportFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		fmt.Println(file.Name)
	}

	return nil
}
