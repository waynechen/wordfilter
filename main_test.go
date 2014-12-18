package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type respData struct{
	Code  int `json:"code"`
	Error string `json:"error"`
	Mess  string `json:"mess"`
	Keywords []string `json:"keywords"`
	Text     string `json:"text"`
}

func Test_V1_Add_Fail(t *testing.T) {
	postUrl := "http://127.0.0.1:8080/v1/add"
	resp, err := http.PostForm(postUrl, url.Values{})
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected HTTP status is 200, but get %d\n", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	data := respData{}
	json.Unmarshal(body, &data)

	t.Log(string(body))
	if data.Code != 0 {
		t.Errorf("code: %d, error: %s, mess: %s", data.Code, data.Error, data.Mess)
	}
}

func Test_V1_Add_Success(t *testing.T) {
	postUrl := "http://127.0.0.1:8080/v1/add"
	resp, err := http.PostForm(postUrl, url.Values{"q": {"测试"}})
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected HTTP status is 200, but get %d\n", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	data := respData{}
	json.Unmarshal(body, &data)

	t.Log(string(body))
	if data.Code == 0 {
		t.Errorf("code: %d, error: %s, mess: %s", data.Code, data.Error, data.Mess)
	}
}

func Test_V1_Del_Success(t *testing.T){
	postUrl := "http://127.0.0.1:8080/v1/del"
	resp, err := http.PostForm(postUrl, url.Values{"q": {"测试,test02"}})
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected HTTP status is 200, but get %d\n", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	data := respData{}
	json.Unmarshal(body, &data)

	t.Log(string(body))
	if data.Code == 0 {
		t.Errorf("code: %d, error: %s, mess: %s", data.Code, data.Error, data.Mess)
	}
}

func Test_V1_Filter_Success(t *testing.T){
	postUrl := "http://127.0.0.1:8080/v1/filter"
	resp, err := http.PostForm(postUrl, url.Values{"q": {"我要测试敏感词"}})
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected HTTP status is 200, but get %d\n", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	data := respData{}
	json.Unmarshal(body, &data)

	if data.Code == 0 {
		t.Errorf("code: %d, error: %s, mess: %s", data.Code, data.Error, data.Mess)
	}

	t.Log(data.Keywords)
	t.Log(data.Text)
	if len(data.Keywords) == 0{
		t.Errorf("过滤失败")
	}
}
