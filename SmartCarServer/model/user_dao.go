package model

type userDao byte

var ud userDao

type IUserDao interface {
	IDao
	Insert(u *User) error
	Delete(uid int) error
	GetUserByTel(tel string) (*User, error)
	GetUserById(id int) (*User, error)
	ModifyInfo(u User) error
}

func GetUserDao() IUserDao {
	return ud
}

func (userDao) GetTable() string {
	return "user"
}

func (u userDao) Insert(user *User) error {
	return DB().Table(u.GetTable()).Create(user).Error

}

func (u userDao) Delete(uid int) error {
	return DB().Table(u.GetTable()).Delete(User{}, "id=?", uid).Error
}

func (u userDao) GetUserByTel(tel string) (*User, error) {
	user := new(User)
	return user, DB().Table(u.GetTable()).First(user, "tel=?", tel).Error
}

func (u userDao) GetUserById(id int) (*User, error) {
	user := new(User)
	err := DB().Table(u.GetTable()).First(user, "id=?", id).Error
	return user, err
}

func (u userDao) ModifyInfo(user User) error {
	return DB().Table(u.GetTable()).Update(user).Error
}
