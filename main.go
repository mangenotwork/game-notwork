package main

import (
	"encoding/json"
	"fmt"
	"github.com/mangenotwork/game-notwork/entity"
	"github.com/mangenotwork/game-notwork/jh"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upGrader = websocket.Upgrader{
		ReadBufferSize:  1024*100,
		WriteBufferSize: 65535,
		HandshakeTimeout: 5*time.Second,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	//游戏房间
	gameRoom sync.Map

)

func main(){


	r := gin.Default()
	r.StaticFS("/hilo", http.Dir("./hilo"))
	r.StaticFS("/resource", http.Dir("./resource"))
	r.StaticFS("/game_js", http.Dir("./game_js"))
	r.LoadHTMLGlob("game/*")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html","")
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html","")
	})
	// ws连接
	r.Any("/conn", func(ctx *gin.Context) {
		name := ctx.Query("name")
		log.Println("name = ", name)

		//if  _, ok := gameRoom.Load(name); ok {
		//	log.Println("已经登录")
		//	return
		//}

		if websocket.IsWebSocketUpgrade(ctx.Request) {
			log.Println("is websocket")

			conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
			if err != nil {
				log.Println("websocket upgrade error:", err)
				return
			}
			conn.WriteMessage(websocket.TextMessage, []byte("wxm.alming"))

			//存储conn
			gameRoom.Store(name, conn)

			welcome(name)

			go func(conn *websocket.Conn, name string) {
				for {
					err := conn.SetReadDeadline(time.Now().Add(3 * time.Hour)) // 超时min
					if err != nil {
						gameRoom.Delete(name)
						return
					}

					t, c, err := conn.ReadMessage()
					if err != nil && t == -1 {
						gameRoom.Delete(name)
						log.Println(err)
						return
					}

					if err != nil {
						log.Println(err)
						return
					}

					var msg entity.Msg
					if err := json.Unmarshal(c, &msg); err != nil {
						log.Println(err)
						sendErr(conn, err)
						continue
					}
					fmt.Println(t, msg)

					switch msg.Cmd {

					// OUT : 退出指令
					case "OUT":
						out(msg.User)
						gameRoom.Delete(name)

					// Game1 : game游戏指令
					case "Game1":
						game1(msg.User, msg.Data)

					// FP : 发牌指令
					case "FP":
						fapai(conn, name)

					// JH-NewRoot : 创建炸金花房间指令
					case "JH-NewRoot":
						// 创建并加入房间
						rootName := jh.NewJhRoot(msg.User, conn)
						// 欢迎进入
						jh.WelcomeJhRoot(name,rootName)
						// 下发房间列表
						JhRootList()
						// 下发房间信息
						jh.GetJhRootInfo(rootName)

					// JH-GetRootList : 获取炸金花房间列表指令
					case "JH-GetRootList":
						JhRootList()

					// JH-InRoom : 进入炸金花房间指令
					case "JH-InRoom":
						//进入房间
						jh.InJhRoot(msg.User, msg.Data, conn)
						// 欢迎进入
						jh.WelcomeJhRoot(name,msg.Data)

					// JH-RoomInfo : 获取炸金花房间信息指令
					case "JH-RoomInfo":
						jh.GetJhRootInfo(msg.Data)

					// JH-RoomOut : 退出炸金花房间指令
					case "JH-RoomOut":
						//退出
						jh.JhRoomUserOut(msg.User, msg.Data)
						jh.GetJhRootInfo(msg.Data)

					// JH-UserZB : 炸金花游戏准备指令
					case "JH-UserZB":
						//游戏准备
						isStart := jh.JhRoomUserZB(msg.User, msg.Data)
						jh.GetJhRootInfo(msg.Data)
						// 游戏开始
						if isStart {
							jh.JhStart(msg.User, msg.Data)
						}

					// JH-GameEnd : 结束当前游戏按钮
					case "JH-GameEnd":
						jh.JhRoomGameEnd(msg.User, msg.Data)

					// 红黑大战
					case "HongHei":
						//进入房间拉取当前信息

					}






				}
			}(conn, name)
		}

	})

	r.Run(":22222")
}

