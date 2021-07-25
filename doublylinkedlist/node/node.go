package node

type Node struct {
	Val  interface{}
	Prev *Node
	Next *Node
}

func New(value interface{}) *Node {
	return &Node{
		Val:  value,
		Next: nil,
		Prev: nil,
	}
}
