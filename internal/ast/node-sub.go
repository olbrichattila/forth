package ast

// NodeSub is a node to subtract the two numbers at the top of the stack
type NodeSub struct {
}

// GetType returns the node type
func (*NodeSub) GetType() NodeType {
	return NodeTypeSub
}