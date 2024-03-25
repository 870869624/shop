package conf

import (
	"github.com/go-ini/ini"
)

var (
	AppModel string
	HttpPort string

	DB     string
	DBHost string
	DBPort string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	RedisDB     string
	RedisAddr   string
	RedisPw     string
	RedisDBName string

	Host        string
	ProductPath string
	AvatarPath  string
)

func Init() {
	//
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMySql(file)
	LoadEmail(file)
	LoadRedis(file)
	LoadPath(file)
	//此处做一个实例，可以做到msql读写分离,以及怎么将DB等配置做成config形式
	//pathRead := strings.Join([]string{root, ":", 123456, "@tcp(", DBHost, ":", DBPort, ")/", taobao, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	//pathWrite := strings.Join([]string{root, ":", 123456, "@tcp(", DBHost, ":", DBPort, ")/", taobao, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
}

func LoadServer(file *ini.File) {
	AppModel = file.Section("service").Key("AppModel").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}
func LoadMySql(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DBHost = file.Section("mysql").Key("DBHost").String()
	DBPort = file.Section("mysql").Key("DBPort").String()
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}

func LoadRedis(file *ini.File) {
	RedisDB = file.Section("redis").Key("RedisDB").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDBName = file.Section("redis").Key("RedisDBName").String()
}

func LoadPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}
