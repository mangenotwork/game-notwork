package entity

// 上下行数据结构
type Msg struct {
	Cmd string `json:"cmd"` // 指令
	User string `json:"user"` // 用户标识
	Data string `json:"data"` // 数据
	Time int64 `json:"time"` // 时间
}