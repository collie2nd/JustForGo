package utils

import (
	"CloudRestaurant/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var Db *Orm

type Orm struct {
	*xorm.Engine
}

func GetOrmEngine() (*Orm, error) {
	cfg := GetConfig()
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	log.Println("connect to " + conn)

	engine.ShowSQL(database.ShowSql)

	if err := engine.Sync2(new(model.SmsCode)); err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	Db = orm
	return orm, err

}
