package util

import (
	"strings"
)

//func GetRootDir() string {
//	_, callerFile, _, _ := runtime.Caller(0) //Who
//	dir := path.Join(path.Dir(callerFile))
//	root := filepath.Join(filepath.Dir(dir), "../..")
//
//	return root
//}

func ExtractImgTypeFromBase64(data string) string {
	index := strings.Index(data, "base64")
	if index == -1 { //no such index
		return ""
	}
	info := data[:index-1]
	index = strings.Index(info, "/")
	infoType := strings.Split(info[index+1:], ";")
	return infoType[0]
}
