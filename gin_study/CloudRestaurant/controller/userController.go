package controller

import (
	"CloudRestaurant/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (uc *UserController) Router(engine *gin.Engine) {
	engine.GET("/user/sendsmscode", uc.SendSmsCode)
}

func (uc *UserController) SendSmsCode(ctx *gin.Context) {
	//发送验证码
	phoneNum, exist := ctx.GetQuery("phone")
	if !exist {
		ctx.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "参数解析失败",
		})
		return
	}

	us := service.UserService{}
	if res := us.SendSmsCode(phoneNum); res {
		ctx.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "短信发送成功",
		})
		return
	} else {
		ctx.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "短信发送失败",
		})
		return
	}

}
