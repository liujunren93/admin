package server

import (
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/global"
	"github.com/liujunren93/admin/utils"
)

func NewCtrl(gs ...*core.Group) {
	var ctrl []File
	for _, g := range gs {
		var f File
		f.Pkg = "ctrl"
		f.Name = g.Name
		f.Import = append(f.Import,
			"github.com/gin-gonic/gin",
			"strconv",
			global.Mod+"/utils",
			global.Mod+"/"+global.FilePath[global.TypeDao],
			global.Mod+"/"+global.FilePath[global.TypeModel],
		)

		var className = utils.UcFirst(g.Name + "Ctrl")
		f.AddVariable(fmt.Sprintf("var %s *%s \n", utils.UcFirst(g.Name), className))
		class := Class{
			Name: className,
		}
		importEntity := false
		for _, dom := range g.List {
			if dom.HasSearch || dom.HPagination {
				importEntity = true
			}
			class.AddFunc(NewFunc(dom.Name+"List", "ctx *gin.context", "", listCtrlStr(dom.Name, dom.Name, dom.HasSearch, dom.HPagination)))
			class.AddFunc(NewFunc(dom.Name+"Create", "ctx *gin.context", "", createCtrlStr(dom.Name, dom.Name)))
			class.AddFunc(NewFunc(dom.Name+"Update", "ctx *gin.context", "", updateCtrlStr(dom.Name, dom.Name)))
			class.AddFunc(NewFunc(dom.Name+"Info", "ctx *gin.context", "", infoCtrlStr(dom.Name)))
			class.AddFunc(NewFunc(dom.Name+"Delete", "ctx *gin.context", "", delCtrlStr(dom.Name)))
		}
		if importEntity {
			f.Import = append(f.Import, global.Mod+"/"+global.FilePath["entity"])
		}
		f.Class = append(f.Class, class)
		f.AddFunc(NewFunc(utils.UpFirst(className), "", "*"+className, ctrlConstructBody(utils.UcFirst(g.Name), className)))
		ctrl = append(ctrl, f)
	}
	Build(global.TypeCtrl,ctrl...)
}

func ctrlConstructBody(variable, className string) string {
	code := `if %s==nil{
	%s=new(%s)
	}
	return %s
`
	return fmt.Sprintf(code, variable, variable, className, variable)
}

func listCtrlStr(entClassName, daoClassName string, search, pagination bool) (bodyStr string) {

	code := ""
	if search || pagination {
		code = `var req entity.%sListReq
	if err := ctx.ShouldBindQuery(&req);err!=nil{
		utils.ResErr400Json(ctx,err,nil)
		return
	}
	list:= dao.%sDao().List(req)
    utils.ResSuccessJson(ctx, "ok", list)`
		bodyStr = fmt.Sprintf(code, entClassName, daoClassName)
		return
	} else {
		code = `list:= dao.%sDao{}.List()
	utils.ResSuccessJson(ctx, "ok", list)`
	}

	bodyStr = fmt.Sprintf(code, daoClassName)
	return

}

func updateCtrlStr(modelClassName, daoClassName string) string {
	code := `param := ctx.Param("id")
	id, _ := strconv.Atoi(param)
	var req model.%s
	if err := ctx.ShouldBindJSON(&req);err!=nil{
	  	utils.ResErrJson(ctx,400,err,nil)
		return
	}
	err:= dao.%sDao().Update(uint(id),req)
	if err!=nil{
		utils.Res500Json(ctx, err, nil)
		return
	}
	utils.ResSuccessJson(ctx, "ok", nil)`
	bodyStr := fmt.Sprintf(code, modelClassName, daoClassName)
	return bodyStr
}

func delCtrlStr(daoClassName string) string {
	code := `param := ctx.Param("id")
	id, _ := strconv.Atoi(param)
	err:= dao.%sDao().Delete(uint(id))
	if err!=nil{
		utils.ResErrJson(ctx,400,err, nil)
		return
	}
	utils.ResSuccessJson(ctx, "ok", nil)`
	bodyStr := fmt.Sprintf(code, daoClassName)
	return bodyStr
}

func infoCtrlStr(daoClassName string) string {
	code := `param := ctx.Param("id")
	id, _ := strconv.Atoi(param)
	info,err:= dao.%sDao().Info(uint(id))
	if err!=nil{
		utils.ResErrJson(ctx,400, err, nil)
		return
	}
	utils.ResSuccessJson(ctx, "ok", map[string]interface{}{"info":info})`
	bodyStr := fmt.Sprintf(code, daoClassName)
	return bodyStr
}

func createCtrlStr(modelClassName, daoClassName string) string {
	code := `var req model.%s
	if err := ctx.ShouldBindJSON(&req);err!=nil{
		utils.ResErrJson(ctx,400,err,nil)
		return
	}
	err:= dao.%sDao().Create(req)
	if err!=nil{
		utils.Res500Json(ctx, err, nil)
		return
	}
	utils.ResSuccessJson(ctx, "ok", nil)`
	bodyStr := fmt.Sprintf(code, modelClassName, daoClassName)
	return bodyStr
}
