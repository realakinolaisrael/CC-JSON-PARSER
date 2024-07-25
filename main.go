package main

import (
    "fmt"
    "os"
    "unicode"
)

// TokenType is an enumeration of the possible token types in the JSON input.
type TokenType int

// Define the token types: left brace, right brace, and invalid token.
const (
    TOKEN_LBRACE TokenType = iota // {
    TOKEN_RBRACE                  // }
    TOKEN_INVALID
)

// Token struct represents a token with a type and a value.
type Token struct {
    Type  TokenType
    Value string
}

// lexer function takes an input string and tokenizes it into a slice of Tokens.
func lexer(input string) []Token {
    var tokens []Token
    // Iterate over each character in the input string.
    for _, char := range input {
        switch char {
        case '{':
            // Add a TOKEN_LBRACE token for the '{' character.
            tokens = append(tokens, Token{Type: TOKEN_LBRACE, Value: string(char)})
        case '}':
            // Add a TOKEN_RBRACE token for the '}' character.
            tokens = append(tokens, Token{Type: TOKEN_RBRACE, Value: string(char)})
        default:
            // Add a TOKEN_INVALID token for any non-whitespace invalid character.
            if !unicode.IsSpace(char) {
                tokens = append(tokens, Token{Type: TOKEN_INVALID, Value: string(char)})
            }
        }
    }
    return tokens
}

// parser function checks if the sequence of tokens represents a valid JSON object.
func parser(tokens []Token) bool {
    // A valid JSON object must have exactly two tokens: TOKEN_LBRACE followed by TOKEN_RBRACE.
    if len(tokens) == 2 &&
        tokens[0].Type == TOKEN_LBRACE &&
        tokens[1].Type == TOKEN_RBRACE {
        return true
    }
    return false
}

func main() {
    // Sample input (valid JSON)
    input := "{}"
    // Uncomment to test with invalid JSON
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
