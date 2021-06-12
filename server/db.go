package server

import (
	_ "embed"
	"fmt"
	"github.com/liujunren93/admin/global"
	"os"
)
//go:embed static/db.file
var dbStr string



func NewDBFile() {
	err := os.MkdirAll(global.FilePath[global.TypeDb], 0777)
	if err != nil {
		panic(err)
	}
	create, err := os.Create(global.FilePath[global.TypeDb] + "/mysql.go")
	if err != nil {
		panic(err)
	}
	defer create.Close()
	dsn :=`fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.DB)`
	create.WriteString(fmt.Sprintf(dbStr, global.FilePath[global.TypeDb],dsn))

}
