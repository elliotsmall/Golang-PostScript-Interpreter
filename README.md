# Golang-PostScript-Interpreter
PostScript Interpreter written in Golang as part of my programming language design class.

# Requirements:
Go 1.18 or later

# Setup
To setup project, either clone repo if files aren't already on local machine, then navigate to the needed directory.

If a Go module has not been instantiated already run the following:
go mod init postscript

# Running the interpreter
To run in REPL mode, enter the following command: go run .
You will then be taken into an interactive REPL where you can run commands freely.
To quit the REPL, enter "quit"

To run a file, enter the following: go run . program.ps

To run with lexical scoping, add the --lexical tag when running the program. ( -l is also accepted )

# Run tests
To run all tests, enter 'go test'
To run specific tests, enter 'go test -run TestFor'
To run with verbose output, add '-v'

# Unimplemented commands
maxlength - Maps in Go don't expose their allocated capacity separately from their current length. Because of this, maxlength returns the current number of key-value pairs.

putinterval - Go strings are immutable. In PostScript the putinterval function does an in-place mutation of a PostScript string, which cannot be done on Go strings.

