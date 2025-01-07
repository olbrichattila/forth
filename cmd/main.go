// Package main is the entry point, running the script from file or prompt
package main

import (
	"bufio"
	"fmt"
	"forth/internal/forth"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		err := runFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	runConsole()
}

func runFile(filePath string) error {
	code, err := LoadFileToString(filePath)
	if err != nil {
		return err
	}

	return forth.Run(code)
}

func runConsole() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Run code (press Enter to quit): ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if input == "" {
			break
		}

		err := forth.Run(input)
		if err != nil {
			fmt.Println(err.Error());
		}
	}
}

func LoadFileToString(filePath string) (string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", filePath)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return string(content), nil
}