package server

import (
	_ "embed"
	"github.com/liujunren93/admin/global"
	"os"
	"path/filepath"
)

const utilsDir = "utils"

//go:embed static/utils.file
var utilsStr string

func NewUtils() {
	path, err := filepath.Abs(global.ApiRoot + "/" + global.FilePath[global.TypeUtils])
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}

	create, err := os.Create(path + "/utils.go")
	if err != nil {
		panic(err)
	}
	defer create.Close()
	create.WriteString(utilsStr)

}
