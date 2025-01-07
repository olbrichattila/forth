package ast

// NodeAdd is an operator node to add the two number together at the top of the stack
type NodeAdd struct {
}

// GetType returns the node type
func (*NodeAdd) GetType() NodeType {
	return NodeTypeAdd
}