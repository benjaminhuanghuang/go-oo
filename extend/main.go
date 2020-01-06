package main
import (
	"./tree"
)
type myTreeNode struct {
	node *tree.Node
}


func (myNode  *myTreeNode) postOrder(){
	if(myNode == nil || myNode.node == nil ){
		return
	}
	left:=myTreeNode{myNode.node.Left}
	left.postOrder()
	right:=myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}