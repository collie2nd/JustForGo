package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := utils.ParseConfig("./config/app.json")

	if err != nil {
		log.Panic(err)
	}

	_, err = utils.GetOrmEngine()
	if err != nil {
		log.Fatalln(err)
	}

	app := gin.Default()
	registerRouter(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine) {
	//curl http://localhost:8090/hello
	new(controller.HelloController).Router(router)

	//需要注册阿里云短信验证码服务
	//curl http://localhost:8090/user/sendsmscode?phone=19572468158
	new(controller.UserController).Router(router)

}
