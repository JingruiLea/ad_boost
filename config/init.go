package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisUsername string
	RedisDBIndex  int

	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
}

var workDir string
var Configs *Config

func init() {
	workDir = os.Getenv("WORK_DIR")
	if workDir == "" {
		workDir, _ = os.Getwd()
	}
	env := os.Getenv("ENV")
	if env != "" {
		_ = godotenv.Overload(".env." + env)
	} else {
		_ = godotenv.Overload("/Users/ljr/go/src/github.com/JingruiLea/ad_boost/.env")
	}

	Configs = &Config{}
	Configs.RedisHost = os.Getenv("REDIS_HOST")
	if Configs.RedisHost == "" {
		panic("redis host is empty")
	}
	Configs.RedisPort = os.Getenv("REDIS_PORT")
	if Configs.RedisPort == "" {
		panic("redis_dal port is empty")
	}
	Configs.RedisPassword = os.Getenv("REDIS_PASSWORD")
	if Configs.RedisPassword == "" {
		panic("redis_dal password is empty")
	}
	Configs.RedisUsername = os.Getenv("REDIS_USERNAME")
	if Configs.RedisUsername == "" {
		panic("redis_dal username is empty")
	}
	Configs.RedisDBIndex = GetRedisDBIndex()

	Configs.DBHost = os.Getenv("DB_HOST")
	if Configs.DBHost == "" {
		panic("db host is empty")
	}
	Configs.DBPort = os.Getenv("DB_PORT")
	if Configs.DBPort == "" {
		panic("db port is empty")
	}
	Configs.DBName = os.Getenv("DB_NAME")
	if Configs.DBName == "" {
		panic("db name is empty")
	}
	Configs.DBUser = os.Getenv("DB_USER")
	if Configs.DBUser == "" {
		panic("db auth is empty")
	}
	Configs.DBPass = os.Getenv("DB_PASSWORD")
	if Configs.DBPass == "" {
		panic("db pass is empty")
	}

}

func GetLogPath() string {
	return fmt.Sprintf("%s/log/%s.app.log", workDir, MustGetSvcName())
}

func MustGetSvcName() string {
	ret := os.Getenv("SERVICE_NAME")
	if ret == "" {
		return "taimer.backend.go"
	}
	return ret
}

func GetRedisDBIndex() int {
	return int(Str2I64(os.Getenv("REDIS_DB_INDEX")))
}

func Str2I64(s string, defaultValue ...int64) int64 {
	var ret int64
	ret, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return ret
}
