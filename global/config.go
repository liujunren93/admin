package global

const (
	TypeCtrl    = "ctrl"
	TypeDao     = "dao"
	TypeModel   = "model"
	TypeEntity  = "entity"
	TypeDb      = "db"
	TypeRouter  = "router"
	TypeUtils   = "utils"
	TypeConf    = "conf"
	TypeHView   = "hView"
	TypeHApi    = "hApi"
	TypeHRouter = "hrouter"
)

var (
	ApiRoot  = ""
	Mod      = "" //go mod
	WebRoot  = ""
	FilePath = map[string]string{
		TypeCtrl:   "handler/ctrl",
		TypeEntity: "handler/entity",
		TypeModel:  "handler/model",
		TypeDao:    "handler/dao",
		TypeDb:     "db",
		TypeRouter: "router",
		TypeUtils:  "utils",
		TypeConf:   "conf",
		TypeHView:  "/src/views",
		TypeHApi:   "/src/config/routers",
	}
)
