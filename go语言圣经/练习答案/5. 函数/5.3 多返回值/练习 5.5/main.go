package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
			continue
		}
		fmt.Printf("图片数量：%d， 单词数量 %d", images, words)
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" { // 如果是html 元素，并且标签为a标签
		for _, a := range n.Attr { // 遍历该节点属性，找到他的href属性
			if a.Key == "href" {
				links = append(links, a.Val) // 放入到links中
			}
		}
	}
	// 如果当前元素有子元素，获取第一个子元素，然后对该子元素的兄弟元素遍历，达到遍历所有子元素的效果
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}


func countWordsAndImages(n *html.Node) (words, images int) {
	//fmt.Printf("%v, %v", n.Type, n.Data)
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
		return
	}
	if n.Type == html.TextNode {
		words += len(strings.Split(n.Data, " "))
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}
