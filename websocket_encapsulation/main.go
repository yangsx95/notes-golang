package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

/*
封装websocket，提供两个方法：
- ReadMessage 方法，负责将消息从 in channel中读取
- SendMessage 方法，负责将消息投递到 out channel

内部原理：
- 启动读协程，循环读取Websocket，将消息投递到 in channel
- 启动写协程，循环读取out channel，并将消息传递给websocket发送

技术难点：
- 内核瓶颈，100万在线 * 10条/s = 1000万条/s，而linux内核发送tcp的极限包频 ≈ 100万/秒
  解决方案：减少网络小包的发送（kafka也有这种机制），将N条消息合并为1条消息，合并后每秒推送次数大致等于在线连接数
- 锁瓶颈，需要在系统中维护已经登录的用户长链接集合(100w在线)
         推送消息要对整个集合遍历，循环发送，耗时极长
         推送期间，客户端仍旧可以正常上下线，所以集合需要上锁，防止并发操作出现问题
  解决方案：将长链接打散到多个集合中，每个集合有自己的锁，多线程并发推送集合，避免锁竞争
          将互斥锁使用读写锁替代，减少锁竞争
- CPU瓶颈，浏览器与服务端采取json格式通讯，json编码非常耗费CPU资源
  解决方案：json编码前置，1次编码，100万次推送

单机瓶颈：
- 内存瓶颈，长链接维护
- CPU瓶颈，消息推送瞬间消耗大量CPU
- 带宽瓶颈，瞬间可达到400-600M（主要瓶颈）

*/

func main() {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		var (
			wsConn *websocket.Conn
			conn   *Connection
			err    error
			data   []byte
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
		if wsConn, err = upgrader.Upgrade(writer, request, nil); err != nil {
			log.Println(err)
			return
		}

		// 封装conn
		if conn, err = InitConnection(wsConn); err != nil {
			log.Println(err)
			return
		}

		// 启动协程，循环向客户端推送消息
		go func() {
			var (
				err error
			)
			for {
				if err = conn.SendMessage([]byte("heartbeat")); err != nil {
					return
				}
				time.Sleep(1 * time.Second)
			}
		}()

		// 循环读取客户端信息
		for {
			// 参数1 为messageType
			// 参数2 为数据
			if data, err = conn.ReadMessage(); err != nil {
				log.Println(err)
				goto ERR
			}
			if err = conn.SendMessage(data); err != nil { // 发的什么就返回什么
				log.Println(err)
				goto ERR
			}
		}
	ERR:
		conn.Close()
	})
	_ = http.ListenAndServe("0.0.0.0:7777", nil)
}
