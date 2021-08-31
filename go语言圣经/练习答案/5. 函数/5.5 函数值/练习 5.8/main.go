package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {

	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("请求失败 %s", url)
	}
	defer resp.Body.Close()

	doc, _ := html.Parse(resp.Body)
	node := forEachNode(doc, startElement, endElement, "Popover8-toggle")
	fmt.Printf("%v", node)
}

func endElement(n *html.Node, id string) bool {
	return true
}

func startElement(n *html.Node, id string) bool {
	if n.Type != html.ElementNode {
		return true
	}
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return false
		}
	}
	return true
}

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用，返回true代表继续循环，否则中断循环
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node, s string) bool, s string) *html.Node {
	if pre != nil {
		if !pre(n, s) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, s)
	}
	if post != nil {
		if !post(n, s) {
			return n
		}
	}
	return n
}
