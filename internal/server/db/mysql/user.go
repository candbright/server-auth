package mysql

import (
	"github.com/candbright/go-core/errors"
	"github.com/candbright/server-auth/internal/server/db/options"
	"github.com/candbright/server-auth/internal/server/domain"
	"gorm.io/gorm"
	"time"
)

func (DB *DB) ListUsers() ([]domain.User, error) {
	var result []domain.User
	if err := DB.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (DB *DB) GetUserById(id string) (domain.User, error) {
	return DB.GetUser(options.WhereId(id))
}

func (DB *DB) GetUser(opts ...options.Option) (domain.User, error) {
	var result domain.User
	if err := DB.Options(opts...).Take(&result).Error; err != nil {
		return domain.User{}, err
	}
	return result, nil
}

func (DB *DB) AddUser(data domain.User) error {
	var before domain.User
	err := DB.Where(domain.User{Id: data.Id}).First(&before).Error
	if err == nil {
		return errors.ExistError{Id: data.Id, Type: "user"}
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	err = DB.Where(domain.User{PhoneNumber: data.PhoneNumber}).First(&before).Error
	if err == nil {
		return errors.ExistError{Id: data.PhoneNumber, Type: "phone number"}
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	data.CreateAt = time.Now()
	data.UpdateAt = time.Now()
	if err := DB.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (DB *DB) UpdateUser(id string, data domain.User) error {
	data.UpdateAt = time.Now()
	if err := DB.Where(domain.User{Id: id}).UpdateColumns(&data).Error; err != nil {
		return err
	}
	return nil
}

func (DB *DB) DeleteUser(id string) error {
	if err := DB.Where(domain.User{Id: id}).Delete(&domain.User{}).Error; err != nil {
		return err
	}
	return nil
}c
