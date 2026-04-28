// Entry point only, parse CLI flags, open input file or stdin, and hand off to interpreter.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lexical := false

	args := os.Args[1:]
	var filename string

	for _, arg := range args {
		if arg == "--lexical" || arg == "-l" {
			lexical = true
		} else {
			filename = arg
		}
	}

	interp := NewInterpreter()
	interp.SetLexical(lexical)

	if filename != "" {
		runFile(interp, filename)
	} else {
		runREPL(interp)
	}
}

func runFile(interp *Interpreter, filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	tokens := Tokenize(string(data))
	err = interp.Execute(tokens)
	if err != nil && err != ErrQuit {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func runREPL(interp *Interpreter) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("PostScript Interpreter")
	fmt.Println("Type 'quit' to exit, --lexical flag for lexical scoping")
	fmt.Print(">> ")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == " " {
			fmt.Print(">> ")
			continue
		}

		token := Tokenize(line)
		err := interp.Execute(token)
		if err != nil {
			if err.Error() == "quit" {
				fmt.Println("Exiting...")
				return
			}
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		fmt.Print(">> ")
	}
}
