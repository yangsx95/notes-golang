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
	forEachNode(doc, startElement, endElement)
}

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		s := fmt.Sprintf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			s += " " + a.Key + "=\"" + a.Val + "\""
		}
		if n.FirstChild != nil {
			s += ">"
			if  n.FirstChild.Type != html.TextNode {
				s += "\n"
			}
		}
		depth++
		fmt.Print(s)
	}
	if n.Type == html.TextNode {
		fmt.Printf("%*s%s", depth*2, "", n.Data)
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild == nil {
			fmt.Printf("/>\n")
			return
		}
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}

}
