package main

import (
	"fmt"
)

func displayHelp() {
	fmt.Print(`ctllm - Code to LLM

Usage:
  ctllm init [--force]       Initialize the project
  ctllm sync                 Sync the configuration with the current project structure
  ctllm                      Process the project and export code files

Options:
  --help                     Show this help message
  --force                    Force re-initialization of the project
`)
}

func colorRed(format string, a ...interface{}) {
	fmt.Printf("\033[31m"+format+"\033[0m\n", a...)
}

func colorGreen(format string, a ...interface{}) {
	fmt.Printf("\033[32m"+format+"\033[0m\n", a...)
}

func colorYellow(format string, a ...interface{}) {
	fmt.Printf("\033[33m"+format+"\033[0m\n", a...)
}

func colorRedBold(format string, a ...interface{}) {
	fmt.Printf("\033[1;31m"+format+"\033[0m\n", a...)
}
