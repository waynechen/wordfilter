# 敏感词过滤服务
---
基于词典的敏感词过滤程序

程序敏感词词典使用Trie树存储， 提供HTTP API访问

## 使用

```
go run main.go 127.0.0.1:8080
```
then visit http://127.0.0.1:8080/v1/filter?q=你的敏感词

## API

### 1.替换敏感词
输入一段文本，返回敏感词及敏感词替换为*号后的文本

* **Request:**  /v1/filter 
* **Request Method:** GET or POST 
* **Params**:
|Name| Required| Type | Default Value| Example | Desc. |
| -- | ------- | ---- | ------------ | ------- | ----  |
|q   | Yes     | string | | 中华人民共和国 | 文章内容 |
*  **Response:**
```
{
  "code": 1,
  "error": "", // 当code=0时，返回的错误消息
  "keywords": ["k1","k2"], //敏感词
  "text": "" //将敏感词替换为*号后的文本
}
```

### 2.添加敏感词
TOOD

## 词库说明
敏感词词库在 dictionary 目录里
每个敏感词独立一行。
