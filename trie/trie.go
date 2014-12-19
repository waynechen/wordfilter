package trie

import "sync"

type Trie struct{
	Root *TrieNode
	Mutex sync.RWMutex
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

// 添加一个敏感词(UTF-8的)
func (t *Trie) Add(keyword string) {
	chars := []rune(keyword)

	if len(chars) == 0 {
		return
	}

	t.Mutex.Lock()

	node := t.Root
	for _, char := range chars {
		if _, ok := node.Node[char]; !ok {
			node.Node[char] = NewTrieNode()
		}
		node = node.Node[char]
	}
	node.End = true

	t.Mutex.Unlock()
}

// 删除一个敏感词
func (t *Trie) Del(keyword string) {
	chars := []rune(keyword)
	if len(chars) == 0 {
		return
	}

	t.Mutex.Lock()
	node := t.Root
	t.cycleDel(node, chars, 0)
	t.Mutex.Unlock()
}

func (t *Trie) cycleDel(node *TrieNode, chars []rune, index int) (shouldDel bool) {
	char := chars[index]
	l := len(chars)

	if n, ok := node.Node[char]; ok {
		if index+1 < l {
			shouldDel = t.cycleDel(n, chars, index+1)
			if shouldDel {
				if n.End { // 说明这是一个敏感词，不能删除
					shouldDel = false
				}else {
					delete(node.Node, char)
				}
			}
		}else if n.End {
			if len(n.Node) == 0 { // 是最后一个节点
				shouldDel = true
			}else { // 不是最后一个节点
				n.End = false
			}
		}
	}

	return
}

// 查找替换
// 将text中在trie里的敏感字，替换为*号
// 返回结果: 是否有敏感字, 敏感字数组, 替换后的文本
func (t *Trie) Replace(text string) (bool, []string, string) {
	found := []string{}
	chars := []rune(text)
	l := len(chars)
	if l == 0 {
		return false, found, text
	}

	var (
		i, j, jj int
		ok bool
	)

	node := t.Root
	for i = 0; i < l; i++ {
		if _, ok = node.Node[chars[i]]; !ok {
			continue
		}

		jj = 0

		node = node.Node[chars[i]]
		for j = i+1; j < l; j++ {
			if _, ok = node.Node[chars[j]]; !ok {
				if jj > 0 {
					found = t.replaceToAsterisk(found, chars, i, jj)
					i = jj
				}
				break
			}

			node = node.Node[chars[j]]
			if node.End {
				jj = j //还有子节点的情况, 记住上次找到的位置, 以匹配最大串 (eg: AV, AV女优)

				if len(node.Node) == 0 { // 是最后节点, break
					found = t.replaceToAsterisk(found, chars, i, j)
					i = j
					break;
				}
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

// 替换为*号
func (t *Trie) replaceToAsterisk(found []string, chars []rune, i, j int) []string {
	tmpFound := chars[i:j+1]
	found = append(found, string(tmpFound))
	for k := i; k <= j; k++ {
		chars[k] = 42 // *的rune为42
	}
	return found
}

func (t *Trie) ReadAll() (words []string) {
	t.Mutex.Lock()
	words = []string{}
	words = t.cycleRead(t.Root, words, "")
	t.Mutex.Unlock()
	return
}

func (t *Trie) cycleRead(node *TrieNode, words []string , parentWord string) []string {
	for char, n := range node.Node {
		if n.End {
			words = append(words, parentWord+string(char))
		}
		if len(n.Node) > 0 {
			words = t.cycleRead(n, words, parentWord+string(char))
		}
	}
	return words
}
