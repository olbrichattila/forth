package ast

// NodeFunction stores a function with it's content in the Body as an Ast sub tree
type NodeFunction struct {
	Name string
	Body *Ast
}

// GetType returns the node type
func (*NodeFunction) GetType() NodeType {
	return NodeTypeFunction
}