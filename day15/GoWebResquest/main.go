package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 1，向网络地址发送请求，并获取一段json格式数据
	urlPath := "https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=6915066165949631752"
	req, _ := http.NewRequest("GET", urlPath, nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.26")
	client := &http.Client{}
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body) // 发送网络请求，读取字节
	res.Body.Close()
	// fmt.Println(body)
	content := string(body)
	fmt.Println("接收返回值内容：", content)

	// 2，将json数据反序列化为GO数据类型
	var responseObject map[string]interface{}
	json.Unmarshal(body, &responseObject)
	fmt.Println(responseObject)
}
