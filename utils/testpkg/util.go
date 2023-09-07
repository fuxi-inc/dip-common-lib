package testpkg

import (
	"encoding/base64"
	"io/ioutil"
	"path"
	"runtime"
)

// GetMockDataContent 获取testpkg下相对路径的文件内容
// 如： testpkg.GetMockDataContent("/mock_data/user/bob/private.hex") 读取的是 /utils/testpkg/mock_data/user/bob/private.hex的文件内容
func GetMockDataContent(fp string) []byte {
	var dir string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		dir = path.Dir(filename)
	}

	data, _ := ioutil.ReadFile(dir + fp)
	return data
}

func Base64DecodeFile(sourceFile, destFile string) error {
	fileContent, _ := ioutil.ReadFile(sourceFile)
	destContent := make([]byte, base64.StdEncoding.DecodedLen(len(string(fileContent))))
	base64.StdEncoding.Decode(destContent, fileContent)
	return ioutil.WriteFile(destFile, destContent, 0644)
}
