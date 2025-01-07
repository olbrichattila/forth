package ast

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type nodeTestSuite struct {
	suite.Suite
}

func TestNodeRunner(t *testing.T) {
	suite.Run(t, new(nodeTestSuite))
}

func (t *nodeTestSuite) SetupTest() {
	
}

func (t *nodeTestSuite) TestNodeDrop() {
	node := &NodeDrop{}
	t.Equal(NodeTypeDrop, node.GetType())
}

func (t *nodeTestSuite) TestNodeAdd() {
	node := &NodeAdd{}
	t.Equal(NodeTypeAdd, node.GetType())
}

func (t *nodeTestSuite) TestNodeDup() {
	node := &NodeDup{}
	t.Equal(NodeTypeDup, node.GetType())
}

func (t *nodeTestSuite) TestNodeFunction() {
	node := &NodeFunction{}
	t.Equal(NodeTypeFunction, node.GetType())
}

func (t *nodeTestSuite) TestNodeMultiply() {
	node := &NodeMultiply{}
	t.Equal(NodeTypeMultiply, node.GetType())
}

func (t *nodeTestSuite) TestNodeName() {
	node := &NodeName{}
	t.Equal(NodeTypeName, node.GetType())
}

func (t *nodeTestSuite) TestNodePrint() {
	node := &NodePrintResult{}
	t.Equal(NodeTypePrintResult, node.GetType())
}

func (t *nodeTestSuite) TestNodePush() {
	node := &NodePush{}
	t.Equal(NodeTypePush, node.GetType())
}

func (t *nodeTestSuite) TestNodeSub() {
	node := &NodeSub{}
	t.Equal(NodeTypeSub, node.GetType())
}

func (t *nodeTestSuite) TestNodeSwap() {
	node := &NodeSwap{}
	t.Equal(NodeTypeSwap, node.GetType())
}
