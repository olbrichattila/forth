// Package forth is a simplified implementation of the forth language (Coding KATA)
// This package orchestrates tokenization, creating AST tree and runs.
package forth

import (
	"forth/internal/ast"
	"forth/internal/interpreter"
	"forth/internal/lexer"
)

// Run will execute the source code
func Run(code string) error {
	// Lexer tokenize the source code
	lexer := lexer.New()
	tokens, err := lexer.Tokenize(code)
	if err != nil {
		return err
	}

	// Ast creates a Abstract Syntax Tree
	ast := ast.New()
	astTree, err := ast.Build(tokens)
	if err != nil {
		return err
	}

	// the interpreter executes the Ast, running the script
	interpreter := interpreter.New()
	return interpreter.Execute(astTree);
}
