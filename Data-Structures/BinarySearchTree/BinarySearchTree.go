package binarysearchtree

import "fmt"
//搜索二叉树
type Node struct{
	Value int
	leftNode *Node
	rightNode *Node
}

type BinaryTree struct{
	len int
	root *Node
}

func NewNode(data int) *Node{
	return &Node{data,nil,nil}
}

func(n *Node)Insert(data int){
	if n.Value>data{
		if n.leftNode!=nil{
			n.leftNode.Insert(data)
		}else{
			n.leftNode=NewNode(data)
		}
	}else{
		if n.rightNode!=nil{
			n.rightNode.Insert(data)
		}else{
			n.rightNode=NewNode(data)
		}
	}
}

func(n *Node)middleShow(){
	if n.leftNode!=nil{
		n.leftNode.middleShow()
	}
	fmt.Println(n.Value)
	if n.rightNode!=nil{
		n.rightNode.middleShow()
	}
}

func(n *Node)search(data int) *Node{
	if n.Value==data{
		return n
	}else if n.Value>data{
		if n.leftNode!=nil{
			return n.leftNode.search(data)
		}
	}else{
		if n.rightNode!=nil{
			return n.rightNode.search(data)
		}
	}
	return nil
}

func NewTree()*BinaryTree{
	return &BinaryTree{0,nil}
}

func(t *BinaryTree)Insert(data ...int){
	if t.root!=nil{
		for _,d:=range data{
			t.len++
			t.root.Insert(d)
		}
	}else{
		t.root=NewNode(data[0])
		t.Insert(data[1:]...)
	}
}

func(t *BinaryTree)Search(data int) *Node{
	if t.root!=nil{
		return t.root.search(data)
	}else{
		return nil
	}
}

func(t *BinaryTree)MiddleShow(){
	if t.root==nil{
		return
	}
	t.root.middleShow()
}
