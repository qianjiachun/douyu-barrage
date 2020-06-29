package client

import (
	"encoding/binary"
	"github.com/qianjiachun/douyu/common"
	"github.com/qianjiachun/douyu/msg"
	"io"
	"net"

	"time"
)

const ADDR  = "119.96.201.28:8601"

type DouyuClient struct {
	Rid string
	conn net.Conn
}

func (client *DouyuClient) Connect(callback func(data string)) {
	var err error
	client.conn, err = net.Dial("tcp", ADDR)

	if err != nil {
		println("connect to server failed!")
		return
	}
	_, err = client.conn.Write(msg.BuildDouyuPkg("type@=loginreq/roomid@=" + client.Rid + "/"))
	_, err = client.conn.Write(msg.BuildDouyuPkg("type@=joingroup/rid@=" + client.Rid + "/gid@=-9999/"))
	common.CheckErr(err)
	go client.heartBeat()
	go client.recv(callback)
}

func (client *DouyuClient) recv(callback func(data string)) {
	//danmuReg  := regexp.MustCompile("type@=chatmsg/.*rid@=(\\d*?)/.*uid@=(\\d*).*nn@=(.*?)/txt@=(.*?)/(.*)/")
	for {
		buf := make([]byte, 512)
		if _, err := io.ReadFull(client.conn, buf[:12]); err != nil {
			break
		}
		pl := binary.LittleEndian.Uint32(buf[:4])
		cl := pl - 8
		if cl > 512 {
			buf = make([]byte, cl)
		}

		if _, err := io.ReadFull(client.conn, buf[:cl]); err != nil {
			break
		}
		callback(common.Bytes2str(buf[:cl-1]))
	}

}
func (client *DouyuClient) heartBeat() {
	for {
		_, err := client.conn.Write(msg.BuildDouyuPkg("type@=mrkl/"))
		common.CheckErr(err)
		time.Sleep(time.Second * 40)
	}
}

func (client *DouyuClient) Close() {
	client.Close()
}

