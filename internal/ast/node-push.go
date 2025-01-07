package ast

// NodePush is an instruction to push the value to the stack
type NodePush struct {
	Value int
}

// GetType returns the node type
func (*NodePush) GetType() NodeType {
	return NodeTypePush
}
