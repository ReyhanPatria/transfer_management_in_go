package connect

import (
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const dsn string = "root@tcp(127.0.0.1:3306)/db_transfer_management?parseTime=True"

type Database struct {
	Instance *gorm.DB
	Cache    *cache.Cache
}

func NewDB() Database {
	var database Database
	database.Instance = connectPrimary()
	database.Cache = connectCache()
	return database
}

func connectPrimary() *gorm.DB {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbl_",
			SingularTable: true,
		},
	})

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	return db
}

func connectCache() *cache.Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	cache := cache.New(&cache.Options{
		Redis: rdb,
	})

	return cache
}
