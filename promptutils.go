package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func prompt(label, defaultVal string) string {
	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultVal,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		return defaultVal
	}
	return result
}

func promptInt(label string, defaultVal int) int {
	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		return err
	}

	prompt := promptui.Prompt{
		Label:    label,
		Default:  fmt.Sprintf("%d", defaultVal),
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		return defaultVal
	}
	intValue, err := strconv.Atoi(result)
	if err != nil {
		return defaultVal
	}
	return intValue
}

func confirmAction(label string) bool {
	prompt := promptui.Prompt{
		Label:     label + " (y/N)",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		return false
	}
	return strings.ToLower(result) == "y"
}

func confirmActionDanger(label string) bool {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . | red | bold }} ",
		Valid:   "{{ . | red | bold }} ",
		Invalid: "{{ . | red | bold }} ",
		Success: "{{ . | red | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     label + " (y/N)",
		Templates: templates,
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		return false
	}
	return strings.ToLower(result) == "y"
}
