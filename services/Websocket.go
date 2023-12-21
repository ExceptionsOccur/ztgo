package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"ztgo/requests"
	"ztgo/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域请求
	},
}

var connections = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

func HandleWebSocketConnection(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Failed to upgrade to WebSocket:", err)
		// 网络错误
		utils.ZTResponseError(ctx, "500, 网络错误")
		return
	}
	connections[conn] = true
	go websocketLogic(conn)
	go websocketSendListener()
	go pushData()
}

// 消息推送
func pushData() {
	// 2s 刷新一次
	checkInterval := 2 * time.Second
	checkTicker := time.NewTicker(checkInterval)
	var res map[string]interface{}
	for {
		select {
		case <-checkTicker.C:
			// 成员变更推送
			res = make(map[string]interface{})
			diffJoin, diffLeave := requests.CheckMembers()
			// 没有更新，不做操作
			if len(diffJoin) == 0 && len(diffLeave) == 0 {
				continue
			}
			// 新增
			if len(diffJoin) != 0 {
				res["join"] = diffJoin
			}
			// 离开
			if len(diffLeave) != 0 {
				res["leave"] = diffLeave
			}
			// 有更新就推送
			if len(res) != 0 {
				res["eventcode"] = 2001
				sendData, _ := json.Marshal(res)
				WebsocketSend(sendData)
			}
		default:
			{
			}
		}
	}
}

// 给外部调用的websocket发送方法
func WebsocketSend(data []byte) {
	broadcast <- data
}

func websocketSendListener() {
	for {
		// 从广播通道获取消息
		message := <-broadcast

		// 将消息发送给所有连接的客户端
		for conn := range connections {
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				// 处理发送错误
				conn.Close()
				delete(connections, conn)
			}
		}
	}
}

func websocketLogic(conn *websocket.Conn) {
	pingInterval := 30 * time.Second
	pingMessage := []byte("pong")
	// 狗粮
	isDead := false
	// 设置心跳检测的定时器
	pingTicker := time.NewTicker(pingInterval)
outerLoop:
	for {
		select {
		case <-pingTicker.C:
			// 看门狗原理，不喂狗就退出
			if isDead {
				break outerLoop
			}
			// 下次退出标志，如果收到心跳消息则置否，等效为放狗粮
			isDead = true

		default:
			// 读取客户端发送的消息
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Failed to read message from WebSocket:", err)
				break outerLoop
			}
			// 如果是心跳消息则返回pong作为响应
			if string(message) == "ping" {
				err := conn.WriteMessage(websocket.TextMessage, pingMessage)
				if err != nil {
					fmt.Println("Failed to send ping message:", err)
					break outerLoop
				}
				// 喂狗
				isDead = false
			}
		}
	}
	conn.Close()
	delete(connections, conn)
	pingTicker.Stop()
}
