package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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

	case "/v1/filter":
		filterWords(w, r)

	case "/v1/add": // 添加敏感词
		// TODO

	default:
		notFound(w)
	}
}

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func apiHelper(w http.ResponseWriter) {
	help := make(map[string]string)
	help["filter_url_get"] = "/v1/filter?q={text}"
	help["filter_url_post"] = "/v1/filter"
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
	http.ListenAndServe(ipAddr, &router{})
}

