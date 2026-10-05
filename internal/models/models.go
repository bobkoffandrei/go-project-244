package models

type Node struct {
	Type     string
	Key      string
	Value    interface{}
	OldValue interface{}
	Children []Node
}

const (
	UNCHANGED = "unchanged"
	ADDED     = "added"
	REMOVED   = "removed"
	CHANGED   = "changed"
	NESTED    = "nested"
)
