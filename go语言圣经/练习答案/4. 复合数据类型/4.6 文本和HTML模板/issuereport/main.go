package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Item
}

type Item struct {
	Title     string    `json:"title"`
	Number    int       `json:"number"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
	//Body   string   `json:"body"`
	Labels []*Label `json:"labels"`
}

type Label struct {
	Name string `json:"name"`
}

type User struct {
	Login string `json:"login"`
}

var url = "https://api.github.com/repos/spring-projects/spring-framework/issues?q=test"

// {{action}}
// 当前值“.”最初被初始化为调用模板时的参数，这里也就是 IssuesSearchResult 变量
func main() {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer resp.Body.Close()
	//all, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%s", all)
	items := make([]*Item, 0)
	if err = json.NewDecoder(resp.Body).Decode(&items); err != nil {
		log.Fatalln(err)
		return
	}

	//fmt.Println(items)

	// 数据准备
	data := IssuesSearchResult{len(items), items}
	// 模板准备
	const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`
	// 构建模板对象
	t, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 渲染
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// http://localhost:8000/issues
