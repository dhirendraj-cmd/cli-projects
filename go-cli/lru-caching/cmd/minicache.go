package cmd

import "fmt"


type Node struct{
	Key int
	Val any
	Next *Node
	Prev *Node
}

func NewNode(key int, val any) *Node{
	return &Node{
		Key: key,
		Val: val,
		Next: nil,
		Prev: nil,
	}
}

type Doubly struct{
	headNode *Node
	tailNode *Node
}

func NewDoubly() *Doubly{
	return &Doubly{
		headNode: nil,
		tailNode: nil,
	}
}

type LRUCache struct{
	Capacity int
	CacheDict map[any]*Node
	List *Doubly
}

func NewLRUCache(capacity int) *LRUCache{
	return &LRUCache{
		Capacity: capacity,
		CacheDict: make(map[any]*Node),
		List: NewDoubly(),
	}
}

func (d *Doubly) IsEmpty() bool{
	return d.headNode == nil
}

func (d *Doubly) AddInBeginning(node *Node){
	if d.IsEmpty(){
		d.headNode = node
		d.tailNode = node
		node.Next = nil
		node.Prev = nil
		return
	}

	node.Prev = nil
	node.Next = d.headNode

	if d.headNode != nil {
        d.headNode.Prev = node
    }

	d.headNode = node
}

func (d *Doubly) RemoveStaleNode(node *Node){
	if node == nil{
		return
	}

	if node.Prev != nil{
		node.Prev.Next = node.Next
	} else {
		d.headNode = node.Next
	}

	if node.Next != nil{
		node.Next.Prev = node.Prev
	} else {
		d.tailNode = node.Prev
	}

	node.Next = nil
	node.Prev = nil

}

func (d *Doubly) MovetoFront(node *Node){
	d.RemoveStaleNode(node)
	d.AddInBeginning(node)
}

func (lru *LRUCache) Put(key int, val any) {
	// checking if key already exists
	if node, ok := lru.CacheDict[key]; ok{
		node.Val = val
		lru.List.MovetoFront(node)
	} else {

		// checking if new key and mapping
		newNode := NewNode(key, val)
		lru.List.AddInBeginning(newNode)
		lru.CacheDict[key] = newNode
	}


	// checking if size > Capacity
	if len(lru.CacheDict) > lru.Capacity{
		fmt.Println("Sie limit Exceded")
		lastNode := lru.List.tailNode
		if lastNode != nil{
			lru.List.RemoveStaleNode(lastNode)
			delete(lru.CacheDict, lastNode.Key)
		}
	}

	// for k, v := range lru.CacheDict{
	// 	fmt.Printf("Key: %v, Value: %v\n", k, v.Val)
	// }
	lru.PrintCache()
}

func (lru *LRUCache) Get(key int) (any, bool){
	if node, ok := lru.CacheDict[key]; ok{
		lru.List.MovetoFront(node)
		fmt.Println("Key Found: ", node.Val)
		lru.PrintCache()
		return node.Val, true
	}

	fmt.Println("Key Not found!!")
	return "-1", false
}

func (lru *LRUCache) PrintCache(){
	temp := lru.List.headNode
	for temp != nil {
		fmt.Printf("[%v: %v] <-> ", temp.Key, temp.Val)
		temp = temp.Next
	}
	fmt.Println("nil")
}



func MiniLRUCache(){
	fmt.Println("Implementing Mini Cache!!!")

	lru := NewLRUCache(3)
	lru.Put(1, "A")
	lru.Put(2, "B")
	lru.Put(3, "C")
	
	lru.Get(2)
	// lru.Put(4, "E")
	lru.Put(2, "D")
}