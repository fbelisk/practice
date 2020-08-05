package trieTree

type TrieTree struct {
	Root *Node
}

type Node struct {
	IsWord bool
	Next   map[byte]*Node
}

func (t *TrieTree) Insert(word []byte) {
	wordLen := len(word)
	node := t.Root
	for k, v := range word {
		letterNode, ok := node.Next[v]
		if !ok {
			letterNode = &Node{
				Next:map[byte]*Node{},
			}
		}
		if k == wordLen - 1 {
			letterNode.IsWord = true
		}
		node.Next[v] = letterNode
		node = letterNode
	}
}
