package main
import (
  _ "money-record/app/config" // 执行init函数
  "money-record/app/route"
  "money-record/app/middlewares"
  "github.com/gin-gonic/gin"
  "github.com/spf13/viper"
  "fmt"
)

func main() {
  router := gin.Default()
  // 注册中间件
  router.Use(middlewares.Cors())
  // 注册路由组
	route.CollectRoute(router);
	port := viper.GetString("server.port")
	fmt.Println("当前端口", port)
	if port != "" {
		router.Run(":" + port)
	} else {
		router.Run()
	}
}
