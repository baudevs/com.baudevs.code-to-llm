package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

const configFileName = "ctllm-config.yaml"
const ignorePatternsFileName = "ignore_patterns.yaml"

type Config struct {
	Root           string   `yaml:"root"`
	OutputDir      string   `yaml:"output_dir"`
	TokenLimit     int      `yaml:"token_limit"`
	ProjectType    string   `yaml:"project_type"`
	IgnorePatterns []string `yaml:"ignore_patterns"`
}

func isInitialized() bool {
	_, err := os.Stat(configFileName)
	return err == nil
}

func loadConfig() (Config, error) {
	var config Config
	data, err := os.ReadFile(configFileName)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func saveConfig(config Config) {
	data, err := yaml.Marshal(&config)
	if err != nil {
		colorRed("Error saving configuration: %v", err)
		return
	}
	err = os.WriteFile(configFileName, data, 0644)
	if err != nil {
		colorRed("Error writing configuration file: %v", err)
	}
}

func initializeProject(force bool) {
	if isInitialized() && !force {
		colorYellow("This project is already initialized.")
		return
	}

	if isInitialized() && force {
		// Ask for confirmation with danger colors
		colorRedBold("Warning: You are about to re-initialize the project. This will overwrite the existing configuration.")
		if !confirmActionDanger("Are you sure you want to proceed?") {
			colorYellow("Re-initialization aborted.")
			return
		}
	}

	config := Config{
		Root:       prompt("Enter the root directory of your project", "."),
		OutputDir:  prompt("Enter the output directory", "ctllm_output"),
		TokenLimit: promptInt("Enter the maximum token limit per file", 8000),
	}

	detectedType := detectProjectType(config.Root)
	config.ProjectType = confirmProjectType(detectedType)

	// Load ignore patterns from ignore_patterns.yaml
	patterns, err := loadIgnorePatterns(config.ProjectType)
	if err != nil {
		colorRed("Error loading ignore patterns: %v", err)
		return
	}
	config.IgnorePatterns = patterns

	// Ask the user if they want to edit the ignore patterns
	colorYellow("Default ignore patterns have been set based on the project type.")
	if confirmAction("Would you like to review and edit the ignore patterns?") {
		// Let the user edit the patterns as a comma-separated list
		patternsStr := prompt("Enter ignore patterns separated by commas", strings.Join(config.IgnorePatterns, ","))
		config.IgnorePatterns = strings.Split(patternsStr, ",")
		// Trim whitespace from each pattern
		for i, pattern := range config.IgnorePatterns {
			config.IgnorePatterns[i] = strings.TrimSpace(pattern)
		}
	}

	// **New Feature: Ask about adding the output directory to .gitignore**
	if confirmAction(fmt.Sprintf("Would you like to add the output directory '%s' to your .gitignore file?", config.OutputDir)) {
		err := addOutputDirToGitignore(config.Root, config.OutputDir)
		if err != nil {
			colorRed("Error updating .gitignore: %v", err)
		} else {
			colorGreen("Output directory '%s' added to .gitignore.", config.OutputDir)
		}
	} else {
		colorYellow("You have chosen to track the output directory with Git.")
	}

	saveConfig(config)
	colorGreen("ctllm project initialized successfully.")
}

func addOutputDirToGitignore(root string, outputDir string) error {
	gitignorePath := filepath.Join(root, ".gitignore")
	// Check if .gitignore exists
	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		// .gitignore does not exist, offer to create it
		colorYellow(".gitignore file does not exist in the project root.")
		if confirmAction("Would you like to create a new .gitignore file?") {
			// Use os.WriteFile instead of ioutil.WriteFile
			err := os.WriteFile(gitignorePath, []byte(outputDir+"\n"), 0644)
			if err != nil {
				return fmt.Errorf("failed to create .gitignore: %v", err)
			}
			colorGreen("Created new .gitignore file and added '%s' to it.", outputDir)
			return nil
		} else {
			colorYellow("Output directory not added to .gitignore.")
			return nil
		}
	}

	// Read existing .gitignore content using os.ReadFile
	contentBytes, err := os.ReadFile(gitignorePath)
	if err != nil {
		return fmt.Errorf("failed to read .gitignore: %v", err)
	}
	content := string(contentBytes)

	// Check if outputDir is already in .gitignore
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == outputDir {
			// Output directory already in .gitignore
			colorYellow("Output directory '%s' is already in .gitignore.", outputDir)
			return nil
		}
	}

	// Append outputDir to .gitignore
	f, err := os.OpenFile(gitignorePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open .gitignore for writing: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString(outputDir + "\n"); err != nil {
		return fmt.Errorf("failed to write to .gitignore: %v", err)
	}

	colorGreen("Added '%s' to .gitignore.", outputDir)
	return nil
}

func syncConfig() {
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

	// Re-detect project type and ask for confirmation
	detectedType := detectProjectType(config.Root)
	config.ProjectType = confirmProjectType(detectedType)

	// Update default ignore patterns based on the new project type
	config.IgnorePatterns = getDefaultIgnorePatterns(config.ProjectType)

	saveConfig(config)
	colorGreen("Configuration synced successfully.")
}

func loadIgnorePatterns(projectType string) ([]string, error) {
	// Check if ignore_patterns.yaml exists
	if _, err := os.Stat(ignorePatternsFileName); os.IsNotExist(err) {
		colorYellow("%s not found. Using default ignore patterns.", ignorePatternsFileName)
		return getDefaultIgnorePatterns(projectType), nil
	}

	// Read the ignore_patterns.yaml file using os.ReadFile
	data, err := os.ReadFile(ignorePatternsFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %v", ignorePatternsFileName, err)
	}

	// Define a map to hold the ignore patterns
	patternsMap := make(map[string][]string)

	// Unmarshal the YAML data
	err = yaml.Unmarshal(data, &patternsMap)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %v", ignorePatternsFileName, err)
	}

	// Get the common patterns
	commonPatterns, ok := patternsMap["common"]
	if !ok {
		commonPatterns = []string{}
	}

	// Get the project-specific patterns
	projectPatterns, ok := patternsMap[projectType]
	if !ok {
		colorYellow("No ignore patterns found for project type '%s'. Using common patterns only.", projectType)
		projectPatterns = []string{}
	}

	// Combine the common and project-specific patterns
	combinedPatterns := append(commonPatterns, projectPatterns...)

	return combinedPatterns, nil
}
