// Tokenizer

package main

import (
	"strconv"
	"strings"
	"unicode"
)

type TokenType int

const (
	TokInt TokenType = iota
	TokFloat
	TokBool
	TokString
	TokName
	TokOperator
	TokProcStart
	TokProcEnd
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(input string) []Token {
	var tokens []Token
	i := 0
	for i < len(input) {
		ch := rune(input[i])

		//Skip whitespace
		if unicode.IsSpace(ch) {
			i++
			continue
		}

		//Skip comments
		if ch == '%' {
			for i < len(input) && input[i] != '\n' {
				i++
			}
			continue
		}

		//Procedure delimiters
		if ch == '{' {
			tokens = append(tokens, Token{Type: TokProcStart, Value: "{"})
			i++
			continue
		}
		if ch == '}' {
			tokens = append(tokens, Token{Type: TokProcEnd, Value: "}"})
			i++
			continue
		}

		//String literals e.g: (Hello World)
		if ch == '(' {
			i++ //Skip (
			var sb strings.Builder
			depth := 1
			for i < len(input) && depth > 0 {
				c := input[i]
				if c == '(' {
					depth++
				} else if c == ')' {
					depth--
					if depth == 0 {
						i++
						break
					}
				}
				sb.WriteByte(c)
				i++
			}
			tokens = append(tokens, Token{Type: TokString, Value: sb.String()})
			continue
		}

		//Name literal e.g: /foo
		if ch == '/' {
			i++ //Skip /
			start := i
			for i < len(input) && !unicode.IsSpace(rune(input[i])) && input[i] != '{' && input[i] != '}' {
				i++
			}
			tokens = append(tokens, Token{Type: TokName, Value: input[start:i]})
			continue
		}

		// Number or operator
		start := i
		for i < len(input) && !unicode.IsSpace(rune(input[i])) && input[i] != '{' && input[i] != '}' {
			i++
		}
		word := input[start:i]
		if word == "true" {
			tokens = append(tokens, Token{Type: TokBool, Value: "true"})
		} else if word == "false" {
			tokens = append(tokens, Token{Type: TokBool, Value: "false"})
			//If err is anything other than nil, the string is not a valid int. Either float or operator
		} else if _, err := strconv.Atoi(word); err == nil {
			tokens = append(tokens, Token{Type: TokInt, Value: word})
			//If err is anything other than nil, string is not a valid float. Must be an operator
		} else if _, err := strconv.ParseFloat(word, 64); err == nil {
			tokens = append(tokens, Token{Type: TokFloat, Value: word})
		} else {
			tokens = append(tokens, Token{Type: TokOperator, Value: word})
		}

	}

	return tokens
}
