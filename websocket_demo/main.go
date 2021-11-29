package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		var (
			conn *websocket.Conn
			err  error
			data []byte
		)

		// 创建协议升级器
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool { // 允许跨域
				return true
			},
		}
		// 升级协议为websocket，并获取websocket连接对象
		if conn, err = upgrader.Upgrade(writer, request, nil); err != nil {
			log.Println(err)
			return
		}

		// 循环读取客户端信息
		for {
			// 参数1 为messageType
			// 参数2 为数据
			if _, data, err = conn.ReadMessage(); err != nil {
				log.Println(err)
				goto ERR
			}
			if err = conn.WriteMessage(websocket.TextMessage, data); err != nil { // 发的什么就返回什么
				log.Println(err)
				goto ERR
			}
		}
	ERR:
		conn.Close()
	})
	_ = http.ListenAndServe("0.0.0.0:7777", nil)
}
