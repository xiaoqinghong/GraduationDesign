package model

import (
	"github.com/jinzhu/gorm"
	"github.com/xiaoqinghong/SmartCarServer/params"
	"github.com/xiaoqinghong/SmartCarServer/utils"
)

type User struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Tel      string `json:"tel"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

var user = User{}

type IUser interface {
	Login(tel, password string) params.ResponseData
	Register(tel, password, name string) params.ResponseData
	GetInfo(id int) params.ResponseData
}

func GetIUser() IUser {
	return user
}

func (User) Login(tel, password string) params.ResponseData {
	resp := params.ResponseData{}
	if len(tel) == 0 || len(password) == 0 {
		resp.Error = 1
		resp.Message = "账号或密码不能为空"
	} else if !utils.IsPhoneNum(tel) {
		resp.Error = 1
		resp.Message = "手机号码格式不正确"
	} else {
		u, err := GetUserDao().GetUserByTel(tel)
		if err != nil {
			resp.Error = 1
			resp.Message = err.Error()
		} else {
			if u.Password != password {
				resp.Error = 1
				resp.Message = "密码错误"
			} else {
				resp.Message = "登录成功"
			}
		}
	}
	return resp
}

func (User) Register(tel, password, name string) params.ResponseData {
	resp := params.ResponseData{}
	if len(tel) == 0 || len(password) == 0 {
		resp.Error = 1
		resp.Message = "账号或密码不能为空"
	} else if !utils.IsPhoneNum(tel) {
		resp.Error = 1
		resp.Message = "手机号码格式不正确"
	} else if len(password) < 6 {
		resp.Error = 1
		resp.Message = "密码至少是6位"
	} else {
		_, err := GetUserDao().GetUserByTel(tel)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				userT := new(User)
				userT.Tel = tel
				userT.Password = password
				userT.Name = name
				err := GetUserDao().Insert(userT)
				if err != nil {
					resp.Error = 1
					resp.Message = err.Error()
				} else {
					resp.Message = "注册成功"
				}
			}
		} else {
			resp.Error = 1
			resp.Message = "该用户已注册"
		}
	}
	return resp
}

func (User) GetInfo(id int) params.ResponseData {
	resp := params.ResponseData{}
	u, err := GetUserDao().GetUserById(id)
	if err != nil {
		resp.Error = 1
		if err == gorm.ErrRecordNotFound {
			resp.Message = "用户不存在"
		} else {
			resp.Message = "system error"
		}
	} else {
		resp.Data = u
	}
	return resp
}
