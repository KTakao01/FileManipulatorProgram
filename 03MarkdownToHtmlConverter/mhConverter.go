package main

import (
	"fmt"
	"os"

	"github.com/russross/blackfriday/v2"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run <file_name.go> markdown <inputfile> <outputfile>")
	}

	inputFile := os.Args[2]
	outputFile := os.Args[3]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error reading content: %w", err)
	}

	htmlContent := blackfriday.Run(content, blackfriday.WithNoExtensions())

	err = os.WriteFile(outputFile, htmlContent, 0644)
	if err != nil {
		fmt.Println("error writing to file: %w", err)
	}
}
