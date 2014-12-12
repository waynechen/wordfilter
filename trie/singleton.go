package trie

import (
	"fmt"
)

var sTrie *Trie

func Singleton() *Trie {
	if sTrie == nil {
		sTrie = NewTrie()
	}
	return sTrie;
}

func PrintTrie() {
	t := Singleton()
	node := t.Root
	print(node, " |")
}

func print(node *TrieNode, line string) {
	if len(node.Node) > 0 {
		for char, n := range node.Node {
			fmt.Printf("%s%s\n", line, string(char))
			print(n, line+" |")
		}
	}
}
