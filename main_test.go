package main

import "testing"

func TestDatabase(t *testing.T) {
	node := NewINode()
	childnode := NewINode()
	node.InsertKV(3, childnode)
	// [3]
	if node.nkey != 1 {
		t.Errorf("Got nkey = %v, expect %v", node.nkey, 1)
	}
	if node.keys[0] != 3 {
		t.Errorf("Got key[0] = %v, expect %v", node.keys[0], 1)
	}

	node.InsertKV(10, childnode)
	// [3, 10]
	if node.nkey != 2 {
		t.Errorf("Got nkey = %v, expect %v", node.nkey, 2)
	}
	if node.keys[0] != 3 {
		t.Errorf("Got key[0] = %v, expect %v", node.keys[0], 3)
	}
	if node.keys[1] != 10 {
		t.Errorf("Got key[0] = %v, expect %v", node.keys[1], 10)
	}

	node.InsertKV(5, childnode)
	// [3, 5, 10]
	if node.nkey != 3 {
		t.Errorf("Got nkey = %v, expect %v", node.nkey, 3)
	}
	if node.keys[0] != 3 {
		t.Errorf("Got key[0] = %v, expect %v", node.keys[0], 3)
	}
	if node.keys[1] != 5 {
		t.Errorf("Got key[0] = %v, expect %v", node.keys[1], 5)
	}
	if node.keys[2] != 10 {
		t.Errorf("Got key[0] = %v, expect %v", node.keys[2], 10)
	}

}
