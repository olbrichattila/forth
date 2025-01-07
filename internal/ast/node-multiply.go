package ast

// NodeMultiply is a node for multiplying the two number at the top of the stack
type NodeMultiply struct {
}

// GetType returns the node type
func (*NodeMultiply) GetType() NodeType {
	return NodeTypeMultiply
}