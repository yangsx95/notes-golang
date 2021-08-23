package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Issue struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Labels []string `json:"labels"`
}

func main() {

	url := "https://api.github.com/repos/xf616510229/hadoop-study/issues"

	issue := Issue{"这是来自GO学习示例的Issue", "我是内容", []string{"test"}}

	issueJson, err := json.Marshal(issue)

	r := bytes.NewReader(issueJson)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, r)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token *************")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
