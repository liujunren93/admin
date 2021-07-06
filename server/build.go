package server

import (
	"github.com/liujunren93/admin/global"
	"os"
	"path/filepath"
)

func Build(fileType string,files ...File)  {
	abs, err := filepath.Abs(global.ApiRoot+"/"+global.FilePath[fileType])
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(abs, 0666)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		abs, err := filepath.Abs(abs+"/"+file.Name + ".go")
		create, err := os.Create(abs)

		if err != nil {
			panic(err)
		}
		create.WriteString(file.String())
		create.Close()
	}
}
