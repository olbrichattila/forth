package ast

import (
	"fmt"
	"forth/internal/lexer"
)

// Builder abstracts the concrete implementation of the Builder
type Builder interface {
	Build(tokens []lexer.Token) (*Ast, error)
}

// New creates a new AST builder
func New() Builder {
	return &build{}
}

type build struct {
	tokens []lexer.Token
	pos int
}

// Build constructs an AST from a slice of lexer tokens.
func (p *build) Build(tokens []lexer.Token) (*Ast, error) {
	p.tokens = tokens
	ast := &Ast{};

	for !p.eof() {
		token := p.at()
		node, err := p.parse(token)
		if err != nil {
			return nil, err
		}
	
		if token.GetType() == lexer.TokenTypeFunction {
			if functionNode, ok := node.(*NodeFunction); ok {
				ast.AddFunction(functionNode.Name, functionNode)
			}
		} else {
			ast.body = append(ast.body, node)
		}
		p.pos++
	}

	return ast, nil
}

func (p *build) eof() bool {
	return p.pos >= len(p.tokens) 
}

func (p *build) next() lexer.Token {
    p.pos++
	if p.eof() {
        return lexer.Token{}
    }
    return p.tokens[p.pos]
}

func (p *build) at() lexer.Token {
	return p.tokens[p.pos]
}

func (p *build) expect(tt lexer.TokenType) bool {
	if p.eof() {
		return false
	}

	token := p.next()
	return token.GetType() == tt
}

func (p *build) parse(token lexer.Token) (Node, error) {
	tokenType := token.GetType()
	parsers := getParsers()
	if result, ok := parsers[tokenType]; ok {
		return result(p, token)
	}
	
	return nil, fmt.Errorf("unhandled token type: %d (%s)", tokenType, token.GetValue())
}
