package conf

import (
	"4-ToDoList/cache"
	"4-ToDoList/model"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"strconv"
	"strings"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadMysqlData(file)
	LoadRedis(file)
	path := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
	Redis()
}

// ToDo 使用ini文件对数据库进行连接

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logrus.Info(err)
		panic(err)
	}
	cache.RedisClient = client
	fmt.Println("Redis连接成功")
}
