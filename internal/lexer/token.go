package lexer

// TokenType represents a lexer token type
type TokenType int

// Exported token types
const (
	// Literals
	TokenTypeNumber TokenType = iota
	TokenTypeName

	// Operators
	TokenTypeAdd
	TokenTypeSub
	TokenTypeMultiply

	// Stack manipulation
	TokenTypeDup
	TokenTypeDrop
	TokenTypeSwap

	// Statements
	TokenTypeFunction
	TokenTypePrint
	TokenEndFunc
)

// Initialize new token
func newToken(t TokenType, v string) Token {
    return Token{Type: t, Value: v}
}

// Token is a reference to a keyword, identifier, operator or literal
type Token struct {
	Type TokenType
	Value string
}


// GetType returns the token type
func (t *Token) GetType() TokenType {
	return t.Type
}

// GetValue returns the token value
func (t *Token) GetValue() string {
	return t.Value
}
