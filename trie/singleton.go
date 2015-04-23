package trie

var _blackTrie *Trie
var _whitePrefixTrie *Trie
var _whiteSuffixTrie *Trie

func BlackTrie() *Trie {
	if _blackTrie == nil {
		NewBlackTrie()
	}
	return _blackTrie
}

func WhitePrefixTrie() *Trie {
	if _whitePrefixTrie == nil {
		_whitePrefixTrie = NewTrie()
	}
	return _whitePrefixTrie
}

func WhiteSuffixTrie() *Trie {
	if _whiteSuffixTrie == nil {
		_whiteSuffixTrie = NewTrie()
	}
	return _whiteSuffixTrie
}

func NewBlackTrie() {
	_blackTrie = NewTrie()
	_blackTrie.CheckWhiteList = true
}

func NewWhitePrefixTrie() {
	_whitePrefixTrie = NewTrie()
}

func NewWhiteSuffixTrie() {
	_whiteSuffixTrie = NewTrie()
}
