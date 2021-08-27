// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	var emap map[string]int = make(map[string]int)
	for _, link := range visit(nil, &emap, doc) {
		fmt.Println(link)
	}

	fmt.Println(emap)
}

func visit(links []string, emap *map[string]int, n *html.Node) []string {
	// 记录标签出现次数
	if n.Type == html.ElementNode {
		(*emap)[n.Data]++
	}

	if n.Type == html.ElementNode && n.Data == "a" { // 如果是html 元素，并且标签为a标签
		for _, a := range n.Attr { // 遍历改节点，找到他的href属性
			if a.Key == "href" {
				links = append(links, a.Val) // 放入到links中
			}
		}
	}
	// 如果当前元素有子元素，获取第一个子元素，然后对该子元素的兄弟元素遍历，达到遍历所有子元素的效果
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, emap, c)
	}
	return links
}
