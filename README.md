#敏感词过滤服务
---

基于词典的敏感词过滤程序

程序敏感词词典使用Trie树存储， 提供HTTP API访问

## 使用

```
go run main.go 127.0.0.1:8080
```

then visit http://127.0.0.1:8080/v1/filter?q=文本内容

## API

### 1.查找替换敏感词
输入一段文本，返回敏感词及敏感词替换为*号后的文本

* **Request:**  /v1/filter 
* **Request Method:** GET or POST 
* **Params**:

| Name | Type | Requried | Example | Desc. |
| ---- | ---- | -------- | ------- | ----- |
| q | string | yes | | 需要检查的文本内容 |

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

添加一组敏感词

* **Request:**  /v1/add 
* **Request Method:** POST 
* **Params**:

| Name | Type | Requried | Example | Desc. |
| ---- | ---- | -------- | ------- | ----- |
| q    | string | yes  | 你大爷,走私 | 敏感词，多个之间与逗号相隔 |

*  **Response:**
```
{
  "code": 1,
  "error": "", // 当code=0时，返回的错误消息
}
```

### 3.删除敏感词

删除一组敏感词

* **Request:**  /v1/del 
* **Request Method:** POST 
* **Params**:

| Name | Type | Requried | Example | Desc. |
| ---- | ---- | -------- | ------- | ----- |
| q    | string | yes  | 你大爷,走私 | 敏感词，多个之间与逗号相隔 |

*  **Response:**
```
{
  "code": 1,
  "error": "", // 当code=0时，返回的错误消息
}
```

### 4.查看所有敏感词

* **Request:**  /v1/words 
* **Request Method:** GET
* **Response:**
```
陪睡
陪伴服务
陪聊女
陪聊小姐
```

## 词库说明
敏感词词库在 dictionary 目录里
每个敏感词独立一行。

- dictionary/add 默认载入的词典

- dictionary/del 默认输入的词典中需要删除的字词
  如add中有”情色“, 在del中也有”情色“, 则表示排除掉了”情色“这个词,不会过滤这个词了

