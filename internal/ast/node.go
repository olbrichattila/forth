package ast

// NodeType represent the type of the AST node
type NodeType int

// List of AST node types
const (
	NodeTypePush NodeType  = iota
	NodeTypeAdd
	NodeTypeSub
	NodeTypeMultiply
	NodeTypePrintResult
	NodeTypeFunction
	NodeTypeName
	NodeTypeDup
	NodeTypeDrop
	NodeTypeSwap
)

// Node represents an AST node
type Node interface {
	GetType() NodeType
}
