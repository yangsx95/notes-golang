package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

// go run main.go repo:golang/go is:open json decoder
func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	// 按照日期分类
	aYearAgoDate := time.Now().AddDate(-1, 0, 0)
	aMonthAgoDate := time.Now().AddDate(0, -1, 0)
	dateItem := make(map[string][]Issue)
	dateItem["一年前"], dateItem["最近一年"], dateItem["最近一个月"] = make([]Issue, 0), make([]Issue, 0), make([]Issue, 0)

	for _, item := range result.Items {
		// 一年前
		if item.CreatedAt.Before(aYearAgoDate) {
			dateItem["一年前"] = append(dateItem["一年前"], *item)
		} else { // 最近一年
			dateItem["最近一年"] = append(dateItem["最近一年"], *item)
		}

		// 最近一个月
		if item.CreatedAt.After(aMonthAgoDate) {
			dateItem["最近一个月"] = append(dateItem["最近一个月"], *item)
		}
	}

	for s, issues := range dateItem {
		fmt.Printf("%s:\n", s)
		for _, item := range issues {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
		fmt.Println()
	}
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
