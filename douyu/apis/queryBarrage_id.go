package apis

import (
	"fmt"
	"github.com/qianjiachun/douyu/common"
	"github.com/qianjiachun/douyu/db"
	"github.com/qianjiachun/douyu/msg"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

func queryBarrage_id(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*") // 跨域 "*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	writer.Header().Set("content-type", "application/json")  //返回数据格式是json

	rid := request.PostFormValue("rid")
	id := request.PostFormValue("id")
	page, _ := strconv.Atoi(request.PostFormValue("page"))
	minTime, _ := strconv.Atoi(request.PostFormValue("minTime"))
	maxTime, _ := strconv.Atoi(request.PostFormValue("maxTime"))

	offset := page * 20
	var items []db.Struct_Douyu_Barrage
	err := db.Collection.Find(bson.M{
		"$and": []bson.M{
			bson.M{"rid": bson.M{"$regex": rid}},
			bson.M{"id": bson.M{"$regex": id}},
			bson.M{"timestamp": bson.M{
				"$gte": minTime,
				"$lte": maxTime,
				},
			},
		},
	}).Skip(offset).Limit(20).Sort("-timestamp").All(&items)
	common.CheckErr(err)

	_, _ = fmt.Fprint(writer, msg.BuildRetJson(items))
}

