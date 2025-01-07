package ast

// NodeDup is an instruction to duplicate the top item on the stack.
type NodeDup struct {
}

// GetType returns the node type
func (*NodeDup) GetType() NodeType {
	return NodeTypeDup
}