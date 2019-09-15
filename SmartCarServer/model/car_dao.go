package model

type carDao byte

var cd carDao

type ICarDao interface {
	IDao
	Insert(car *Car) error
	GetInfo(id int) (*Car, error)
	InsertData(data *Data) error
	GetCarList() (list []*Car, err error)
	GetDataList(carId int, startTime, endTime int64) ([]*Data, error)
}

func GetICarDao() ICarDao {
	return cd
}

func (carDao) GetTable() string {
	return "car"
}

func (carDao) GetDataTable() string {
	return "data"
}

func (c carDao) Insert(car *Car) error {
	return DB().Table(c.GetTable()).Create(car).Error
}

func (c carDao) GetInfo(id int) (*Car, error) {
	car := new(Car)
	return car, DB().First(car, "id=?", id).Error
}

func (c carDao) InsertData(data *Data) error {
	return DB().Table(c.GetDataTable()).Create(data).Error
}

func (c carDao) GetCarList() (list []*Car, err error) {
	err = DB().Table(c.GetTable()).Find(&list).Error
	return list, err
}

func (c carDao) GetDataList(carId int, startTime, endTime int64) (list []*Data, err error) {
	err = DB().Table(c.GetDataTable()).Find(&list, "car_id=? AND created <=? AND created >=?", carId, endTime, startTime).Error
	return list, err
}
