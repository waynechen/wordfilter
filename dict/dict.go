package dict

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/huayuego/wordfilter/trie"
)

func init() {
	LoadDict()
}

// 从字典中装入敏感词库
func LoadDict() {
	load("./dictionary/add", "add")
	load("./dictionary/del", "del")
}

func load(path , op string) {

	var loadAllDictWalk filepath.WalkFunc = func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		initTrie(path, op)

		return nil
	}

	err := filepath.Walk(path, loadAllDictWalk)
	if err != nil {
		panic(err)
	}
}

func initTrie(path, op string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("fail to open file %s %s", path, err.Error())
		return
	}

	defer f.Close()

	fmt.Printf("%s Load dict: %s\n", time.Now().Local().Format("2006-01-02 15:04:05 -0700"), path)

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

			if "add" == op {
				trie.Singleton().Add(s)

			}else if "del" == op {
				trie.Singleton().Del(s)
			}
		}
	}

	return
}
