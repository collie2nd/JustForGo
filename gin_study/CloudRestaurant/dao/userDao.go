package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/utils"
	"log"
)

type UserDao struct {
	*utils.Orm
}

func (ud *UserDao) InsertRecord(sc model.SmsCode) int64 {
	res, err := ud.InsertOne(&sc)
	if err != nil {
		log.Panicln(err.Error())
	}
	return res
}
