package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func detectProjectType(root string) string {
	type projectSignature struct {
		Type       string
		Indicators []string
	}

	signatures := []projectSignature{
		{"Next.js", []string{"package.json", "next.config.js"}},
		{"Node.js + Express", []string{"package.json", "app.js", "server.js"}},
		{"React with Remix", []string{"package.json", "remix.config.js"}},
		{"Pure JavaScript", []string{"index.js"}},
		{"Svelte", []string{"package.json", "svelte.config.js"}},
		{"Python Data Science", []string{"requirements.txt", "*.ipynb"}},
		{"Python with Flask", []string{"requirements.txt", "app.py"}},
		{"Go Library", []string{"go.mod", "*.go"}},
		{"Go Web Project", []string{"go.mod", "*.go", "main.go"}},
		{"Rust", []string{"Cargo.toml", "src/main.rs"}},
		// Add more signatures as needed
	}

	for _, sig := range signatures {
		match := true
		for _, indicator := range sig.Indicators {
			if strings.Contains(indicator, "*") {
				files, _ := filepath.Glob(filepath.Join(root, indicator))
				if len(files) == 0 {
					match = false
					break
				}
			} else {
				if _, err := os.Stat(filepath.Join(root, indicator)); os.IsNotExist(err) {
					match = false
					break
				}
			}
		}
		if match {
			return sig.Type
		}
	}

	return "Unknown"
}

func confirmProjectType(detectedType string) string {
	if detectedType != "Unknown" {
		colorGreen("Detected project type: %s", detectedType)
	} else {
		colorYellow("Could not automatically detect the project type.")
	}

	projectTypes := []string{
		"Next.js",
		"Node.js + Express",
		"React with Remix",
		"Pure JavaScript",
		"Svelte",
		"Python Data Science",
		"Python with Flask",
		"Go Library",
		"Go Web Project",
		"Rust",
		"Other",
	}

	defaultIndex := 0
	for i, pt := range projectTypes {
		if pt == detectedType {
			defaultIndex = i
			break
		}
	}

	selector := promptui.Select{
		Label:     "Select your project type",
		Items:     projectTypes,
		CursorPos: defaultIndex,
	}

	index, result, err := selector.Run()
	if err != nil {
		fmt.Println(err)
		return detectedType
	}

	if result == "Other" {
		return prompt("Please enter your project type", "")
	}

	return projectTypes[index]
}

func getDefaultIgnorePatterns(projectType string) []string {
	defaultPatterns := []string{
		".DS_Store",
		"LICENSE",
		"README.md",
		"CHANGELOG.md",
		"CONTRIBUTING.md",
		"CODE_OF_CONDUCT.md",
		"SECURITY.md",
		".github/",
		".vscode/",
		".idea/",
		".gitignore",
	}

	switch projectType {
	case "Next.js", "Node.js + Express", "React with Remix", "Pure JavaScript", "Svelte":
		return append(defaultPatterns, []string{
			"node_modules/",
			"dist/",
			"build/",
			".env",
			"npm-debug.log",
			"yarn-error.log",
		}...)
	case "Python Data Science", "Python with Flask":
		return append(defaultPatterns, []string{
			"venv/",
			"__pycache__/",
			"*.pyc",
			".env",
			".ipynb_checkpoints/",
		}...)
	case "Go Library", "Go Web Project":
		return append(defaultPatterns, []string{
			"vendor/",
			"bin/",
			"*.exe",
			"*.dll",
			"*.so",
			"*.dylib",
		}...)
	case "Rust":
		return append(defaultPatterns, []string{
			"target/",
			"*.rs.bk",
		}...)
	default:
		return defaultPatterns
	}
}
