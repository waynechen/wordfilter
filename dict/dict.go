package dict

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"dooland/wordfilter/trie"
)

func init() {
	LoadDict()
}

// 从字典中装入敏感词库
func LoadDict() {
	getFileLists("./dictionary")
}

func getFileLists(path string) {

	var loadAllDictWalk filepath.WalkFunc = func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		loadAndAddToTrie(path)

		return nil
	}

	err := filepath.Walk(path, loadAllDictWalk)
	if err != nil {
		panic(err)
	}
}

func loadAndAddToTrie(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("fail to open file %s %s", path, err.Error())
		return
	}

	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, isPrefix, e := buf.ReadLine()
		if e != nil {
			if e != io.EOF {
				err = e
			}
			break
		}
		if isPrefix {
			continue
		}

		if word := strings.TrimSpace(string(line)); word != "" {
			tmp := strings.Split(word, " ")
			s := strings.Trim(tmp[0], " ")
			trie.Singleton().Add(s)
		}
	}

	return
}
