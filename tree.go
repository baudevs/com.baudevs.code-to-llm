package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"strings"

	ignore "github.com/sabhiram/go-gitignore"
)

func generateTreeStructure(root string, matcher *ignore.GitIgnore, outputDir string) (string, error) {
	var builder strings.Builder
	err := walkDir(root, root, "", &builder, matcher, outputDir)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}

func walkDir(root, path, prefix string, builder *strings.Builder, matcher *ignore.GitIgnore, outputDir string) error {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for i, entry := range entries {
		relPath, err := filepath.Rel(root, filepath.Join(path, entry.Name()))
		if err != nil {
			return err
		}

		// Skip the output directory
		if relPath == "." || relPath == outputDir || strings.HasPrefix(relPath, outputDir+string(os.PathSeparator)) {
			continue
		}

		// Skip ignored files and directories
		if matcher.MatchesPath(relPath) {
			continue
		}

		connector := "├── "
		if i == len(entries)-1 {
			connector = "└── "
		}
		builder.WriteString(prefix + connector + entry.Name() + "\n")

		if entry.IsDir() {
			newPrefix := prefix + "│   "
			if i == len(entries)-1 {
				newPrefix = prefix + "    "
			}
			err = walkDir(root, filepath.Join(path, entry.Name()), newPrefix, builder, matcher, outputDir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
