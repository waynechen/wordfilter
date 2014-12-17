package trie

import (
	"testing"
)

func TestAdd(t *testing.T) {
	trie := NewTrie()
	trie.Add("中华人民共和国")
	trie.Add("中国")
	trie.Add("中国共产党")
	trie.Add("中国人民解放军")
	trie.Add("中国人民武警")
	trie.Add("华人")
	trie.Add("我men是")

	node := trie.Root
	printTrie(node, t, " |")
}

func printTrie(node *TrieNode, t *testing.T, line string) {
	if len(node.Node) > 0 {
		for char, n := range node.Node {
			t.Logf("%s%s", line, string(char))
			printTrie(n, t, line+" |")
		}
	}
}


func TestReplace(t *testing.T) {
	trie := NewTrie()
	trie.Add("苍井空")
	trie.Add("AV")
	trie.Add("日本AV")
	trie.Add("AV演员")

	text := "苍井空（あおい そら），日本AV演员兼电视、电影演员。苍井空AV女优是从xx出道"
	expect := "***（あおい そら），****演员兼电视、电影演员。*****女优是从xx出道"

	ok, words, newText := trie.Replace(text)

	t.Log("words:", words)
	t.Log("text:", newText)

	if !ok {
		t.Errorf("替换失败\n")
	}

	if newText != expect {
		t.Errorf("希望得到: %s\n实际得到: %s\n", expect, newText)
	}


	ok, _, _ = trie.Replace("完全和谐的文本")
	if ok {
		t.Errorf("替换失败\n")
	}
}

func TestDel(t *testing.T) {
	trie := NewTrie()
	trie.Add("苍井空")
	trie.Add("AV")
	trie.Add("日本AV")
	trie.Add("AV演员")

	node := trie.Root
	printTrie(node, t, " |")
	t.Log("-----")

	trie.Del("AV演员")
	node = trie.Root
	printTrie(node, t, " |")
	//
	//	text := "苍井空（あおい そら），日本AV演员兼电视、电影演员。苍井空AV女优是从xx出道"
	//	expect := "***（あおい そら），****演员兼电视、电影演员。*****女优是从xx出道"
	//
	//	_, _, newText := trie.Replace(text)
	//
	//	if newText != expect {
	//		t.Errorf("希望得到: %s\n实际得到: %s\n", expect, newText)
	//	}
	//
	//
	//	t.Log("删除一个敏感词")
	//
	//	trie.Del("AV")
	//
	//	text := "苍井空（あおい そら），日本AV演员兼电视、电影演员。苍井空AV女优是从xx出道"
	//	expect := "***（あおい そら），日本****兼电视、电影演员。***AV女优是从xx出道"
	//
	//	_, _, newText := trie.Replace(text)
	//	if newText != expect {
	//		t.Errorf("希望得到: %s\n实际得到: %s\n", expect, newText)
	//	}
}
