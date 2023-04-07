package jh

import (
	"errors"
	"fmt"
	"sort"
)

/*
	游戏-金花

	三同 > 顺金 > 金花 > 顺子 > 对子 > 普通牌
	比牌: A最大

 */

// 金花牌
type JhPai struct {
	Value []int
	Px int
	PxName string
	Pz []int
	Bz []int
	Text []string
	Dz int

	// 业务相关
	User string // 谁的牌
}

// 牌
var PaiMap = map[int]string{
	1:"A",	2:"2",	3:"3",	4:"4",	5:"5",	6:"6",	7:"7",	8:"8",	9:"9",	10:"10",	11:"J",	12:"Q",	13:"K",
	14:"A", 15:"2",	16:"3",	17:"4",	18:"5",	19:"6",	20:"7",	21:"8",	22:"9",	23:"10",	24:"J",	25:"Q",	26:"K",
	27:"A",	28:"2", 29:"3", 30:"4",	31:"5",	32:"6",	33:"7",	34:"8",	35:"9",	36:"10",	37:"J",	38:"Q",	39:"K",
	40:"A",	41:"2",	42:"3",	43:"4",	44:"5",	45:"6",	46:"7",	47:"8",	48:"9",	49:"10",	50:"J",	51:"Q",	52:"K",
}

// 牌大小
var PaiValue = map[string]int{
	"A":1, "2":2, "3":3, "4":4, "5":5, "6":6, "7":7, "8":8, "9":9, "10":10, "J":11, "Q":12, "K":13,
}

// 金花牌等级
var PaiLv = map[int]string{
	0:"普通", 1:"对子", 2:"顺子", 3:"同花", 4:"同花顺", 5:"三同",
}

// 计算牌
func GetJHPai(pai []int, user string) (*JhPai, error){
	if len(pai) != 3 {
		return nil, errors.New("牌数量错误!")
	}
	//初始值
	px := 0 // 牌等级
	dz := 0 // 对子的点数
	isTh := false // 是否是同花
	isSz := false // 是否是顺子
	// 对牌的值进行排序
	sort.Ints(pai)
	// 转换牌型
	p1 := []string{PaiMap[pai[0]], PaiMap[pai[1]], PaiMap[pai[2]]}
	// 转换牌的类型
	p2 := []string{}
	for _, v := range pai{
		if v > 0 && v < 14 {
			p2 = append(p2, "红桃")
		}
		if v > 13 && v< 27 {
			p2 = append(p2, "黑桃")
		}
		if v > 26 && v < 40 {
			p2 = append(p2, "方块")
		}
		if v > 39 && v < 53 {
			p2 = append(p2, "梅花")
		}
	}
	// 判断是否是同花
	if p2[0] == p2[1] && p2[1] == p2[2] {
		isTh = true
	}
	// 转换牌的点数
	p3 := []int{PaiValue[p1[0]], PaiValue[p1[1]], PaiValue[p1[2]]}
	// 判断是否顺子
	if p3[2]-p3[1] == 1 && p3[1]-p3[0] == 1 {
		isSz = true
	}
	// 判断 -> 对子, 三同
	if p3[0] == p3[1] && p3[1] ==p3[2] {
		//log.Println("三同 : ", pai, p3)
		px = 5
	} else if p3[0] == p3[1] || p3[0] == p3[2] || p3[1] == p3[2]{
		px = 1
		if p3[0] == p3[1] || p3[0] == p3[2]{
			dz = p3[0]
		}
		if p3[1] == p3[2] {
			dz = p3[1]
		}
		// A最大
		if dz == 1 {
			dz = 21
		}
	}
	// 是否是同花
	if isTh && !isSz{
		px = 3
	}
	// 是否是顺子
	if !isTh && isSz{
		px = 2
	}
	// 是否是同花顺
	if isTh && isSz{
		px = 4
	}
	// 牌的文字
	pTxt := []string{ fmt.Sprintf("%s-%s", p2[0],p1[0]),
		fmt.Sprintf("%s-%s", p2[1],p1[1]),
		fmt.Sprintf("%s-%s", p2[2],p1[2]),
	}
	// 金花牌型对象
	r := &JhPai{
		Value: pai,
		Px : px,
		PxName: PaiLv[px],
		Pz : p3,
		Text : pTxt,
		Dz: dz,
		User: user,
	}
	// 得到对比大小用的值
	for _, v := range p3 {
		if v == 1{
			r.Bz = append(r.Bz,21) // A最大
			continue
		}
		r.Bz = append(r.Bz,v)
	}
	sort.Slice(r.Bz, func(i, j int) bool {
		return r.Bz[i] > r.Bz[j]
	})
	return r, nil
}

// 金花牌的大小对比
// 返回胜利者， 如果平局返回nil
func JHPaiPK(A, B *JhPai) *JhPai {
	// 比牌形
	if A.Px > B.Px {
		return A
	}else if A.Px < B.Px {
		return B
	}else{
		//牌行一样
		//1. 对子比
		if A.Px == 1 && A.Dz > B.Dz{
			return A
		}else if A.Px == 1 && A.Dz < B.Dz{
			return B
		}else if A.Px == 1 && A.Dz == B.Dz{
			// 比大小
			if notIntSlice(A.Bz, A.Dz) > notIntSlice(B.Bz, A.Dz) {
				return A
			}
			if notIntSlice(A.Bz, A.Dz) < notIntSlice(B.Bz, A.Dz) {
				return B
			}
			if notIntSlice(A.Bz, A.Dz) == notIntSlice(B.Bz, A.Dz) {
				return nil
			}
		}
		if A.Px != 1 {
			for _, a := range A.Bz {
				for _, b := range B.Bz {
					if a > b {
						return A
					}
					if a <= b {
						break
					}
				}
			}
			return B
		}
	}
	return nil
}
func notIntSlice(s []int, t int) int {
	z := 0
	if t == 21 {
		z = 1
		t = 1
	}
	for _,v := range s {
		if v != t && z == 1{
			return 21
		}
		if v != t && z == 0{
			return v
		}
	}
	return 0
}
