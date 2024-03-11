package data

import (
	"context"
	"fmt"
	"github.com/dstgo/tracker/conf"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadGormDB(dbConf conf.DBConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConf.User, dbConf.Password, dbConf.Address, dbConf.Params)
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func LoadRedisDB(rdConf conf.RedisConf) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     rdConf.Address,
		Password: rdConf.Auth,
	})
	pingRes := client.Ping(context.Background())
	if err := pingRes.Err(); err != nil {
		return nil, err
	}
	return client, nil
}
