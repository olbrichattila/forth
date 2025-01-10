// Package lexer tokenize the source code
// The lexical analyzer main purpose is to break down raw source code into manageable pieces called tokens
package lexer

import "fmt"

// Lexer abstracts the logic from the concrete implementation
type Lexer interface {
	Tokenize(code string) ([]Token, error)
}

// New creates a new lexer
func New() Lexer {
	return &tokenizer{}
}

var tokens = map[string]TokenType{
	"+":  TokenTypeAdd,
	"-": TokenTypeSub,
	"*": TokenTypeMultiply,
	"dup": TokenTypeDup,
	"drop": TokenTypeDrop,
	"swap": TokenTypeSwap,
	".": TokenTypePrint,
	":": TokenTypeFunction,
	";": TokenEndFunc,
}

type tokenizer struct {
	code string
	pos int
}

func (t *tokenizer) Tokenize(code string) ([]Token, error) {
    result := make([]Token, 0)
    t.code = code
    t.pos = 0

    for t.pos < len(code) {
        char := string(code[t.pos])

        if t.ignoreCharacter(char) {
            t.pos++
            continue
        }

        if intValue, ok := t.parseIntToken(); ok {
            result = append(result, newToken(TokenTypeNumber, intValue))
            continue
        }

        if nameValue, ok := t.parseNameToken(); ok {
            result = append(result, t.resolveNameToken(nameValue))
            continue
        }

        if tokenType, ok := tokens[char]; ok {
            result = append(result, newToken(tokenType, char))
            t.pos++
            continue
        }

        return nil, fmt.Errorf("unexpected character '%s' at position %d", char, t.pos)
    }

    return result, nil
}


func (t *tokenizer) resolveNameToken(tokenName string) Token {
	// if asc code sequence is pre defined like dup, swap... return specific token
	if tk, ok := tokens[tokenName]; ok {
		return newToken(tk, tokenName)
	}

	return newToken(TokenTypeName, tokenName)
}

func (t *tokenizer) parseIntToken() (string, bool) {
    start := t.pos
    for t.pos < len(t.code) && t.isInt(string(t.code[t.pos])) {
        t.pos++
    }
    return t.code[start:t.pos], t.pos > start
}

func (t *tokenizer) parseNameToken() (string, bool) {
    start := t.pos
    for t.pos < len(t.code) && t.isName(string(t.code[t.pos])) {
        t.pos++
    }
    return t.code[start:t.pos], t.pos > start
}

func (t *tokenizer) isInt(s string) bool {
	return s >= "0" && s <= "9"
}

func (t *tokenizer) isName(s string) bool {
	return (s >= "a" && s <= "z") || (s >= "A" && s <= "Z") || (s >= "0" && s <= "9")
}

func (t *tokenizer) ignoreCharacter(s string) bool {
	return s == " " || s == "\n" || s == "\r"  || s == "\t"
}
