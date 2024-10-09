package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	ignore "github.com/sabhiram/go-gitignore"
)

// getGitignoreMatcher reads the .gitignore file (if it exists) and combines its patterns
// with the default ignore patterns provided. It then compiles these patterns into a matcher.
func getGitignoreMatcher(root string, ignorePatterns []string) (*ignore.GitIgnore, error) {
	gitignorePath := filepath.Join(root, ".gitignore")

	// Read .gitignore patterns if the file exists
	if _, err := os.Stat(gitignorePath); err == nil {
		data, err := ioutil.ReadFile(gitignorePath)
		if err != nil {
			return nil, err
		}
		lines := strings.Split(string(data), "\n")
		// Remove any empty lines or comments
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			ignorePatterns = append(ignorePatterns, line)
		}
	}

	// Compile the combined ignore patterns
	matcher := ignore.CompileIgnoreLines(ignorePatterns...)

	return matcher, nil
}

// getFiles recursively walks through the directory starting from 'root',
// and collects all files that are not ignored by the 'matcher'.
func getFiles(root string, matcher *ignore.GitIgnore, outputDir string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the relative path from the root
		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		// Skip the output directory to prevent including generated files
		if relPath == "." || relPath == outputDir || strings.HasPrefix(relPath, outputDir+string(os.PathSeparator)) {
			return nil
		}

		// Skip ignored files and directories
		if matcher.MatchesPath(relPath) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Add the file to the list
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
