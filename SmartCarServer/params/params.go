package params

// 用于记录每一个car的ip
var CarIPMap = make(map[int]string)

// 接口返回结构体
type ResponseData struct {
	Code    int         `json:"code"`
	Error   int         `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 用户登录参数
type UserLogin struct {
	Tel      string `json:"tel" form:"tel"`
	Password string `json:"password" form:"password"`
}

// 用户注册参数
type UserRegister struct {
	Tel      string `json:"tel" form:"tel"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
}

type CarRegister struct {
	Name string `json:"name" form:"name"`
}

// 小车上传数据参数
type UploadData struct {
	CarId int     `json:"car_id" form:"car_id"`
	T     float32 `json:"t" form:"t"`
	H     float32 `json:"h" form:"h"`
}

// 获取小车采集数据的参数
type GetData struct {
	CarId     int   `json:"car_id" form:"car_id"`
	StartTime int64 `json:"start_time" form:"start_time"`
	EndTime   int64 `json:"end_time" form:"end_time"`
}

// 通过id获取信息的参数
type GetInfoById struct {
	Id int `json:"id" form:"id"`
}
