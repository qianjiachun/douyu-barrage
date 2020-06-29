package main

import (
	"encoding/hex"
	"github.com/qianjiachun/douyu/client"
	"github.com/qianjiachun/douyu/db"
	"github.com/qianjiachun/douyu/msg"
)

var arr []client.DouyuClient

func main() {
	//println("*************斗鱼分布式弹幕采集 -By 小淳*************")
	//println("读取直播间号中...")
	//rooms, err := ioutil.ReadFile("rooms.cfg")
	//common.CheckErr(err)
	//roomlist := strings.Split(string(rooms), "\n")
	//println("读取直播间完毕，连接数据库中...")
	//db.ConnetDB("mongodb://xiaochunchun:xiaochunchun77@122.51.5.63:29021")
	//defer _defer()
	//println("连接数据库成功")
	//for i := 0; i < len(roomlist); i++ {
	//	println("开始采集直播间：" + roomlist[i])
	//	arr = append(arr, client.DouyuClient{Rid: roomlist[i]})
	//}
	//for i := 0; i < len(arr); i++ {
	//	arr[i].Connect(func(data string) {
	//		db.InsertDB(data)
	//	})
	//}
	//
	//apis.Init_listen()
	println("原文：type@=loginreq/roomid@=5189167")
	println("[]byte：", msg.BuildDouyuPkg("type@=loginreq/roomid@=5189167"))
	println("16进制：", hex.EncodeToString(msg.BuildDouyuPkg("type@=loginreq/roomid@=5189167")))

}

func _defer() {
	db.Session.Close()
	for i := 0; i < len(arr); i++ {
		arr[i].Close()
	}
}