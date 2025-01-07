package lexer

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type tokenTestSuite struct {
	suite.Suite
}

func TestTokenRunner(t *testing.T) {
	suite.Run(t, new(tokenTestSuite))
}

func (t *tokenTestSuite) TestToken() {
	expectedValue := "1234";
	token := &Token{t: TokenTypeAdd, value: expectedValue}

	t.Equal(TokenTypeAdd, token.GetType())
	t.Equal(expectedValue, token.GetValue())
}
