package ast

// NodeAdd represents an operator node that adds the top two numbers on the stack.
type NodeAdd struct {
}

// GetType returns the type of this node, which is NodeTypeAdd.
func (*NodeAdd) GetType() NodeType {
	return NodeTypeAdd
}