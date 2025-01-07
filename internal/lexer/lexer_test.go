package lexer

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type lexerTestSuite struct {
	suite.Suite
	lexer Lexer
}

func TestLexerRunner(t *testing.T) {
	suite.Run(t, new(lexerTestSuite))
}

func (t *lexerTestSuite) SetupTest() {
	t.lexer = New()
}

func (t *lexerTestSuite) TearDownTest() {
	t.lexer = nil
}

func (t *lexerTestSuite) TestIsSourceTokenized() {
	code := "25 10 ."

	tokens, err := t.lexer.Tokenize(code)
	t.Nil(err)
	t.Len(tokens, 3)

	t.Equal(TokenTypeNumber, tokens[0].GetType())
	t.Equal(TokenTypeNumber, tokens[1].GetType())
	t.Equal(TokenTypePrint, tokens[2].GetType())
}

func (t *lexerTestSuite) TestisFunctionsAndSourceTokenized() {
	code := ":function 50 50 +; function ."

	tokens, err := t.lexer.Tokenize(code)
	t.Nil(err)
	t.Len(tokens, 8)

	t.Equal(TokenTypeFunction, tokens[0].GetType())
	t.Equal(TokenTypeName, tokens[1].GetType())
	t.Equal(TokenTypeNumber, tokens[2].GetType())
	t.Equal(TokenTypeNumber, tokens[3].GetType())
	t.Equal(TokenTypeAdd, tokens[4].GetType())
	t.Equal(TokenEndFunc, tokens[5].GetType())
	t.Equal(TokenTypeName, tokens[6].GetType())
	t.Equal(TokenTypePrint, tokens[7].GetType())
}
