package main

import (
	"bufio"
	"fmt"
	"golox"
	"os"
)

func main() {
	if len(os.Args[1:]) > 1 {
		fmt.Println("Usage golox [script]")
		os.Exit(64)
	} else if len(os.Args[1:]) == 1 {
		runFile(os.Args[1])
	} else {
		runPromt()
	}
}

func runFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error read file: %s\n", err.Error())
		os.Exit(66)
	}

	run(string(data))
	if golox.HadError {
		os.Exit(65)
	}
}

func runPromt() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if input.Scan() {
			run(input.Text())
			golox.HadError = false
		}
	}
}

func run(source string) {
	scanner := golox.NewScanner(source)
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token.String())
	}
}
