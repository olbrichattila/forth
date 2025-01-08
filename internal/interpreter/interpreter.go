// Package interpreter iterates through an AST tree and executes all nodes
package interpreter

import (
	"fmt"
	"forth/internal/ast"
	"forth/internal/stack"
)

// New creates a new interpreter "instance"
func New() Interpreter {
	interpreter := &interpret{
        stack: stack.New(),
    }
    interpreter.initFnMap()
    return interpreter
}

// Interpreter abstracts out the functionality of code execution
type Interpreter interface {
	Execute(executeAST *ast.Ast) error

}

type interpret struct {
	stack stack.Stacker
	fnMap map[ast.NodeType]func(ast.Node, *ast.Ast) error
}

// Execute runs the AST tree
func (e *interpret) Execute(executeAST *ast.Ast) error {
	for _, n := range executeAST.GetBody() {
		if fn, ok := e.fnMap[n.GetType()]; ok {
			err := fn(n, executeAST)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (e *interpret) initFnMap() {
	e.fnMap = map[ast.NodeType]func(ast.Node, *ast.Ast) error{
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
}

func (e *interpret) push(n ast.Node, _ *ast.Ast) error {
	if p, ok := n.(*ast.NodePush); ok {
		e.stack.Push(p.Value)
	}

	return nil
}

func (e *interpret) add(_ ast.Node, _ *ast.Ast) error {
	val1, val2, err := e.pop2()
	if err != nil {
		return fmt.Errorf("add operation failed: %w", err)
	}

	e.stack.Push(val1 + val2)

	return nil
}

func (e *interpret) sub(_ ast.Node, _ *ast.Ast) error {
	val1, val2, err := e.pop2()
	if err != nil {
		return fmt.Errorf("sub operation failed: %w", err)
	}

	e.stack.Push(val2 - val1)

	return nil
}

func (e *interpret) multiply(_ ast.Node, _ *ast.Ast) error {
	val1, val2, err := e.pop2()
	if err != nil {
		return fmt.Errorf("multiply operation failed: %w", err)
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
	val, ok := n.(*ast.NodeName)
	if !ok {
		return fmt.Errorf("invalid node type for function call: %T", n)
	}

	fn, ok := callAst.GetFunctions()[val.Name]
	if !ok {
		return fmt.Errorf("function %s not found", val.Name)
	}

	callable, ok := fn.(*ast.NodeFunction)
	if !ok {
		return fmt.Errorf("invalid function node type for %s: %T", val.Name, fn)
	}

	if err := e.Execute(callable.Body); err != nil {
		return fmt.Errorf("error executing function %s: %w", val.Name, err)
	}

	return nil
}

func (e *interpret) dup(_ ast.Node, _ *ast.Ast) error {
	value, err := e.stack.Last()
	if err != nil {
		return fmt.Errorf("get last value from stack operation failed: %w", err)
	}

	e.stack.Push(value)

	return nil
}

func (e *interpret) drop(_ ast.Node, _ *ast.Ast) error {
	err := e.stack.Drop()
	if err != nil {
		return fmt.Errorf("drop operation failed: %w", err)
	}

	return nil
}

func (e *interpret) swap(_ ast.Node, _ *ast.Ast) error {
	val1, val2, err := e.pop2()
	if err != nil {
		return fmt.Errorf("swap operation failed: %w", err)
	}

	e.stack.Push(val1)
	e.stack.Push(val2)
	return nil
}

func (e *interpret) pop2() (int, int, error) {
	val1, err := e.stack.Pop()
	if err != nil {
		return 0, 0, fmt.Errorf("value 1 pop operation failed: %w", err)
	}

	val2, err := e.stack.Pop()
	if err != nil {
		return 0, 0, fmt.Errorf("value 2 pop operation failed: %w", err)
	}

	return val1, val2, nil
}
