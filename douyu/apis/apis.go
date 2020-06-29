package apis

import (
	"github.com/qianjiachun/douyu/common"
	"net/http"
)

func Init_listen() {
	http.HandleFunc("/douyu/db/barrage/queryBarrage_id", queryBarrage_id)
	http.HandleFunc("/douyu/db/barrage/queryBarrage_uid", queryBarrage_uid)
	http.HandleFunc("/douyu/db/barrage/queryBarrageNum_time", queryBarrageNum_time)
	err := http.ListenAndServe(":27999", nil)
	common.CheckErr(err)
}