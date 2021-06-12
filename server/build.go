package server

import (
	"github.com/liujunren93/admin/global"
	"os"
)

func Build(fileType string,files ...File)  {
	os.MkdirAll(global.FilePath[fileType],0666)
	for _, file := range files {
		create, err := os.Create(global.FilePath[fileType]+"/"+file.Name + ".go")

		if err != nil {
			panic(err)
		}
		create.WriteString(file.String())
		create.Close()
	}
}
