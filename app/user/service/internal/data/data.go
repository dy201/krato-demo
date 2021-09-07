package data

import (
	"dy201.com/doyle-blog/app/user/service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type Data struct {
	db  *gorm.DB
}

func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper( log.With(logger, "module", "data/user"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	d := &Data{
		db:  db,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
	}, nil
}
