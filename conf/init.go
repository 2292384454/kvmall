package conf

import (
	"github.com/joho/godotenv"
	"kvmall/cache"
	model "kvmall/models"
	"os"
)

func Init() {
	// 从本地读取环境变量
	godotenv.Load()
	// 连接 mysql 数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	// 连接 redis 数据库
	cache.Redis()
}
