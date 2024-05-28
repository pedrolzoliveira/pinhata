package token

import (
	"regexp"

	"github.com/pedrolzoliveira/pinhata/array"
)

type TokenType string

const (
	FUNCTION          TokenType = "function"
	IDENTIFIER        TokenType = "identifier"
	OPEN_PARENTHESIS  TokenType = "open parenthesis"
	CLOSE_PARENTHESIS TokenType = "close parenthesis"
	OPEN_BRACE        TokenType = "open brace"
	CLOSE_BRACE       TokenType = "close brace"
	COMMA             TokenType = "comma"
	RETURN            TokenType = "return"
	SEMI_COLUMN       TokenType = "semi column"
	PLUS_SIGN         TokenType = "plus sign"
)

var tokenToRegex = map[TokenType]regexp.Regexp{
	FUNCTION:          *regexp.MustCompile("\\bfunction\\b"),
	RETURN:            *regexp.MustCompile("\\breturn\\b"),
	IDENTIFIER:        *regexp.MustCompile("\\w+"),
	OPEN_PARENTHESIS:  *regexp.MustCompile("\\("),
	CLOSE_PARENTHESIS: *regexp.MustCompile("\\)"),
	OPEN_BRACE:        *regexp.MustCompile("{"),
	CLOSE_BRACE:       *regexp.MustCompile("}"),
	COMMA:             *regexp.MustCompile(","),
	SEMI_COLUMN:       *regexp.MustCompile(";"),
	PLUS_SIGN:         *regexp.MustCompile("\\+"),
}

var keywords = array.Array[TokenType]{
	FUNCTION,
	RETURN,
}

type Token struct {
	Type    TokenType
	Content string
}

func isKeyword(word string) bool {
	for _, token := range keywords {
		regex := tokenToRegex[token]
		if regex.MatchString(word) {
			return true
		}
	}
	return false
}

func findNextToken(content string) (*Token, int) {
	var result *Token
	var tokenEndIndex int

	isFirstIteration := true

	for tokenType, tokenRegex := range tokenToRegex {
		indexes := tokenRegex.FindStringIndex(content)
		if indexes == nil {
			continue
		}

		if indexes[1] < tokenEndIndex || isFirstIteration {
			tokenContent := content[indexes[0]:indexes[1]]
			switch tokenType {
			case IDENTIFIER:
				if isKeyword(tokenContent) {
					token, _ := keywords.Find(func(tt TokenType) bool {
						rgx := tokenToRegex[tt]
						return rgx.MatchString(tokenContent)
					})
					result = &Token{Type: token}
				} else {
					result = &Token{Type: tokenType, Content: tokenContent}
				}
			default:
				result = &Token{Type: tokenType}
			}
			tokenEndIndex = indexes[1]
			isFirstIteration = false
		}
	}

	if isFirstIteration {
		return nil, 0
	}

	return result, tokenEndIndex
}

func Tokenize(sourceCode string) (array.Array[Token], error) {
	tokens := make(array.Array[Token], 0)
	for content := sourceCode; content != ""; {
		token, tokenEndIndex := findNextToken(content)
		if token == nil {
			break
		}

		tokens.Add(*token)
		content = content[tokenEndIndex:]
	}
	return tokens, nil
}
