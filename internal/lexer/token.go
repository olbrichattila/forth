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

// Token is a reference to a keyword, identifier, operator or literal
type Token struct {
	t TokenType
	value string
}

// GetType returns the token type
func (t *Token) GetType() TokenType {
	return t.t
}

// GetValue returns the token value
func (t *Token) GetValue() string {
	return t.value
}
