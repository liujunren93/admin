package global

const (
	TypeCtrl   = "ctrl"
	TypeDao    = "dao"
	TypeModel  = "model"
	TypeEntity = "entity"
	TypeDb     = "db"
	TypeRouter = "router"
	TypeUtils  = "utils"
	TypeConf   = "conf"
)

var (
	Mod      = "" //go mod
	FilePath = map[string]string{
		TypeCtrl:   "handler/ctrl",
		TypeEntity: "handler/entity",
		TypeModel:  "handler/model",
		TypeDao:    "handler/dao",
		TypeDb:     "db",
		TypeRouter: "router",
		TypeUtils:  "handler/utils",
		TypeConf:   "conf",
	}
)
