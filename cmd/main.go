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
			os.Exit(1)
		}
		return
	}

	err := runConsole()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runFile(filePath string) error {
	code, err := loadFileToString(filePath)
	if err != nil {
		return fmt.Errorf("error running file '%s': %v", filePath, err)
	}

	return forth.Run(code)
}

func runConsole() error {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	for {
		fmt.Print("Enter code (or press Enter to exit): ")
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

	return nil
}

func loadFileToString(filePath string) (string, error) {
    content, err := os.ReadFile(filePath)
    if err != nil {
        return "", fmt.Errorf("failed to read file '%s': %v", filePath, err)
    }
    return string(content), nil
}
