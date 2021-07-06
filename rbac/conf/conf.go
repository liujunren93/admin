package config
import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"sync"
)
type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
}

type System struct {
	HttpPort  string `yaml:"httpPort"`
	Name      string `yaml:"name"`
	RunModel  string `yaml:"runModel"`
	LogPath   string `yaml:"logPath"`
	JwtSecret string `yaml:"jwtSecret"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"database"`
}

type Redis struct {
	Use      bool   `yaml:"use"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var (
	Conf    *Config
	v       = viper.New()
	monitor sync.Map
)

type MonitorFunc func(old, new *Config)

func RegistryMonitor(name interface{}, f MonitorFunc) {
	monitor.Store(name, f)
}

//获取配置文件
func init() {
	if Conf == nil {
		v.SetConfigName("conf")
		v.SetConfigType("yaml")
		v.AddConfigPath("config")
		v.WatchConfig()
		v.OnConfigChange(func(in fsnotify.Event) { //更新配置文件
			var tmpConf Config
			err := v.Unmarshal(&tmpConf)
			if err != nil {
				fmt.Println("config err ", err)
			}
			monitor.Range(func(key, value interface{}) bool {
				monitorFunc := value.(MonitorFunc)
				monitorFunc(Conf, &tmpConf)
				return true
			})

			Conf = &tmpConf
		})
		err := v.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
	}

	err := v.Unmarshal(&Conf)
	if err != nil {
		fmt.Println("config err ", err)
	}

}