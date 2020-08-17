package lru

type LRUCache struct {
	Cache map[int]*Node
	Head *Node
	Tail *Node
	Capacity int
}

type Node struct {
	Value int
	Prev *Node
	Next *Node
	Key int
}

func (n *Node) AddToHead(l *LRUCache)  {
	n.Next = l.Head
	n.Prev = nil
	if l.Head != nil {
		l.Head.Prev = n
	}
	l.Head = n
}

func (n *Node) RemoveTail(l *LRUCache)  {
	l.Tail = n.Prev
	l.Tail.Next = nil
	n.Next = nil
	n.Prev = nil
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cache: make(map[int]*Node, capacity),
		Head: nil,
		Tail: nil,
		Capacity: capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	if v, ok := this.Cache[key]; ok {
		v.AddToHead(this)
		return v.Value
	}
	return 0
}


func (this *LRUCache) Put(key int, value int)  {
	n := &Node{
		Value: value,
		Key: key,
		Prev:  nil,
		Next:  nil,
	}

	if len(this.Cache) == 0 {
		this.Tail = n
	}

	n.AddToHead(this)
	this.Cache[key] = n
	if  this.Capacity == len(this.Cache) {
		delete(this.Cache, this.Tail.Key)
		this.Tail.RemoveTail(this)
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
