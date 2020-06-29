package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io"
	"net"
	"regexp"
	"strings"
	"time"
	"unsafe"
)

const ADDR  = "119.96.201.28:8601"

func main() {
	//roomid := flag.String("room","99999","roomid")
	//flag.Parse()
	//println("start connecting server .....")
	//conn,err := net.Dial("tcp",ADDR)
	//if err != nil {
	//	println("connect to server failed!")
	//	return
	//}
	//println("connect to server success!")
	//println("start into room.....")
	//conn.Write(buildRequest("type@=loginreq/roomid@=" + *roomid + "/"))
	//conn.Write(buildRequest("type@=joingroup/rid@=" + *roomid + "/gid@=-9999/"))
	//
	//go reciveMessage(conn)
	//go heartBeat(conn)
	go startRoom("5189167")
	//go startRoom("1322845")

	select {

	}
	//var wg sync.WaitGroup
	//wg.Add(1)
	//wg.Wait()
}
func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func startRoom(rid string) {
	conn,err := net.Dial("tcp",ADDR)
	if err != nil {
		println("connect to server failed!")
		return
	}
	conn.Write(buildRequest("type@=loginreq/roomid@=" + rid + "/"))
	conn.Write(buildRequest("type@=joingroup/rid@=" + rid + "/gid@=-9999/"))
	go heartBeat(conn)
	danmuReg  := regexp.MustCompile("type@=chatmsg/.*rid@=(\\d*?)/.*uid@=(\\d*).*nn@=(.*?)/txt@=(.*?)/(.*)/")
	for true {
		buf := make([]byte, 512)
		if _, err := io.ReadFull(conn, buf[:12]); err != nil {
			break
		}
		pl := binary.LittleEndian.Uint32(buf[:4])
		//code := binary.LittleEndian.Uint16(buf[8:10])
		cl := pl - 8
		if cl > 512 {
			buf = make([]byte, cl)
		}
		if _, err := io.ReadFull(conn, buf[:cl]); err != nil {
			break
		}
		match := danmuReg.FindStringSubmatch(bytes2str(buf[:cl-1]))
		if len(match) >0 {
			println(getFormatTime() +"\t" +match[3]+"\t:" +match[4])
		}



		//buf := make([]byte, 1024 * 20)
		//bufLen, _ := conn.Read(buf)
		//offset := 0
		//
		//for offset <= bufLen + 12 {
		//	eachLen := binary.LittleEndian.Uint32(buf[offset:offset + 4])
		//	println(bytes2str(buf[offset+12:uint32(offset)+eachLen+4]),"\n")
		//	offset = offset + int(eachLen) + 4
		//	//println(offset)
		//
		//}



		//println(n,"||||",bytes2str(buf[:n]),"\n")
		//data := hex.EncodeToString(buf[:n])
		//messages := strings.Split(data,"b2020000")
		//for _, m := range messages {
		//	if strings.Contains(m,"00"){
		//		end := strings.Index(m,"00")
		//		m = string([]rune(m)[:end])
		//		mb, _ := hex.DecodeString(m)
		//		dm := string(mb)
		//		match := danmuReg.FindStringSubmatch(dm)
		//		if len(match) >0 {
		//			println(getFormatTime() +"\t" +match[3]+"\t:" +match[4])
		//		}
		//	}
		//}
	}
}

func reciveMessage(conn net.Conn) {
	danmuReg  := regexp.MustCompile("type@=chatmsg/.*rid@=(\\d*?)/.*uid@=(\\d*).*nn@=(.*?)/txt@=(.*?)/(.*)/")
	for true {
		buf := make([]byte, 1024 * 80)
		n,_ := conn.Read(buf)
		data := hex.EncodeToString(buf[:n])
		messages := strings.Split(data,"b2020000")
		for _, m := range messages {
			if strings.Contains(m,"00"){
				end := strings.Index(m,"00")
				m = string([]rune(m)[:end])
				mb, _ := hex.DecodeString(m)
				dm := string(mb)
				match := danmuReg.FindStringSubmatch(dm)
				if len(match) >0 {
					println(getFormatTime() +"\t" +match[3]+"\t:" +match[4])
				}
			}
		}
	}
}


func getFormatTime() string{
	return time.Now().Format("2006-01-02 15:04:05")
}
func heartBeat(conn net.Conn) {
	for {
		conn.Write(buildRequest("type@=mrkl/"))
		println("beats")
		time.Sleep(time.Second * 40)
	}
}


func buildRequest(str string) []byte{
	data := new(bytes.Buffer)
	rawLen := len([]byte(str)) + 9
	binary.Write(data, binary.LittleEndian, int32(rawLen))
	binary.Write(data, binary.LittleEndian, int32(rawLen))
	binary.Write(data, binary.LittleEndian, int16(689))
	binary.Write(data, binary.LittleEndian, byte(0))
	binary.Write(data, binary.LittleEndian, byte(0))
	binary.Write(data, binary.LittleEndian, []byte(str))
	binary.Write(data, binary.LittleEndian, byte(0))
	return data.Bytes()
}