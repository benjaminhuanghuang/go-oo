package tree

import "fmt"

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
