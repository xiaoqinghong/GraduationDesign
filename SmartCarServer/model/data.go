package model

type Data struct {
	Id      int     `json:"id" gorm:"primary_key"`
	CarId   int     `json:"car_id"`
	T       float32 `json:"t"`
	H       float32 `json:"h"`
	Created int64   `json:"created"`
}
