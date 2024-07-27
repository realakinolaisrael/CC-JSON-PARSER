package main

import (
    "fmt"
    "os"
    "unicode"
)

// token types in the JSON input.
type TokenType int

const (
    TOKEN_LBRACE TokenType = iota 
    TOKEN_RBRACE                
    TOKEN_INVALID
)

// Token struct with a type and a value.
type Token struct {
    Type  TokenType
    Value string
}

// lexer function takes an input string and tokenizes it into a slice of Tokens.
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

// To check if the sequence of tokens represents a valid JSON object.
func parser(tokens []Token) bool {
    return len(tokens) == 2 &&
        tokens[0].Type == TOKEN_LBRACE &&
        tokens[1].Type == TOKEN_RBRACE
}

func main() {
    // Sample input (valid JSON)
    input := "{}"
    // input := "{invalid}"

    // Tokenize the input string.
    tokens := lexer(input)

    // Parse the tokens and determine if the JSON is valid.
    if parser(tokens) {
        fmt.Println("Valid JSON")
        os.Exit(0) // Exit with code 0 for valid JSON.
    } else {
        fmt.Println("Invalid JSON")
        os.Exit(1) // Exit with code 1 for invalid JSON.
    }
}
