package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func processProject(config Config) error {
	// Create output directory if it doesn't exist
	if _, err := os.Stat(config.OutputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(config.OutputDir, 0755); err != nil {
			return err
		}
	}

	// Get the ignore matcher, passing the ignore patterns from config
	matcher, err := getGitignoreMatcher(config.Root, config.IgnorePatterns)
	if err != nil {
		return err
	}

	// **Update Here:** Get all files to process, now passing config.OutputDir
	files, err := getFiles(config.Root, matcher, config.OutputDir)
	if err != nil {
		return err
	}

	// Generate and save the tree structure
	treeStr, err := generateTreeStructure(config.Root, matcher, config.OutputDir)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(config.OutputDir, "file_tree.txt"), []byte(treeStr), 0644)
	if err != nil {
		return err
	}

	// Split files into chunks based on the token limit
	chunks, err := splitFilesIntoChunks(files, config.Root, config.TokenLimit, config.ProjectType)
	if err != nil {
		return err
	}

	// Write each chunk to a separate file
	for i, chunk := range chunks {
		outputPath := filepath.Join(config.OutputDir, fmt.Sprintf("code_chunk_%d.txt", i+1))
		err = ioutil.WriteFile(outputPath, []byte(chunk), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
