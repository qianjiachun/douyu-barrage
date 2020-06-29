package apis

import (
	"encoding/json"
	"fmt"
	"github.com/qianjiachun/douyu/common"
	"github.com/qianjiachun/douyu/db"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)


func queryBarrageNum_time(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*") // 跨域 "*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	writer.Header().Set("content-type", "application/json")  //返回数据格式是json

	rid := request.PostFormValue("rid")

	minTime, _ := strconv.Atoi(request.PostFormValue("minTime"))
	maxTime, _ := strconv.Atoi(request.PostFormValue("maxTime"))

	type list []interface{}
	date := time.Unix(0, 0)
	p1 := bson.M{
		"$match": bson.M{
			"timestamp": bson.M{
				"$gte": minTime,
				"$lt": maxTime,
			},
			"rid": rid,
		},
	}

	sub := list{"$time", date}
	sub_base := bson.M{"$subtract": sub}

	mode_list := list{sub_base, 1000*60*30}
	mode := bson.M{"$mod": mode_list}

	sub_array := []bson.M{}
	sub_array = append(sub_array, sub_base)
	sub_array = append(sub_array, mode)
	var p2 = bson.M{"$group": bson.M{"_id": bson.M{"$subtract": sub_array}, "barrageNum": bson.M{"$sum": 1}}}
	add_lists := list{date, "$_id"}
	p3 := bson.M{
		"$project": bson.M{
			"barrageNum":1,
			"date": bson.M{"$add": add_lists},
		},
	}
	p4 := bson.M{
		"$sort": bson.M{
			"date": 1,
		},
	}
	options := []bson.M{p1, p2, p3, p4}

	pipe := db.Collection.Pipe(options)

	var results = []bson.M{}
	err := pipe.All(&results)

	x := make([]string,0)
	y := make([]int,0)
	ret := make(map[string]interface{})
	for i := 0; i < len(results); i++ {
		x = append(x, results[i]["date"].(time.Time).Format("01-02 15:04:05"))
		y = append(y, results[i]["barrageNum"].(int))
	}
	ret["x"] = x
	ret["y"] = y
	retJson, _ := json.Marshal(ret)
	common.CheckErr(err)

	_, _ = fmt.Fprint(writer, string(retJson))
}