// 下行错误信息
func sendErr(conn *websocket.Conn, err error) {
	serversMsg := entity.Msg{
		Cmd: "servers",
		User: "",
		Data: "【错误】上传未知数据! err:"+err.Error(),
		Time: time.Now().Unix(),
	}
	serversResult, _ := json.Marshal(serversMsg)

	conn.WriteMessage(websocket.TextMessage, serversResult)
}

// 欢迎来宾
func welcome(name string) {
	serversMsg := entity.Msg{
		Cmd: "servers",
		User: "",
		Data: "【进入】 系统: 欢迎 "+name+" 进入房间",
		Time: time.Now().Unix(),
	}
	serversResult, _ := json.Marshal(serversMsg)

	gameRoom.Range(func(key, value interface{}) bool{
		if conn, ok := value.(*websocket.Conn); ok {
			conn.WriteMessage(websocket.TextMessage, serversResult)
		}
		return true
	})
}

// 退出房间
func out(name string) {

	serversMsg := entity.Msg{
		Cmd: "servers",
		User: "",
		Data: "【退出】 系统: "+name+" 退出房间",
		Time: time.Now().Unix(),
	}
	serversResult, _ := json.Marshal(serversMsg)

	gameRoom.Range(func(key, value interface{}) bool {
		if conn, ok := value.(*websocket.Conn); ok {
			conn.WriteMessage(websocket.TextMessage, serversResult)
		}
		return true
	})
}

//有些玩法 game1
func game1(name, data string) {

	msg := entity.Msg{
		Cmd: "Game1",
		User: name,
		Data: data,
		Time: time.Now().Unix(),
	}
	result ,_ := json.Marshal(msg)

	serversMsg := entity.Msg{
		Cmd: "servers",
		User: "",
		Data: "【游戏】 系统: "+name+" 点击了 "+ data,
		Time: time.Now().Unix(),
	}
	serversResult, _ := json.Marshal(serversMsg)

	gameRoom.Range(func(key, value interface{}) bool {
		if conn, ok := value.(*websocket.Conn); ok {
			conn.WriteMessage(websocket.TextMessage, serversResult)
			conn.WriteMessage(websocket.TextMessage, result)
		}
		return true
	})
}


// pkp 发牌
func fapai(conn *websocket.Conn, name string) {
	pai := jh.PaiMap
	pai1 := []string{}
	for j:=0; j<3; j++ {
		z := rand.Intn(53)
		if _, ok := pai[z]; ok &&  z != 0 {
			pai1 = append(pai1, fmt.Sprintf("%d",z))
			delete(pai, z)
		}else{
			j--
			continue
		}
	}
	log.Println(name," 的牌是 ", pai1)

	serversMsg := entity.Msg{
		Cmd: "FP",
		User: name,
		Data: strings.Join(pai1, ","),
		Time: time.Now().Unix(),
	}
	result ,_ := json.Marshal(serversMsg)
	conn.WriteMessage(websocket.TextMessage, result)
}

// 下发金花房间列表
func JhRootList() {
	data := ""
	for k, _ := range jh.JhRoot {
		data = data+k+","
	}
	msg := entity.Msg{
		Cmd: "JH-RootList",
		User: "",
		Data: data,
		Time: time.Now().Unix(),
	}
	result ,_ := json.Marshal(msg)
	gameRoom.Range(func(key, value interface{}) bool {
		if conn, ok := value.(*websocket.Conn); ok {
			conn.WriteMessage(websocket.TextMessage, result)
		}
		return true
	})

}

// 扎金花比大小
// 1. 3张一样 > 花色一样的顺子 > 花色一样 > 顺子 > 对子 > 普通牌
func jinhua(pai [][]string) map[string]int {

	//


	return nil
}

//func jihuaBP(paiA, paiB []int) {
//	is3 := false
//	isTHS := false
//	isTH := false
//	isSZ := false
//	isDZ := false
//
//	hs := []string{}
//
//
//	paiAList := strings.Split(paiA, ",")
//	for _,v := range paiAList{
//
//	}
//
//}

/*

1. 赌大小，压数字
2. 时时彩
3.



 */