package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoqinghong/SmartCarServer/conf"
	"github.com/xiaoqinghong/SmartCarServer/handler"
)

func init() {
	router := gin.Default()
	if conf.GetConfig().App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	api := router.Group("/api")
	{
		// user
		api.POST("/user/login", handler.Login)
		api.POST("/user/register", handler.Register)
		api.GET("/user/get_info", handler.GetUserInfo)
		// car
		api.POST("/car/register", handler.RegisterCar)
		api.POST("/car/upload_data", handler.UploadData)
		api.GET("/car/get_list", handler.GetCarList)
		api.GET("/car/get_data", handler.GetData)
	}
	err := router.Run(conf.GetConfig().App.ApiPort)
	if err != nil {
		panic(err)
	}
}
