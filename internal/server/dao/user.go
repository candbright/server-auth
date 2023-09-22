package dao

import (
	"github.com/candbright/server-auth/internal/server/db"
	"github.com/candbright/server-auth/internal/server/db/options"
	"github.com/candbright/server-auth/internal/server/domain"
)

type UserDao struct {
	DB db.DB
}

func NewUserDao(db db.DB) *UserDao {
	return &UserDao{db}
}
func (dao *UserDao) ListUsers() ([]domain.User, error) {
	get, err := dao.DB.ListUsers()
	if err != nil {
		return nil, err
	}
	return get, nil
}

func (dao *UserDao) GetUserById(id string) (domain.User, error) {
	get, err := dao.DB.GetUser(options.WhereId(id))
	if err != nil {
		return domain.User{}, err
	}
	return get, nil
}

func (dao *UserDao) GetUser(opts ...options.Option) (domain.User, error) {
	get, err := dao.DB.GetUser(opts...)
	if err != nil {
		return domain.User{}, err
	}
	return get, nil
}

func (dao *UserDao) AddUser(data domain.User) (domain.User, error) {
	err := dao.DB.AddUser(data)
	if err != nil {
		return domain.User{}, err
	}
	after, err := dao.GetUserById(data.Id)
	if err != nil {
		return domain.User{}, err
	}
	return after, nil
}

func (dao *UserDao) UpdateUser(id string, data domain.User) (domain.User, error) {
	err := dao.DB.UpdateUser(id, data)
	if err != nil {
		return domain.User{}, err
	}
	after, err := dao.GetUserById(id)
	if err != nil {
		return domain.User{}, err
	}
	return after, nil
}

func (dao *UserDao) DeleteUser(id string) error {
	err := dao.DB.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserDao) GetUserByPhoneNumber(phoneNumber string) (domain.User, error) {
	user, err := dao.DB.GetUser(options.Where("phone_number", phoneNumber))
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
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
