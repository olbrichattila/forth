// Package ast parses a token list into a Abstract Syntax Tree
// The Abstract Syntax Tree (AST) is a fundamental data structure used in programming language
// interpreters and compilers.
// Its primary purpose is to provide a structured, hierarchical representation of the source code
// that is easier to analyze and manipulate than raw text or tokens.
package ast

import (
	"fmt"
	"forth/internal/lexer"
)

// Ast Abstract Syntax Tree with the main code flow (body) and functions as a separate AST tree
type Ast struct {
	functions map[string]Node
	body []Node
}

// AddNode Add a new node to the AST
func (a *Ast) AddNode(n Node)  {
	a.body = append(a.body, n)
}

// AddFunction Adds a new function to the AST
func (a *Ast) AddFunction(name string, n Node)  {
	if a.functions == nil {
		a.functions = make(map[string]Node)
	}
	a.functions[name] = n
}

// GetBody returns the main AST list, except functions
func (a *Ast) GetBody() []Node {
	return a.body
}

// GetFunctions returns the defined functions as an AST tree, searchable hash map
func (a *Ast) GetFunctions() map[string]Node {
	return a.functions
}

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
	p.pos++;
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
