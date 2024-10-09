package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func splitFilesIntoChunks(files []string, root string, tokenLimit int, projectType string) ([]string, error) {
	var chunks []string
	var currentChunk []string
	var currentTokenCount int

	// Include project type in the first chunk
	projectTypeHeader := fmt.Sprintf("Project Type: %s\n\n", projectType)
	currentChunk = append(currentChunk, projectTypeHeader)
	currentTokenCount += estimateTokens(projectTypeHeader)

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		// Prepend the relative path as a header
		relPath, err := filepath.Rel(root, file)
		if err != nil {
			return nil, err
		}
		fileHeader := fmt.Sprintf("File: %s\n", relPath)
		fileContent := fileHeader + string(content)
		tokenCount := estimateTokens(fileContent)

		// Check if adding this file exceeds the token limit
		if currentTokenCount+tokenCount > tokenLimit && len(currentChunk) > 0 {
			chunks = append(chunks, strings.Join(currentChunk, "\n\n"))
			currentChunk = []string{}
			currentTokenCount = 0
		}

		currentChunk = append(currentChunk, fileContent)
		currentTokenCount += tokenCount
	}

	if len(currentChunk) > 0 {
		chunks = append(chunks, strings.Join(currentChunk, "\n\n"))
	}
	return chunks, nil
}

func estimateTokens(text string) int {
	// Rough estimate: 1 token per 4 characters
	return len(text) / 4
}
