package jh

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mangenotwork/game-notwork/entity"
	"log"
	"math/rand"
	"strings"
	"time"
)

//金花房间
var JhRoot map[string]*JhRoomData

// 金花房间数据结构
type JhRoomData struct {
	Users map[string]*websocket.Conn //房间用户的连接
	IsZb map[string]bool // 是否准备

	//游戏流程玩法
	UserList []string //玩家顺序  庄家在第一个
	NowUser string //当前操作玩家
	UserLos map[string]bool // 玩家是否失败
	UserWin int //剩余未失败玩家数量
	UserPai map[string]*JhPai //玩家的牌
	UserFen map[string]int //玩家上的分
	GameChan chan JhChan // 游戏流程通道
	IsEnd bool // 是否是结束的状态
}

type JhChan struct {
	Cmd string
	User string
	DataStr string
	DataInt int
}

// 创建房间
func NewJhRoot(user string, conn *websocket.Conn) string {
	if JhRoot == nil {
		JhRoot = make(map[string]*JhRoomData)
	}
	name := fmt.Sprintf("root_%d", len(JhRoot))

	JhRoot[name] = &JhRoomData{
		Users: make(map[string]*websocket.Conn),
		IsZb: make(map[string]bool),

		// 初始话游戏玩法
		UserList: make([]string,0),
		UserWin: 0,
		UserLos: make(map[string]bool), // 发牌的时候才写值
		UserPai: make(map[string]*JhPai), // 发牌的时候才写值
		UserFen: make(map[string]int), // 发牌的时候才写值
		GameChan: make(chan JhChan), // 游戏流程通道
		IsEnd: true,
	}

	JhRoot[name].Users[user] = conn
	JhRoot[name].IsZb[user] = false

	//玩家顺序
	JhRoot[name].UserList = append(JhRoot[name].UserList, user)

	return name
}

// 用户加入房间
func InJhRoot(user, rootName string, conn *websocket.Conn) {
	JhRoot[rootName].Users[user] = conn
	JhRoot[rootName].IsZb[user] = false

	//玩家顺序
	JhRoot[rootName].UserList = append(JhRoot[rootName].UserList, user)
}

// 广播系统欢迎进入房间的用户
func WelcomeJhRoot(name, rootName string) {
	for _, v := range JhRoot[rootName].Users {
		serversMsg := entity.Msg{
			Cmd: "servers",
			User: "",
			Data: "【进入】 系统: 欢迎 "+name+" 进入金花房间 "+ rootName,
			Time: time.Now().Unix(),
		}
		result ,_ := json.Marshal(serversMsg)
		v.WriteMessage(websocket.TextMessage, result)
	}
}

// 返回金花游戏房间数据结构
type JHRoomInfo struct {
	RoomId string `json:"id"`
	Title string `json:"title"` // 房间名称
	Users map[string]string `json:"users"` // 用户准备情况
}

// 广播 房间的信息
func GetJhRootInfo(rootName string) {
	info := &JHRoomInfo{
		RoomId: rootName,
		Title: "炸金花游戏房间: "+rootName,
		Users: make(map[string]string),
	}
	for k, v := range JhRoot[rootName].IsZb{
		if v {
			info.Users[k] = "已经准备"
		}else{
			info.Users[k] = "未准备"
		}
	}
	infoJson, _ := json.Marshal(info)
	for _, v := range JhRoot[rootName].Users {
		serversMsg := entity.Msg{
			Cmd: "JH-RoomInfo",
			User: "",
			Data: string(infoJson),
			Time: time.Now().Unix(),
		}
		result ,_ := json.Marshal(serversMsg)
		v.WriteMessage(websocket.TextMessage, result)
	}
}

// 广播用户退出
func JhRoomUserOut(user, roomName string) {
	for _, v := range JhRoot[roomName].Users {
		serversMsg := entity.Msg{
			Cmd: "servers",
			User: "",
			Data: "【进入】 系统:  "+user+" 退出游戏 ",
			Time: time.Now().Unix(),
		}
		result ,_ := json.Marshal(serversMsg)
		v.WriteMessage(websocket.TextMessage, result)
	}

	userList := JhRoot[roomName].UserList
	for k,v := range userList {
		if v == user {
			JhRoot[roomName].UserList = append(JhRoot[roomName].UserList[:k], JhRoot[roomName].UserList[k+1:]...)
		}
	}

	delete(JhRoot[roomName].Users, user)
	delete(JhRoot[roomName].IsZb, user)

}

// 用户准备并广播
func JhRoomUserZB(user, roomName string) bool {
	JhRoot[roomName].IsZb[user] = true

	// 触发 验证玩家是否都准备
	for _, v := range JhRoot[roomName].IsZb{
		if !v {
			return false
		}
	}

	for _, v := range JhRoot[roomName].Users {
		serversMsg := entity.Msg{
			Cmd: "servers",
			User: "",
			Data: "【进入】 系统:  "+user+" 已经准备 ",
			Time: time.Now().Unix(),
		}
		result ,_ := json.Marshal(serversMsg)
		v.WriteMessage(websocket.TextMessage, result)
	}

	return true
}


