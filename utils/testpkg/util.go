package testpkg

import (
	"io/ioutil"
	"path"
	"runtime"
)

// GetMockDataContent 获取testpkg下相对路径的文件内容
func GetMockDataContent(fp string) string {
	var dir string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		dir = path.Dir(filename)
	}

	data, _ := ioutil.ReadFile(dir + fp)
	return string(data)
}
