// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// 改造，输入地址，下载对应地址的html文件并分析
	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	//fmt.Printf("%v = %v\n", n.Type, n.Data)
	if n.Type == html.ElementNode && n.Data == "script" {
		return links
	}

	if n.Type == html.TextNode {
		d := strings.TrimSpace(n.Data)
		if d != "" {
			fmt.Println(d)
		}
	}

	// 如果当前元素有子元素，获取第一个子元素，然后对该子元素的兄弟元素遍历，达到遍历所有子元素的效果
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