// 炸金花游戏开始
func JhStart(user, roomName string){
	send(roomName,"【进入】 系统:  游戏即将开始!!!  ")

	// 开始游戏初始化
	JhRoot[roomName].IsEnd = false
	JhRoot[roomName].UserPai = make(map[string]*JhPai)
	JhRoot[roomName].UserLos = make(map[string]bool)
	JhRoot[roomName].UserWin = 0
	JhRoot[roomName].UserFen = make(map[string]int)
	// 随机指定用户
	JhRoot[roomName].NowUser = JhRoot[roomName].UserList[0]

	// 倒计时
	for i:=3; i>0; i-- {
		time.Sleep(1*time.Second)
		send(roomName,fmt.Sprintf("【进入】 系统:  游戏开始倒计时 %d s", i))
	}

	// 即将发牌
	send(roomName,"【进入】 系统:  请坐稳～ 即将发牌 ")

	// JH-FP : 炸金花发牌
	pai := PaiMap
	for k, v := range JhRoot[roomName].Users {
		paiInt := []int{}
		pai1 := []string{}
		for j:=0; j<3; j++ {
			z := rand.Intn(53)
			if _, ok := pai[z]; ok &&  z != 0 {
				pai1 = append(pai1, fmt.Sprintf("%d",z))
				paiInt = append(paiInt, z)
				delete(pai, z)
			}else{
				j--
				continue
			}
		}
		log.Println(k," 的牌是 ", pai1)

		// 存储用户的牌
		JhRoot[roomName].UserPai[k], _ = GetJHPai(paiInt,k)
		JhRoot[roomName].UserLos[k] = false
		JhRoot[roomName].UserWin++
		JhRoot[roomName].UserFen[k] = 50 // 底分

		serversMsg := entity.Msg{
			Cmd: "JH-FP",
			User: k,
			Data: strings.Join(pai1, ","),
			Time: time.Now().Unix(),
		}
		result ,_ := json.Marshal(serversMsg)
		v.WriteMessage(websocket.TextMessage, result)

	}

	log.Println("UserList = ", JhRoot[roomName].UserList)

	// 游戏时间轮，用户排队操作
	go func(){
		for {
			for _, v := range JhRoot[roomName].UserList{

				if JhRoot[roomName].IsEnd {
					return
				}

				JhRoot[roomName].NowUser = v
				log.Println("v = ", v)
				time.Sleep(10*time.Second)
				JhRoot[roomName].GameChan <- JhChan{
					Cmd : "qh",
					User : v,
					DataStr : "",
					DataInt : 0,
				}

				////启动定时 -- 时间到了游戏结束
				//timer1 := time.NewTicker(3 * time.Second)
				//select {
				//case <-timer1.C:
				//	send(roomName,"【进入】 系统: "+ v +" 操作结束。")
				//}
			}
		}
	}()

	// 游戏主流程
	go func(){
		for {
			select {
				case chanData := <- JhRoot[roomName].GameChan:
					log.Println(chanData, JhRoot[roomName].NowUser)

					// 游戏结束
					if chanData.Cmd == "end" {
						if chanData.User == JhRoot[roomName].NowUser{
							send(roomName,"【进入】 系统: "+chanData.DataStr+"。")
							for k, v := range JhRoot[roomName].Users {
								serversMsg := entity.Msg{
									Cmd: "servers",
									User: "",
									Data: k + "的牌是: " + JhRoot[roomName].UserPai[k].PxName,
									Time: time.Now().Unix(),
								}
								result ,_ := json.Marshal(serversMsg)
								v.WriteMessage(websocket.TextMessage, result)
							}
							send(roomName,"【进入】 系统: "+chanData.User+"点击结束，本轮游戏结束。")
							JhRoot[roomName].IsEnd = true
							return
						} else {
							userSend(JhRoot[roomName].Users[chanData.User], "【进入】 系统: 没有到你的操作时间.")
						}

					}

					// 游戏 - 玩家操作时间结束
					if chanData.Cmd == "qh"{
						send(roomName,"【进入】 系统: "+ chanData.User +" 操作结束。")
					}

					//

				}
		}
	}()


}


func send(roomName, connet string){
	for _, v := range JhRoot[roomName].Users {
		serversMsg := entity.Msg{
			Cmd: "servers",
			User: "",
			Data: connet,
			Time: time.Now().Unix(),
		}
		result ,_ := json.Marshal(serversMsg)
		v.WriteMessage(websocket.TextMessage, result)
	}
}

func userSend(conn *websocket.Conn, connet string){
	serversMsg := entity.Msg{
		Cmd: "servers",
		User: "",
		Data: connet,
		Time: time.Now().Unix(),
	}
	result ,_ := json.Marshal(serversMsg)
	conn.WriteMessage(websocket.TextMessage, result)
}

// 游戏结束按钮
func JhRoomGameEnd(user, roomName string){
	//JhRoot[roomName].GameChan <- JhChan{
	//	Cmd : "end",
	//	User : user,
	//	DataStr : "",
	//	DataInt : 0,
	//}

	JhBP(user, roomName)
}

// 比牌按钮
func JhBP(user, roomName string) {
	//比牌
	user2 := ""
	for _,v := range JhRoot[roomName].UserList {
		if v != user{
			user2 = v
			break
		}
	}
	jgdata := ""
	jg := JHPaiPK(JhRoot[roomName].UserPai[user], JhRoot[roomName].UserPai[user2])
	log.Println("比牌结果: ", jg)
	if jg == nil {
		jgdata = "平局"
	}else{
		jgdata = jg.User + "胜利"
	}

	JhRoot[roomName].GameChan <- JhChan{
		Cmd : "end",
		User : user,
		DataStr : jgdata,
		DataInt : 0,
	}
}

/*


 */