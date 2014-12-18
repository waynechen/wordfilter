package trie

var sTrie *Trie

func Singleton() *Trie {
	if sTrie == nil {
		sTrie = NewTrie()
	}
	return sTrie;
}
