package token

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/pedrolzoliveira/pinhata/array"
)

type TokenType string

const (
	INTEGER_LITERAL   TokenType = "integer_literal"
	VARIABLE_NAME     TokenType = "variable_name"
	FUNCTION          TokenType = "function"
	OPEN_PARENTHESES  TokenType = "("
	CLOSE_PARENTHESES TokenType = ")"
	OPEN_BRACES       TokenType = "{"
	CLOSE_BRACES      TokenType = "}"
	RETURN            TokenType = "return"
)

type Token struct {
	Type  TokenType
	Value any
}

func generateToken(identifier string) (Token, error) {
	switch TokenType(identifier) {
	case FUNCTION, OPEN_BRACES, CLOSE_BRACES, OPEN_PARENTHESES, CLOSE_PARENTHESES, RETURN:
		return Token{Type: TokenType(identifier)}, nil
	default:
		if value, error := strconv.Atoi(identifier); error == nil {
			return Token{Type: INTEGER_LITERAL, Value: value}, nil
		} else if unicode.IsLetter(rune(identifier[0])) {
			return Token{Type: VARIABLE_NAME, Value: identifier}, nil
		} else {
			return Token{}, fmt.Errorf("unknown identifier: %s", identifier)
		}
	}
}

func Tokenize(code string) (array.Array[Token], error) {
	trimmedCode := strings.Trim(code, " ")
	identifiers := regexp.MustCompile(`\s+`).Split(trimmedCode, -1)
	tokens := make(array.Array[Token], 0)
	for _, identifier := range identifiers {
		token, error := generateToken(identifier)
		if error != nil {
			return nil, error
		}
		tokens.Add(token)
	}
	return tokens, nil
}
