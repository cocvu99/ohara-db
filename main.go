package main

import "fmt"

const INTERNAL_MAX_KEY = 8

type Node interface{}

type BTreeInternalNode struct {
	nkey     int // number of key = number of node
	keys     [INTERNAL_MAX_KEY]int
	children [INTERNAL_MAX_KEY]*Node
}

// Find the last position so that the key <= find_key
// Last Less or Equal
func (node *BTreeInternalNode) FindLastLE(findKey int) int {
	pos := node.nkey
	for i := 0; i < node.nkey; i++ {
		if node.keys[i] <= findKey {
			pos = i
		}
	}
	return pos
}

// Insert a key-children pair into the Internal Node
func (node *BTreeInternalNode) InsertKV(insertKey int, insertChild Node) {
	// Find last less or equal as position to insert
	pos := node.FindLastLE(insertKey)

}

func main() {
	fmt.Println("Hello World! First step!")
}
