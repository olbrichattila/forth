package ast

// NodeDrop represent an instruction to drop the top element from the stack
type NodeDrop struct {
}

// GetType returns the node type
func (*NodeDrop) GetType() NodeType {
	return NodeTypeDrop
}
