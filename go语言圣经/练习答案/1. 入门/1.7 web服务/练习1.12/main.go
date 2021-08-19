// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var count int
var mu sync.Mutex // 创建一个互斥锁

var palette = []color.Color{color.White, color.RGBA{G: 0xFF, A: 0xFF}}

const (
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler 匹配该规则的路径都将会进去handler函数
	http.HandleFunc("/lissajous", func(writer http.ResponseWriter, request *http.Request) {

		//var cycles = 5 // number of complete x oscillator revolutions
		const (
			cycles  = 5     // 常量不能是复杂的表达式生成的
			res     = 0.001 // angular resolution
			size    = 100   // image canvas covers [-size..+size]
			nframes = 64    // number of animation frames 动画帧数
			delay   = 8     // delay between frames in 10ms units
		)
		var iRCycles int
		rCycles := request.URL.Query().Get("cycles")
		if rCycles != "" {
			var err error
			iRCycles, err = strconv.Atoi(rCycles)
			if err != nil {
				_, _ = fmt.Fprintf(writer, "参数cycles有误,必须为字符串类型")
				return
			}
		} else {
			iRCycles = cycles
		}

		freq := rand.Float64() * 3.0 // relative frequency of y oscillator
		anim := gif.GIF{LoopCount: nframes}
		phase := 0.0 // phase difference
		for i := 0; i < nframes; i++ {
			rect := image.Rect(0, 0, 2*size+1, 2*size+1)
			img := image.NewPaletted(rect, palette)
			for t := 0.0; t < float64(iRCycles)*2*math.Pi; t += res {
				x := math.Sin(t)
				y := math.Sin(t*freq + phase)
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					blackIndex)
			}
			phase += 0.1
			anim.Delay = append(anim.Delay, delay)
			anim.Image = append(anim.Image, img)
		}
		_ = gif.EncodeAll(writer, &anim) // NOTE: ignoring encoding errors
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // 监听并开启服务
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	// 打印请求头信息
	_, _ = fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	_, _ = fmt.Fprintf(w, "Host = %q\n", r.Host)
	_, _ = fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	// 打印请求体信息
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		_, _ = fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	_, _ = fmt.Fprintf(w, "URL.Path = %q\n请求次数: %d \n", r.URL.Path, count)
}
