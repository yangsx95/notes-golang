package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string // 只有导出的结构体才会被编码，必须使用大写字母开头
	Year   int    `json:"released"`        // 结构体成员Tag。  类似注解，代表 Year 这个字段，转换为json时，使用 released
	Color  bool   `json:"color,omitempty"` // omitempty表示，如果该成员值为零值，则忽略该json字段
	Actors []string
}

func main() {
	var movies []Movie = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	// 序列化
	//jsonBytes, err := json.Marshal(movies)
	jsonBytes, err := json.MarshalIndent(movies, "", "    ") //格式化
	if err != nil {
		log.Fatalf("解析失败, %v", err)
		return
	}

	fmt.Printf("%s\n\n", jsonBytes)

	// 反序列化
	newMovie := make([]Movie, 0)
	err = json.Unmarshal(jsonBytes, &newMovie)
	if err != nil {
		return
	}
	fmt.Printf("%#v\n\n", newMovie)
}
