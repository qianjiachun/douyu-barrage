package db

import (
	"github.com/qianjiachun/douyu/common"
	"gopkg.in/mgo.v2"
	"strconv"
	"time"
)

var Session *mgo.Session
var Collection *mgo.Collection

//func ConnetDB(url string) (*mgo.Session, *mgo.Collection) {
//	session, err := mgo.Dial(url)
//	common.CheckErr(err)
//	db := session.DB("douyu").C("data")
//	return session, db
//}

func ConnetDB(url string) {
	var err error
	Session, err = mgo.Dial(url)
	common.CheckErr(err)
	Collection = Session.DB("douyu").C("data")
	//defer Session.Close()
}

func InsertDB(str string) {
	msgType := common.GetStrMiddle(str, "type@=", "/")
	if msgType == "chatmsg" {
		rid := common.GetStrMiddle(str, "rid@=", "/")
		uid := common.GetStrMiddle(str, "uid@=", "/")
		id := common.GetStrMiddle(str, "nn@=", "/")
		level, _ := strconv.Atoi(common.GetStrMiddle(str, "level@=", "/"))
		barrage := common.GetStrMiddle(str, "txt@=", "/")
		nowTime := time.Now()
		timestamp := int(nowTime.Unix())
		time := nowTime.UTC()
		data := Struct_Douyu_Barrage{
			RID: rid,
			UID: uid,
			ID: id,
			Level: level,
			Barrage: barrage,
			Timestamp: timestamp,
			Time: time,
		}

		err := Collection.Insert(&data)
		common.CheckErr(err)
	}

}