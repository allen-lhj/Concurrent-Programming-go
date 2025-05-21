package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handler websocket connection

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 确保在函数退出时关闭连接
	defer ws.Close()

	log.Println("客户端已连接")

	// 无限循环来监听来自客户端的消息
	for {
		// ReadMessage 从 WebSocket 连接中读取一条消息
		// 消息类型可以是 TextMessage 或 BinaryMessage
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			// 如果读取消息出错（比如客户端断开连接），记录错误并退出循环
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("错误: %v", err)
			}
			log.Println("客户端断开连接。")
			break
		}

		// 打印收到的消息
		log.Printf("收到客户端消息: %s", string(p))

		// 准备要回复的消息
		responseMessage := "服务器收到： " + string(p)

		// WriteMessage 将消息写回 WebSocket 连接
		if err := ws.WriteMessage(messageType, []byte(responseMessage)); err != nil {
			log.Println(err) // 如果写入失败，记录错误并退出循环
			break
		}
		log.Printf("已发送回复: %s", responseMessage)
	}
}

func main() {
	// 设置路由，当访问 /ws 时，调用 handleConnections 函数
	http.HandleFunc("/ws", handleConnections)

	// 启动 HTTP 服务器，监听在 8080 端口
	log.Println("WebSocket 服务器启动，监听端口 :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
