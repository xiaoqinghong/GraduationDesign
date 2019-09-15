package model

import (
	"fmt"
	"time"

	"github.com/xiaoqinghong/SmartCarServer/params"
)

type Car struct {
	Id      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"name"`
	Created int64  `json:"created" gorm:"created"`
	IP      string `json:"ip" gorm:"_"`
}

var car = Car{}

type ICar interface {
	Register(name string) params.ResponseData
	GetInfo(id int) params.ResponseData
	UploadData(carId int, t, h float32)
	GetCarList() params.ResponseData
	GetData(carId int, startTime, endTime int64) params.ResponseData
}

func GetICar() ICar {
	return car
}

func (Car) Register(name string) params.ResponseData {
	resp := params.ResponseData{}
	if len(name) == 0 {
		resp.Error = 1
		resp.Message = "注册名称不能为空"
	} else {
		carT := new(Car)
		carT.Name = name
		carT.Created = time.Now().Unix()
		err := GetICarDao().Insert(carT)
		if err != nil {
			resp.Error = 1
			resp.Message = "system error"
		} else {
			resp.Message = "注册成功"
		}
	}
	return resp
}

func (Car) GetInfo(id int) params.ResponseData {
	return params.ResponseData{}
}

func (Car) UploadData(carId int, t, h float32) {
	data := new(Data)
	data.CarId = carId
	data.T = t
	data.H = h
	data.Created = time.Now().Unix()
	err := GetICarDao().InsertData(data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (Car) GetCarList() params.ResponseData {
	resp := params.ResponseData{}
	resList, err := GetICarDao().GetCarList()
	for i := 0; i < len(resList); i++ {
		resList[i].IP = params.CarIPMap[resList[i].Id]
	}
	if err != nil {
		resp.Error = 1
		resp.Message = err.Error()
	} else {
		resp.Data = resList
	}
	return resp
}

func (Car) GetData(carId int, startTime, endTime int64) params.ResponseData {
	resp := params.ResponseData{}
	resList, err := GetICarDao().GetDataList(carId, startTime, endTime)
	if err != nil {
		resp.Error = 1
		resp.Message = err.Error()
	} else {
		resp.Data = resList
	}
	return resp
}
