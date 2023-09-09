package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run file_manipulator.go reverse <inputfile> <outputfile>")
		return
	}

	command := os.Args[1]
	inputFile := os.Args[2]
	outputFile := os.Args[3]

	if command == "reverse" {
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading content:", err)
			return
		}

		reversedData := reverseString(string(content))
		err = os.WriteFile(outputFile, []byte(reversedData), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
