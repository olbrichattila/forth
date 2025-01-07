package ast

// NodeName stores an ASC word, az-AZ
// TODO could add validation to some of those
type NodeName struct {
	Name string
}

// GetType returns the node type
func (*NodeName) GetType() NodeType {
	return NodeTypeName
}
