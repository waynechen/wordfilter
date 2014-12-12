package trie

type Trie struct{
	Root *TrieNode
}

type TrieNode struct{
	Node map[rune]*TrieNode
	End bool
}

func NewTrie() *Trie {
	t := new(Trie)
	t.Root = NewTrieNode()
	return t
}

func NewTrieNode() *TrieNode {
	n := new(TrieNode)
	n.Node = make(map[rune]*TrieNode)
	n.End = false
	return n
}

// 输入一个UTF8的 string, 创建
func (t *Trie) Add(keyword string) {
	chars := []rune(keyword)
	if len(chars) == 0 {
		return
	}

	node := t.Root

	for _, char := range chars {
		if _, ok := node.Node[char]; !ok {
			node.Node[char] = NewTrieNode()
		}
		node = node.Node[char]
	}
	node.End = true
}

// 将text中在trie里的关键字，替换为*号
// 返回结果: 是否有关键字, 关键字数组, 替换后的文本
func (t *Trie) Replace(text string) (bool, []string, string) {
	found := []string{}
	chars := []rune(text)
	l := len(chars)
	if l == 0 {
		return false, found, text
	}

	var (
		i, j, k int
		tmpFound []rune
		ok bool
	)

	node := t.Root
	for i = 0; i < l; i ++ {
		if _, ok = node.Node[chars[i]]; !ok {
			continue
		}

		tmpFound = []rune{}
		tmpFound = append(tmpFound, chars[i])

		node = node.Node[chars[i]]

		for j = i+1; j < l; j++ {
			if _, ok = node.Node[chars[j]]; !ok {
				break
			}

			tmpFound = append(tmpFound, chars[j])

			node = node.Node[chars[j]]
			if node.End {
				for k = i; k <= j; k++ {
					chars[k] = 42 // *的rune为42
				}

				found = append(found, string(tmpFound))
				i = j
				break;
			}
		}
		node = t.Root
	}

	exist := false
	if len(found) > 0 {
		exist = true
	}

	return exist, found, string(chars)
}
