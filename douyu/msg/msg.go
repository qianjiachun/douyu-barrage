package msg

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/qianjiachun/douyu/common"
	"github.com/qianjiachun/douyu/db"
)



func BuildDouyuPkg(str string) []byte {
	var err error
	data := new(bytes.Buffer)
	rawLen := len([]byte(str)) + 9
	err = binary.Write(data, binary.LittleEndian, int32(rawLen))
	err = binary.Write(data, binary.LittleEndian, int32(rawLen))
	err = binary.Write(data, binary.LittleEndian, int16(689))
	err = binary.Write(data, binary.LittleEndian, byte(0))
	err = binary.Write(data, binary.LittleEndian, byte(0))
	err = binary.Write(data, binary.LittleEndian, []byte(str))
	err = binary.Write(data, binary.LittleEndian, byte(0))
	common.CheckErr(err)
	return data.Bytes()
}

func BuildRetJson(items []db.Struct_Douyu_Barrage) string {
	arr := make([]map[string]interface{},0)
	ret := make(map[string]interface{})
	for i := 0; i < len(items); i++ {
		tmp := make(map[string]interface{})
		tmp["rid"] = items[i].RID
		tmp["uid"] = items[i].UID
		tmp["id"] = items[i].ID
		tmp["level"] = items[i].Level
		tmp["barrage"] = items[i].Barrage
		tmp["timestamp"] = items[i].Timestamp
		arr = append(arr, tmp)
	}
	ret["data"] = arr
	retJson, _ := json.Marshal(ret)
	return string(retJson)

}