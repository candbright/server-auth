package dao

import (
	"github.com/candbright/go-core/errors"
	"piano-server/server/db"
	"piano-server/server/domain"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}
func (dao *UserDao) ListUsers() ([]domain.User, error) {
	ins, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	get, err := ins.ListUsers()
	if err != nil {
		return nil, err
	}
	return get, nil
}

func (dao *UserDao) GetUser(id string) (domain.User, error) {
	ins, err := db.GetDB()
	if err != nil {
		return domain.User{}, err
	}
	get, err := ins.GetUser(id)
	if err != nil {
		return domain.User{}, err
	}
	return get, nil
}

func (dao *UserDao) AddUser(data domain.User) (domain.User, error) {
	ins, err := db.GetDB()
	if err != nil {
		return domain.User{}, err
	}
	err = ins.AddUser(data)
	if err != nil {
		return domain.User{}, err
	}
	after, err := ins.GetUser(data.Id)
	if err != nil {
		return domain.User{}, err
	}
	return after, nil
}

func (dao *UserDao) UpdateUser(id string, data domain.User) (domain.User, error) {
	ins, err := db.GetDB()
	if err != nil {
		return domain.User{}, err
	}
	err = ins.UpdateUser(id, data)
	if err != nil {
		return domain.User{}, err
	}
	after, err := ins.GetUser(id)
	if err != nil {
		return domain.User{}, err
	}
	return after, nil
}

func (dao *UserDao) DeleteUser(id string) error {
	ins, err := db.GetDB()
	if err != nil {
		return err
	}
	err = ins.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserDao) GetUserByPhoneNumber(phoneNumber string) (domain.User, error) {
	list, err := dao.ListUsers()
	if err != nil {
		return domain.User{}, err
	}
	for _, user := range list {
		if user.PhoneNumber == phoneNumber {
			return user, nil
		}
	}
	return domain.User{}, errors.NotExistError{Id: phoneNumber, Type: "phone number"}
}

func (dao *UserDao) DeleteUserByPhoneNumber(phoneNumber string) error {
	user, err := dao.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return err
	}
	err = dao.DeleteUser(user.Id)
	if err != nil {
		return err
	}
	return nil
}
