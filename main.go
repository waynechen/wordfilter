package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	_ "dooland/wordfilter/dict"
	"dooland/wordfilter/trie"
	"os"
)

type router struct{

}

func (this *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path{
	case "/":
		apiHelper(w)

	case "/v1/filter": // 查找过滤
		filterWords(w, r)

	case "/v1/add": // 添加敏感词
		addWords(w, r)

	case "/v1/del": // 删除敏感词
		delWords(w, r)

	case "/v1/words": // 显示所有敏感词
		viewWords(w, r)

	default:
		notFound(w)
	}
}

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func apiHelper(w http.ResponseWriter) {
	help := make(map[string]string)
	help["filter_url (GET)"] = "/v1/filter?q={text}"
	help["filter_url (POST)"] = "/v1/filter"
	help["add_words (POST)"] = "/v1/add"
	help["delete_words (POST)"] = "/v1/del"
	serveJson(w, help)
}

func filterWords(w http.ResponseWriter, r *http.Request) {
	paramName := "q"

	type resp struct{
		Code     int    `json:"code"`
		Error    string `json:"error"`
		Keywords []string `json:"keywords"`
		Text     string `json:"text"`
	}

	text := ""
	if r.Method == "GET" {
		params, err := url.ParseQuery(r.URL.RawQuery)
		if err == nil {
			if q, ok := params[paramName]; ok {
				text = q[0]
			}
		}else {
			fmt.Println(err)
		}

	}else if r.Method == "POST" {
		text = r.FormValue(paramName)
	}

	res := resp{
		Keywords: []string{},
	}

	if text != "" {
		res.Code = 1
		ok, keyword, newText := trie.Singleton().Replace(text)
		if ok {
			res.Keywords = keyword
			res.Text = newText
		}
	}else {
		res.Code = 0
		res.Error = "参数"+paramName+"不能为空"
	}
	serveJson(w, res)
}

func addWords(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	if r.Method == "POST" {
		q := r.FormValue("q")
		if q == "" {
			resp["code"] = 0
			resp["error"] = "参数q不能为空"
		}else {
			i := 0
			words := strings.Split(q, ",")
			for _, s := range words {
				trie.Singleton().Add(strings.Trim(s, " "))
				i ++
			}

			resp["code"] = 1
			resp["mess"] = fmt.Sprintf("共添加了%d个敏感词", i)
		}
	}else {
		resp["code"] = 0
		resp["error"] = "只允许POST方式"
	}
	serveJson(w, resp)
}

func delWords(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	if r.Method == "POST" {
		q := r.FormValue("q")
		if q == "" {
			resp["code"] = 0
			resp["error"] = "参数q不能为空"
		}else {
			i := 0
			words := strings.Split(q, ",")
			for _, s := range words {
				trie.Singleton().Del(strings.Trim(s, " "))
				i ++
			}

			resp["code"] = 1
			resp["mess"] = fmt.Sprintf("共删除了%d个敏感词", i)
		}
	}else {
		resp["code"] = 0
		resp["error"] = "只允许POST方式"
	}
	serveJson(w, resp)
}

func viewWords(w http.ResponseWriter, r *http.Request) {
	words := trie.Singleton().ReadAll()
	str := strings.Join(words,"\n")
	w.Header().Set("Server", "DOOLAND")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte(str))
}

func serveJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Server", "DOOLAND")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	content, err := json.Marshal(data)
	if err == nil {
		w.Write(content)
	}else {
		w.Write([]byte(`{"code":0, "error":"解析JSON出错"}`))
	}
}

func main() {
	ipAddr := ":8080"
	if len(os.Args) > 1 {
		ipAddr = os.Args[1]
	}

	t := time.Now().Local().Format("2006-01-02 15:04:05 -0700")
	fmt.Printf("%s Listen %s\n", t, ipAddr)
	http.ListenAndServe(ipAddr, &router{})
}

