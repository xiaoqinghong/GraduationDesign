package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaoqinghong/SmartCarServer/model"
	"github.com/xiaoqinghong/SmartCarServer/params"
	"net/http"
)

func Login(c *gin.Context) {
	data := new(params.UserLogin)
	data.Tel = c.PostForm("tel")
	data.Password = c.PostForm("password")
	//c.ShouldBindJSON(data)
	c.JSON(http.StatusOK, model.GetIUser().Login(data.Tel, data.Password))
}

func Register(c *gin.Context) {
	data := new(params.UserRegister)
	data.Tel = c.PostForm("tel")
	data.Password = c.PostForm("password")
	data.Name = c.PostForm("name")
	//c.ShouldBindJSON(data)
	c.JSON(http.StatusOK, model.GetIUser().Register(data.Tel, data.Password, data.Name))
}

func GetUserInfo(c *gin.Context) {
	data := new(params.GetInfoById)
	c.ShouldBindJSON(data)
	c.JSON(http.StatusOK, model.GetIUser().GetInfo(data.Id))
}

func RegisterCar(c *gin.Context) {
	data := new(params.CarRegister)
	c.ShouldBindJSON(data)
	fmt.Println(data.Name)
	c.JSON(http.StatusOK, model.GetICar().Register(data.Name))
}

func UploadData(c *gin.Context) {
	data := new(params.UploadData)
	c.ShouldBindJSON(data)
	params.CarIPMap[data.CarId] = c.Request.RemoteAddr
	model.GetICar().UploadData(data.CarId, data.T, data.H)
	c.Status(http.StatusOK)
}

func GetCarList(c *gin.Context) {
	c.JSON(http.StatusOK, model.GetICar().GetCarList())
}

func GetData(c *gin.Context) {
	data := new(params.GetData)
	c.BindQuery(data)
	c.JSON(http.StatusOK, model.GetICar().GetData(data.CarId, data.StartTime, data.EndTime))
}
