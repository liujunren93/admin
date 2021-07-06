package db
import (
	"fmt"
	"github.com/jinzhu/inflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"sync"
	"time"
	"db"

)

var (
	smap                      sync.Map
	Db                        *gorm.DB
	commonInitialisms         = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
	commonInitialismsReplacer *strings.Replacer
)

func init() {

	var commonInitialismsForReplacer []string
	for _, initialism := range commonInitialisms {
		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
	}
	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)
	Db = getConnect()
}

//创建连接
func getConnect() *gorm.DB {
	conf := config.Conf

	//将服务注册到配置文件中
	config.RegistryMonitor(config.Mysql{}, monitor)

	mysqlConf := conf.Mysql
	dsn :=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: NamingStrategy{
			NamingStrategy: schema.NamingStrategy{

			},
		},
	})

	if err != nil {
		panic(err)
	}
	d, err := db.DB()
	if err != nil {
			panic(err)
	}
	d.SetConnMaxIdleTime(10)
	d.SetMaxIdleConns(100)
	d.SetConnMaxIdleTime(time.Hour)
	if conf.System.RunModel == "debug" {
		db = db.Debug()
	}

	return db
}

type NamingStrategy struct {
	schema.NamingStrategy
	TrimStr string
}

// TableName convert string to table name
func (ns NamingStrategy) TableName(table string) string {
	if ns.TrimStr != "" {
		index := strings.LastIndex(table, ns.TrimStr)
		if index >= 0 {
			table = table[:index]
		}

	}

	if ns.SingularTable {
		return ns.TablePrefix + toDBName(table)
	}

	return ns.TablePrefix + inflection.Plural(toDBName(table))

}
func toDBName(name string) string {
	if name == "" {
		return ""
	} else if v, ok := smap.Load(name); ok {
		return fmt.Sprint(v)
	}

	var (
		value                          = commonInitialismsReplacer.Replace(name)
		buf                            strings.Builder
		lastCase, nextCase, nextNumber bool // upper case == true
		curCase                        = value[0] <= 'Z' && value[0] >= 'A'
	)

	for i, v := range value[:len(value)-1] {
		nextCase = value[i+1] <= 'Z' && value[i+1] >= 'A'
		nextNumber = value[i+1] >= '0' && value[i+1] <= '9'

		if curCase {
			if lastCase && (nextCase || nextNumber) {
				buf.WriteRune(v + 32)
			} else {
				if i > 0 && value[i-1] != '_' && value[i+1] != '_' {
					buf.WriteByte('_')
				}
				buf.WriteRune(v + 32)
			}
		} else {
			buf.WriteRune(v)
		}

		lastCase = curCase
		curCase = nextCase
	}

	if curCase {
		if !lastCase && len(value) > 1 {
			buf.WriteByte('_')
		}
		buf.WriteByte(value[len(value)-1] + 32)
	} else {
		buf.WriteByte(value[len(value)-1])
	}

	return buf.String()
}

func monitor(old, new *config.Config) {
	if old.Mysql != new.Mysql {
		fmt.Println("mysql")
		getConnect()
	}
}