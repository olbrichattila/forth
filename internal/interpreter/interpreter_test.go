package interpreter

import (
	"forth/internal/ast"
	"forth/internal/stdcapture"
	"testing"

	"github.com/stretchr/testify/suite"
)

type interpreterTestSuite struct {
	suite.Suite
	interpreter Interpreter
	capturer *stdcapture.StdoutCapture
}

func TestRunner(t *testing.T) {
	suite.Run(t, new(interpreterTestSuite))
}

func (t *interpreterTestSuite) SetupTest() {
	t.capturer = stdcapture.New()
	t.interpreter = New()
}

func (t *interpreterTestSuite) TearDownTest() {
	t.interpreter = nil
	t.capturer = nil
}

func (t *interpreterTestSuite) TestAstExecutes() {
	// Start capturing str output
	t.capturer.StartCapture()
	defer t.capturer.StopCapture()

	// Add 10 10 . (which prints 20)
	astNodes := &ast.Ast{}
	astNodes.AddNode(&ast.NodePush{Value: 10})
	astNodes.AddNode(&ast.NodePush{Value: 20})
	astNodes.AddNode(&ast.NodePrintResult{})

	
	// Add function: testFunction 50 . (which prints 50)
	functionNodes := &ast.Ast{}
	functionNodes.AddNode(&ast.NodePush{Value: 50})
	functionNodes.AddNode(&ast.NodePrintResult{})
	functionNode := &ast.NodeFunction{Name: "testFunction", Body: functionNodes}
	
	astNodes.AddFunction("testFunc", functionNode)
	astNodes.AddNode(&ast.NodeName{Name: "testFunc"})

	err := t.interpreter.Execute(astNodes)
	output := t.capturer.StopCapture()

	t.Nil(err)
	// 20 and 50 printed
	t.Equal("20\n50\n", output)
}

