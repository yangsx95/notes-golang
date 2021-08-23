package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
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
	Labels  []*Label `json:"labels"`
	State   string   `json:"state"`
	HTMLURL string   `json:"html_url"`
}

type Label struct {
	Name string `json:"name"`
}

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

var url = "https://api.github.com/repos/spring-projects/spring-framework/issues?q=test"

// 模板准备
const templ = `<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>`

var cache *IssuesSearchResult
var t *template.Template

func main() {
	http.HandleFunc("/issues", handler)                   // each request calls handler 匹配该规则的路径都将会进去handler函数
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // 监听并开启服务
}

func handler(w http.ResponseWriter, r *http.Request) {
	if cache == nil {
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
		cache = &IssuesSearchResult{len(items), items}
		// 构建模板对象
		t, err = template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}

	// 渲染
	err := t.Execute(w, cache)
	if err != nil {
		log.Fatalln(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
