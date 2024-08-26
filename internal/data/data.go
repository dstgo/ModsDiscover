package data

import (
	"context"
	"fmt"
	"github.com/dstgo/tracker/conf"
	"github.com/dstgo/tracker/internal/assets"
	"github.com/oschwald/geoip2-golang"
	"github.com/qiniu/qmgo"
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

func LoadMongoDB(ctx context.Context, dbConf conf.DBConf) (*qmgo.QmgoClient, error) {
	//  [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	uri := fmt.Sprintf("mongodb://%s:%s@%s", dbConf.User, dbConf.Password, dbConf.Address)
	mgocli, err := qmgo.Open(ctx, &qmgo.Config{Uri: uri, Database: dbConf.DataBase})
	if err != nil {
		return nil, err
	}
	// ping
	if err := mgocli.Ping(5); err != nil {
		return nil, err
	}

	return mgocli, nil
}

func LoadGeoIpDBInMem(file string) (*geoip2.Reader, error) {
	bytes, err := assets.FS.ReadFile(assets.GeopIp2CityDB)
	if err != nil {
		return nil, err
	}
	reader, err := geoip2.FromBytes(bytes)
	if err != nil {
		return nil, err
	}
	return reader, nil
}
