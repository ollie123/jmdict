package index

import "jmdict"

// RuneTree represents a prefix tree (or trie) mapping []rune to *jmdict.Entry.
type RuneTree struct {
	Entries  []*jmdict.Entry
	Children map[rune]*RuneTree
}

// NewRuneTree returns a new RuneTree.
func NewRuneTree() *RuneTree {
	return &RuneTree{
		Children: make(map[rune]*RuneTree),
	}
}

// Add adds an entry to the tree for the specified key.
func (t *RuneTree) Add(key []rune, entry *jmdict.Entry) {
	node := t
	for _, r := range key {
		child, ok := node.Children[r]
		if !ok {
			child = NewRuneTree()
			node.Children[r] = child
		}
		node = child
	}
	node.Entries = append(node.Entries, entry)
}

// KeyFunc is the type of function called to derive keys from a *jmdict.Entry.
type KeyFunc func(*jmdict.Entry) [][]rune

// AddAll adds all of the entries to the tree with keys determined by keyFn.
func (t *RuneTree) AddAll(keyFn KeyFunc, entries []*jmdict.Entry) {
	for i, entry := range entries {
		keys := keyFn(entry)
		for _, key := range keys {
			t.Add(key, entries[i])
		}
	}
}

// Get returns the node for a key, or nil, false if the key is not present.
func (t *RuneTree) Get(key []rune) (*RuneTree, bool) {
	node := t
	for _, r := range key {
		child, ok := node.Children[r]
		if !ok {
			return nil, false
		}
		node = child
	}
	return node, true
}

// WalkFunc is the type of function called for each node visited by Walk.
// If the function returns false for a node, its subtree is not walked.
type WalkFunc func(*RuneTree) bool

// Walk walks thef tree, calling walkFn for each child.
func (t *RuneTree) Walk(walkFn WalkFunc) {
	if walkFn(t) {
		for _, child := range t.Children {
			child.Walk(walkFn)
		}
	}
}
