package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to ctllm (Code to LLM)!")

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help", "help":
			displayHelp()
			return
		case "init":
			force := false
			// Check for --force flag
			if len(os.Args) > 2 && os.Args[2] == "--force" {
				force = true
			}
			initializeProject(force)
			return
		case "sync":
			syncConfig()
			return
		}
	}

	if !isInitialized() {
		colorRed("Error: This project is not initialized.")
		fmt.Println("Please run 'ctllm init' to initialize the project.")
		return
	}

	config, err := loadConfig()
	if err != nil {
		colorRed("Error loading configuration: %v", err)
		return
	}

	err = processProject(config)
	if err != nil {
		colorRed("Error processing project: %v", err)
	} else {
		colorGreen("Project files exported successfully!")
	}
}
