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

// Build convert tokens created with lexer to an AST tree
func (p *build) Build(tokens []lexer.Token) (*Ast, error) {
	p.tokens = tokens
	Ast := &Ast{};

	for {
		if p.eof() {
			break;
		}

		token := p.at()
		Node, err := p.parse(token);
		if err != nil {
			return nil, err
		}

		if token.GetType() == lexer.TokenTypeFunction {
			if n, ok := Node.(*NodeFunction); ok {
				Ast.AddFunction(n.Name, Node)
			}

			p.pos++;
			continue
		}
		
		Ast.body = append(Ast.body, Node)
		p.pos++
	}

	return Ast, nil
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
	parsers := getParsers()
	if result, ok := parsers[token.GetType()]; ok {
		return result(p, token)
	}
	
	return nil, fmt.Errorf("token type %d not defined", token.GetType())
}
