package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"CloudRestaurant/utils"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {
}

func (us *UserService) SendSmsCode(phoneNum string) bool {
	//1.产生一个验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	fmt.Println(code)
	//code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))

	//2.调用阿里云SDK，完成发送
	//cfg := utils.GetConfig().Sms
	//client, err := dysmsapi.NewClientWithAccessKey(cfg.RegionId, cfg.AppKey, cfg.AppSecret)
	//if err != nil {
	//	//log.Fatalln(err.Error())
	//	fmt.Println(phoneNum)
	//	return false
	//}
	//
	//request := dysmsapi.CreateSendSmsRequest()
	//request.Scheme = "https"
	//request.SignName = cfg.SignName
	//request.TemplateCode = cfg.TemplateCode
	//request.PhoneNumbers = phoneNum
	//codeParam, err := json.Marshal(map[string]interface{}{
	//	"code": code,
	//})
	//request.TemplateParam = string(codeParam)
	//
	//response, err := client.SendSms(request)
	//if err != nil {
	//	//log.Fatalln(err.Error())
	//	fmt.Println(phoneNum)
	//	return false
	//}

	//3.接受返回结果，并判断发送状态
	//if response.Code == "OK" {
	//	return true
	//}
	//return false

	//将结果储存到数据库中
	smsCode := model.SmsCode{
		Phone:      phoneNum,
		BizId:      "0",
		Code:       code,
		CreateTime: time.Now().Unix(),
	}
	userDao := dao.UserDao{Orm: utils.Db}
	res := userDao.InsertRecord(smsCode)
	fmt.Println(res)
	return res > 0
}
