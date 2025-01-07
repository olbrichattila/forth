package ast

// NodePrintResult is an instruction to print out the value at the top of the stack
type NodePrintResult struct {
}

// GetType returns the node type
func (*NodePrintResult) GetType() NodeType {
	return NodeTypePrintResult
}