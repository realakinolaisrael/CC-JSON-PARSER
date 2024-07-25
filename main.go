package main

import (
    "fmt"
    "os"
    "unicode"
)

type TokenType int

const (
	TOKEN_LBRACE TokenType = iota
	TOKEN_RBRACE
	TOKEN_INVALID
)

type Token struct{
	Type TokenType
	Value string
}

func lexer(input string) []Token {
    var tokens []Token
    for _, char := range input {
        switch char {
        case '{':
            tokens = append(tokens, Token{Type: TOKEN_LBRACE, Value: string(char)})
        case '}':
            tokens = append(tokens, Token{Type: TOKEN_RBRACE, Value: string(char)})
        default:
            if !unicode.IsSpace(char) {
                tokens = append(tokens, Token{Type: TOKEN_INVALID, Value: string(char)})
            }
        }
    }
    return tokens
}

func parser(tokens []Token) bool {
    if len(tokens) == 2 &&
        tokens[0].Type == TOKEN_LBRACE &&
        tokens[1].Type == TOKEN_RBRACE {
        return true
    }
    return false
}

func main() {
    // Sample input
    input := "{}"
    // input := "{invalid}"

    tokens := lexer(input)

    if parser(tokens) {
        fmt.Println("Valid JSON")
        os.Exit(0)
    } else {
        fmt.Println("Invalid JSON")
        os.Exit(1)
    }
}