package main

import "fmt"

const INTERNAL_MAX_KEY = 8

type Node interface{}

type BTreeInternalNode struct {
	nkey     int // number of key = number of node
	keys     [INTERNAL_MAX_KEY]int
	children [INTERNAL_MAX_KEY]*Node
}

func NewINode() BTreeInternalNode {
	var new_keys [INTERNAL_MAX_KEY]int
	var new_children [INTERNAL_MAX_KEY]*Node
	return BTreeInternalNode{
		nkey:     0,
		keys:     new_keys,
		children: new_children,
	}
}

// Find the last position so that the key <= find_key
// Last Less or Equal
func (node *BTreeInternalNode) FindLastLE(findKey int) int {
	pos := -1
	for i := 0; i < node.nkey; i++ {
		if node.keys[i] <= findKey {
			pos = i
			break
		}
	}
	return pos
}

// Insert a key-children pair into the Internal Node
func (node *BTreeInternalNode) InsertKV(insertKey int, insertChild Node) {
	// Find last less or equal as position to insert
	pos := node.FindLastLE(insertKey)
	// [ 1,4,7,| | ] -> insert 3
	// [ 1,| |,4,7 ] -> insert 3
	// [ ] -> position = -1
	for i := node.nkey - 1; i > pos; i-- {
		node.keys[i+1] = node.keys[i]
		node.children[i+1] = node.children[i]
	}
	node.keys[pos+1] = insertKey
	node.children[pos+1] = &insertChild
	// [1,3,4,7]
	node.nkey += 1
}

func main() {
	fmt.Println("Hello World! First step!")
}
