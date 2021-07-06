package server

import (
	_ "embed"
	"github.com/liujunren93/admin/global"
	"os"
	"path/filepath"
)

const confDir = "config"

//go:embed static/conf.file
var confStr string

//go:embed static/conf.template
var ymlStr string

func NewConfigFile() {
	abs, err := filepath.Abs(global.ApiRoot + "/" + global.FilePath[global.TypeConf])
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(abs, 0777)
	if err != nil {
		panic(err)
	}
	// 创建结构体
	create, err := os.Create(abs+ "/conf.go")
	if err != nil {
		panic(err)
	}
	create.WriteString(confStr)
	// 创建yaml文件
	yml, err := os.Create(abs + "/conf.yml")
	if err != nil {
		panic(err)
	}
	yml.WriteString(ymlStr)
	create.Close()
}
