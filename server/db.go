package server

import (
	_ "embed"
	"fmt"
	"github.com/liujunren93/admin/global"
	"os"
	"path/filepath"
)
//go:embed static/db.file
var dbStr string



func NewDBFile() {
	path, err := filepath.Abs(global.ApiRoot + "/" + global.FilePath[global.TypeDb])
	err = os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}
	create, err := os.Create(path+ "/mysql.go")
	if err != nil {
		panic(err)
	}
	defer create.Close()
	dsn :=`fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.DB)`
	create.WriteString(fmt.Sprintf(dbStr,global.Mod+"/conf",dsn))

}
