package model

type SmsCode struct {
	Id         int    `json:"id" xorm:"pk autoincr"`
	Phone      string `json:"phone" xorm:"varchar(11)"`
	BizId      string `json:"biz_id" xorm:"varchar(30)"`
	Code       string `json:"code" xorm:"varchar(6)"`
	CreateTime int64  `json:"create_time" xorm:"bigint"`
}
