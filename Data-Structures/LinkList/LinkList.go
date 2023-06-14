package linklist

import "fmt"

//单向链表

type Node struct {
    Data interface{}
    Next *Node
}

type LList struct {
    Header *Node 
    Length int
}

func CreateNode(v interface{}) *Node{
    return &Node{v, nil}
}
func CreateList()  *LList{
    header := CreateNode(nil)
    return &LList{header,0}
}

func (l *LList)Add(data interface{}){
    newNode := CreateNode(data)
    defer func() {
        l.Length++
    }()

    if l.Length == 0 {
        l.Header = newNode
    }else{
        newNode.Next  = l.Header
        l.Header = newNode 
    }

}

func (l *LList)Append(data interface{}){
    newNode := CreateNode(data)
    defer func() {
        l.Length++
    }()

    if l.Length==0{
        l.Header = newNode
    }
    if l.Length>0 {
        current := l.Header
        for current.Next != nil{ 
            current = current.Next
        }
        current.Next = newNode
    }
}

func (l *LList)Insert(i int, data interface{}){
    defer func() {
        l.Length++
    }()

    if i>=l.Length {
        l.Append(data)
        return
    }
    newNode := CreateNode(data)
    j := 1
    pre := l.Header
    for j!=i{
        pre = pre.Next
        j++
    }
    after := pre.Next
    pre.Next = newNode
    newNode.Next = after
}

func (l *LList)Delete(i int){
    defer func() {
        l.Length--
    }()
    if i==1{ 
        l.Header = l.Header.Next
        return
    }
    j := 0
    current := l.Header
    for j == i-1 {
        current = current.Next
        j++
    }
    after := current.Next.Next
    current.Next = after
}

func (l *LList)Scan(){
    current := l.Header
    i := 1
    for current.Next != nil{
        fmt.Printf("第%d的节点是%d\n", i, current.Data)
        current = current.Next
        i++
    }
    fmt.Printf("第%d的节点是%d\n", i, current.Data)

}
//反转单向链表
func (l *LList)ReverseList(){
    current := l.Header
    var pre *Node
    for current != nil{
        nextNode := current.Next
        current.Next = pre
        pre = current 
        current = nextNode
    }
    l.Header = pre
}
