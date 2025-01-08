// Package forth provides a simplified implementation of the Forth language.
// It orchestrates tokenization, Abstract Syntax Tree (AST) construction, and script execution.
package forth

import (
	"fmt"
	"forth/internal/ast"
	"forth/internal/interpreter"
	"forth/internal/lexer"
)

// Run executes the provided Forth source code.
func Run(code string) error {
	if code == "" {
		return fmt.Errorf("input code is empty")
	}

	// Lexer tokenizes the source code.
	lexerInstance := lexer.New()
	tokens, err := lexerInstance.Tokenize(code)
	if err != nil {
		return fmt.Errorf("lexer error: %w", err)
	}

	// AST creates an Abstract Syntax Tree from tokens.
	astInstance := ast.New()
	astTree, err := astInstance.Build(tokens)
	if err != nil {
		return fmt.Errorf("AST build error: %w", err)
	}

	// The interpreter executes the AST.
	interpreterInstance := interpreter.New()
	if err := interpreterInstance.Execute(astTree); err != nil {
		return fmt.Errorf("interpreter execution error: %w", err)
	}

	return nil
}
