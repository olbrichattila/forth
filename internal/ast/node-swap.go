package ast

// NodeSwap is an instruction to swap the two values at the top of the stack
type NodeSwap struct {
}

// GetType returns the node type
func (*NodeSwap) GetType() NodeType {
	return NodeTypeSwap
}