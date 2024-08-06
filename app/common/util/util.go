package util

import (
	"path"
	"path/filepath"
	"runtime"
)

func GetRootDir() string {
	_, callerFile, _, _ := runtime.Caller(0) //Who
	dir := path.Join(path.Dir(callerFile))
	root := filepath.Join(filepath.Dir(dir), "../..")

	return root
}
