package main

import "fmt"

// Trie is a prefix tree
type Trie struct {
	isWord bool
	nexts  map[byte]*Trie
}

// Constructor initialize your data structure here.
func Constructor() Trie {
	return Trie{
		isWord: false,
		nexts:  make(map[byte]*Trie),
	}
}

// Insert inserts a word into the trie.
func (t *Trie) Insert(word string) {
	c := word[0]
	if _, ok := t.nexts[c]; !ok {
		nextTrie := &Trie{
			isWord: false,
			nexts:  make(map[byte]*Trie),
		}
		t.nexts[c] = nextTrie
	}

	if len(word) == 1 {
		t.nexts[c].isWord = true
		return
	}

	t.nexts[c].Insert(word[1:])
}

// Search returns if the word is in the trie.
func (t *Trie) Search(word string) bool {
	c := word[0]
	if _, ok := t.nexts[c]; !ok {
		return false
	}

	if len(word) == 1 {
		return t.nexts[c].isWord
	}

	return t.nexts[c].Search(word[1:])
}

// StartsWith returns if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	c := prefix[0]
	if _, ok := t.nexts[c]; !ok {
		return false
	}

	if len(prefix) == 1 {
		return true
	}

	return t.nexts[c].StartsWith(prefix[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// algorithm data structure: Tire, Prefix Tree
func main() {
	trie := Constructor()
	trie.Insert("apple")

	fmt.Println(trie.Search("apple") == true)
	fmt.Println(trie.Search("app") == false)
	fmt.Println(trie.StartsWith("app") == true)

	trie.Insert("app")
	fmt.Println(trie.Search("app") == true)

}
