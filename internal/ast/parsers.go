package ast

import (
	"fmt"
	"forth/internal/lexer"
	"strconv"
)

type parserFullFunc = func(*build, lexer.Token) (Node, error)
type parserLexerFunc = func(lexer.Token) (Node, error)
type parserBuildFunc = func(*build) (Node, error)
type parserSimpleFunc = func() (Node, error)

func getParsers() map[lexer.TokenType]interface{} {
	return map[lexer.TokenType]interface{}{
		lexer.TokenTypeNumber:   evalPushToStack,
		lexer.TokenTypeAdd:      evalAdd,
		lexer.TokenTypeSub:      evalSub,
		lexer.TokenTypeMultiply: evalMultiply,
		lexer.TokenTypeDup:      evalDup,
		lexer.TokenTypeDrop:     evalDrop,
		lexer.TokenTypeSwap:     evalSwap,
		lexer.TokenTypePrint:    evalPrint,
		lexer.TokenTypeFunction: evalFunction,
		lexer.TokenTypeName:     evalName,
	}
}

func evalPushToStack(t lexer.Token) (Node, error) {
	if v, err := strconv.Atoi(t.GetValue()); err == nil {
		return &NodePush{Value: v}, nil
	}

	return nil, fmt.Errorf("incorrect token value %s", t.GetValue())
}

func evalAdd() (Node, error) {
	return &NodeAdd{}, nil
}

func evalSub() (Node, error) {
	return &NodeSub{}, nil
}

func evalMultiply() (Node, error) {
	return &NodeMultiply{}, nil
}

func evalDup() (Node, error) {
	return &NodeDup{}, nil
}

func evalDrop() (Node, error) {
	return &NodeDrop{}, nil
}

func evalSwap() (Node, error) {
	return &NodeSwap{}, nil
}

func evalPrint() (Node, error) {
	return &NodePrintResult{}, nil
}

func evalFunction(p *build) (Node, error) {
	if !p.expect(lexer.TokenTypeName) {
		return nil, fmt.Errorf("function : should be followed by a function name")
	}

	body := &Ast{};
	nextToken := p.at();
	
	for {
		if p.expect(lexer.TokenEndFunc) {
			break
		}

		if p.eof() {
			return nil, fmt.Errorf("word (function) must end with semicolon")
		}
		
		fnToken := p.at()
		fnNode, err := p.parse(fnToken)
		if err != nil {
			return nil, err;
		}
		body.AddNode(fnNode)
	}

	return &NodeFunction{Name: nextToken.GetValue(), Body: body}, nil
}

func evalName(t lexer.Token) (Node, error) {
	return &NodeName{Name: t.GetValue()}, nil
}
