package notion

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/crockeo/dotfiles/cmd/migrate-to-obsidian/util"
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
		defer reader.Close()
		contents, err := io.ReadAll(reader)
		if err != nil {
			return err
		}

		if filepath.Ext(zipFile.Name) == ".md" {
			contents = processContents(contents)
		}

		path := filepath.Join(destFolder, stripPathUUID(zipFile.Name))
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.Write(contents); err != nil {
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

var UUID_RE = regexp.MustCompile(`(?P<filename>.+) [a-z0-9]{32}(?P<suffix>\..+)?`)

func stripPathUUID(path string) string {
	parts := strings.Split(path, "/")
	for i, part := range parts {
		groups, ok := util.Match(UUID_RE, part)
		if ok {
			parts[i] = fmt.Sprintf("%s%s", groups["filename"], groups["suffix"])
		}
	}
	return filepath.Join(parts[1:]...)
}

func processContents(rawContents []byte) []byte {
	contents := string(rawContents)
	lines := strings.Split(contents, "\n")

	skip := 2
	if len(lines) < skip {
		skip = len(lines)
	}
	lines = lines[skip:]

	// TODO: scan each line for links to local files
	// and then correct those paths
	// with stripPathUUID

	return []byte(strings.Join(lines, "\n"))
}
