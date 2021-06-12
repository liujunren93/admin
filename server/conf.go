package server

import (
	_ "embed"
	"github.com/liujunren93/admin/global"
	"os"
)

const confDir = "config"

//go:embed static/conf.file
var confStr string

//go:embed static/conf.template
var ymlStr string

func NewConfigFile() {
	err := os.MkdirAll(global.FilePath[global.TypeConf], 0777)
	if err != nil {
		panic(err)
	}
	// 创建结构体
	create, err := os.Create(global.FilePath[global.TypeConf] + "/conf.go")
	if err != nil {
		panic(err)
	}
	create.WriteString(confStr)
	// 创建yaml文件
	yml, err := os.Create(global.FilePath[global.TypeConf] + "/conf.yml")
	if err != nil {
		panic(err)
	}
	yml.WriteString(ymlStr)
	create.Close()
}
