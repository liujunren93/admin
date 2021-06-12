package server

import (
	_ "embed"
	"github.com/liujunren93/admin/global"
	"os"
)

const utilsDir = "utils"

//go:embed static/utils.file
var utilsStr string

func NewUtils() {
	err := os.MkdirAll(global.FilePath[global.TypeUtils], 0777)
	if err != nil {
		panic(err)
	}

	create, err := os.Create(global.FilePath[global.TypeUtils] + "/utils.go")
	if err != nil {
		panic(err)
	}
	defer create.Close()
	create.WriteString(utilsStr)

}
