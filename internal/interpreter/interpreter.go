// Package interpreter iterates through an AST tree and executes all nodes
package interpreter

import (
	"fmt"
	"forth/internal/ast"
	"forth/internal/stack"
)

// New creates a new interpreter "instance"
func New() Interpreter {
	return &interpret{
		stack.New(),
	}
}

// Interpreter abstracts out the functionality of code execution
type Interpreter interface {
	Execute(executeAST *ast.Ast) error
}

type interpret struct {
	stack stack.Stacker
}

// Execute runs the AST tree
func (e *interpret) Execute(executeAST *ast.Ast) error {
	fnMap := map[ast.NodeType]func(ast.Node, *ast.Ast) error{
		ast.NodeTypePush: e.push,
		ast.NodeTypeAdd: e.add,
		ast.NodeTypeSub: e.sub,
		ast.NodeTypeMultiply: e.multiply,
		ast.NodeTypePrintResult: e.printResult,
		ast.NodeTypeName: e.functionCall,
		ast.NodeTypeDup: e.dup,
		ast.NodeTypeDrop: e.drop,
		ast.NodeTypeSwap: e.swap,
	}

	for _, n := range executeAST.GetBody() {
		if fn, ok := fnMap[n.GetType()]; ok {
			err := fn(n, executeAST)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (e *interpret) push(n ast.Node, _ *ast.Ast) error {
	if p, ok := n.(*ast.NodePush); ok {
		e.stack.Push(p.Value)
	}

	return nil
}

func (e *interpret) add(_ ast.Node, _ *ast.Ast) error {
	val1, err := e.stack.Pop()
	if err != nil {
		return err
	}

	val2, err := e.stack.Pop()
	if err != nil {
		return err
	}
	e.stack.Push(val1 + val2)

	return nil
}

func (e *interpret) sub(_ ast.Node, _ *ast.Ast) error {
	val1, err := e.stack.Pop()
	if err != nil {
		return err
	}

	val2, err := e.stack.Pop()
	if err != nil {
		return err
	}

	e.stack.Push(val2 - val1)

	return nil
}

func (e *interpret) multiply(_ ast.Node, _ *ast.Ast) error {
	val1, err := e.stack.Pop()
	if err != nil {
		return err
	}

	val2, err := e.stack.Pop()
	if err != nil {
		return err
	}

	e.stack.Push(val1 * val2)
	return nil
}

func (e *interpret) printResult(_ ast.Node, _ *ast.Ast) error {
	last, err := e.stack.Last() 
	if err != nil {
		return err
	}

	fmt.Println(last)

	return nil
}

func (e *interpret) functionCall(n ast.Node, callAst *ast.Ast) error {
	if val, ok := n.(*ast.NodeName); ok {
		if fn, ok := callAst.GetFunctions()[val.Name]; ok {
			if callable, ok := fn.(*ast.NodeFunction); ok {
				e.Execute(callable.Body)
			}
		}
	}

	return nil
}

func (e *interpret) dup(_ ast.Node, _ *ast.Ast) error {
	value, err := e.stack.Last()
	if err != nil {
		return err
	}

	e.stack.Push(value)

	return nil
}

func (e *interpret) drop(_ ast.Node, _ *ast.Ast) error {
	return e.stack.Drop()
}

func (e *interpret) swap(_ ast.Node, _ *ast.Ast) error {
	val1, err := e.stack.Pop()
	if err != nil {
		return err
	}

	val2, err := e.stack.Pop()
	if err != nil {
		return err
	}

	e.stack.Push(val1)
	e.stack.Push(val2)
	return nil
}
