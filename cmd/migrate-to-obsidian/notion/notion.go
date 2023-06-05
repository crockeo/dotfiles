package notion

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
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

	if err := traverseZipFile(reader, func(zipFile *zip.File) error {
		reader, err := zipFile.Open()
		if err != nil {
			return err
		}
		path := filepath.Join(destFolder, zipFile.Name)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		if _, err := io.Copy(file, reader); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func traverseZipFile(zipFile *zip.ReadCloser, fn func(*zip.File) error) error {
	tmpDir, err := os.MkdirTemp("", "notion-export")
	if err != nil {
		return err
	}
	defer func() {
		os.RemoveAll(tmpDir)
	}()

	for _, file := range zipFile.File {
		var err error
		if filepath.Ext(file.Name) == ".zip" {
			err = handleZipFile(file, tmpDir, fn)
		} else {
			err = fn(file)
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func handleZipFile(file *zip.File, tmpDir string, fn func(*zip.File) error) error {
	zipFile, err := file.Open()
	if err != nil {
		return nil
	}

	tmpPath := filepath.Join(tmpDir, file.Name)
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		return nil
	}

	if _, err := io.Copy(tmpFile, zipFile); err != nil {
		return err
	}

	subZipFile, err := zip.OpenReader(tmpPath)
	if err != nil {
		return err
	}

	return traverseZipFile(subZipFile, fn)
}
