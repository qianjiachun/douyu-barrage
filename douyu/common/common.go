package common

import (
	"log"
	"strings"
	"unsafe"
)

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetStrMiddle(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + len(start)
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}