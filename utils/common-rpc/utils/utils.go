package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

const (
	InboundMode  string = "PortHijackMode"
	OriginMode          = ""
	OutboundMode        = "AddrHijackMode"
)

func GenMeshResourceName(usn, lidc, ip string, port int, HijackAddr string) string {
	return "mesh" + "#" + usn + "#" + lidc + "#" + ip + "#" + strconv.Itoa(port) + "#" + HijackAddr
}

var IfHandleThriftErrno bool

func ReadFullFile(filePath string) (string, error) {
	fd, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer fd.Close()
	buf, err := ioutil.ReadAll(fd)
	if err != nil {
		return "", err
	}
	cont := strings.Trim(string(buf), "\n")
	return cont, nil
}

func PrettyJson(v interface{}, indent bool) []byte {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	if indent {
		encoder.SetIndent("", "    ")
	}
	if err := encoder.Encode(v); err != nil {
		return nil
	}
	return bytes.TrimRight(buffer.Bytes(), "\n")
}

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
