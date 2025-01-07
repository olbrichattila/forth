// Package ast parses a token list into a Abstract Syntax Tree
// The Abstract Syntax Tree (AST) is a fundamental data structure used in programming language
// interpreters and compilers.
// Its primary purpose is to provide a structured, hierarchical representation of the source code
// that is easier to analyze and manipulate than raw text or tokens.
package ast

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

