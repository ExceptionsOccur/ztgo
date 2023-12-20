package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域请求
	},
}

func HandleWebSocketConnection(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Failed to upgrade to WebSocket:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	go websocketLogic(conn)
}

func websocketLogic(conn *websocket.Conn) {
	pingInterval := 20 * time.Second
	pingMessage := []byte("pong")

	isDead := false

	// 设置心跳检测的定时器
	pingTicker := time.NewTicker(pingInterval)
outerLoop:
	for {
		select {
		case <-pingTicker.C:

			if isDead {
				break outerLoop
			}
			isDead = true

		default:
			// 读取客户端发送的消息
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Failed to read message from WebSocket:", err)
				break outerLoop
			}

			if string(message) == "ping" {
				err := conn.WriteMessage(websocket.TextMessage, pingMessage)
				if err != nil {
					fmt.Println("Failed to send ping message:", err)
					break outerLoop
				}
				isDead = false
			}

			fmt.Println("Received message:", string(message))

			// 在这里处理收到的消息，并发送响应消息
			err = conn.WriteMessage(websocket.TextMessage, []byte("Server received your message"))
			if err != nil {
				fmt.Println("Failed to send response message:", err)
				break outerLoop
			}
		}
	}

	conn.Close()
	pingTicker.Stop()
}
