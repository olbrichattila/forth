package ast

import (
	"fmt"
	"forth/internal/lexer"
	"strconv"
)

type parserFunc = func(*parser, lexer.Token) (Node, error)

func getParsers() map[lexer.TokenType]parserFunc {
	return map[lexer.TokenType]parserFunc{
		lexer.TokenTypeNumber:   pushToStack,
		lexer.TokenTypeAdd:      parseAdd,
		lexer.TokenTypeSub:      parseSub,
		lexer.TokenTypeMultiply: parseMultiply,
		lexer.TokenTypeDup:      parseDup,
		lexer.TokenTypeDrop:     parseDrop,
		lexer.TokenTypeSwap:     parseSwap,
		lexer.TokenTypePrint:    printResult,
		lexer.TokenTypeFunction: functionToken,
		lexer.TokenTypeName:     name,
	}
}

func pushToStack(_ *parser, t lexer.Token) (Node, error) {
	if v, err := strconv.Atoi(t.GetValue()); err == nil {
		return &NodePush{Value: v}, nil
	}

	return nil, fmt.Errorf("incorrect token value %s", t.GetValue())
}

func parseAdd(_ *parser, _ lexer.Token) (Node, error) {
	return &NodeAdd{}, nil
}

func parseSub(_ *parser, _ lexer.Token) (Node, error) {
	return &NodeSub{}, nil
}

func parseMultiply(_ *parser, _ lexer.Token) (Node, error) {
	return &NodeMultiply{}, nil
}

func parseDup(_ *parser, _ lexer.Token) (Node, error) {
	return &NodeDup{}, nil
}

func parseDrop(_ *parser, _ lexer.Token) (Node, error) {
	return &NodeDrop{}, nil
}

func parseSwap(_ *parser, _ lexer.Token) (Node, error) {
	return &NodeSwap{}, nil
}

func printResult(_ *parser, _ lexer.Token) (Node, error) {
	return &NodePrintResult{}, nil
}

func functionToken(p *parser, _ lexer.Token) (Node, error) {
	if !p.expect(lexer.TokenTypeName) {
		return nil, fmt.Errorf("function : should be followed by a function name")
	}

	body := &Ast{};
	nextToken := p.at();
	
	for {
		if p.eof() || p.expect(lexer.TokenEndFunc) {
			break
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

func name(_ *parser, t lexer.Token) (Node, error) {
	return &NodeName{Name: t.GetValue()}, nil
}