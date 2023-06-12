package mysql

import (
	"github.com/candbright/go-core/errors"
	"piano-server/server/domain"
)

func (DB *DB) ListUsers() ([]domain.User, error) {
	var result []domain.User
	if err := DB.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (DB *DB) GetUser(id string) (domain.User, error) {
	var result domain.User
	if err := DB.Where("id = ?", id).Take(&result).Error; err != nil {
		return domain.User{}, err
	}
	return result, nil
}

func (DB *DB) AddUser(data domain.User) error {
	var before domain.User
	if err := DB.First(&before, data.Id).Error; err == nil {
		return errors.ExistError{Id: data.Id, Type: "user"}
	}
	if err := DB.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (DB *DB) UpdateUser(id string, data domain.User) error {
	if err := DB.Where("id = ?", id).UpdateColumns(&data).Error; err != nil {
		return err
	}
	return nil
}

func (DB *DB) DeleteUser(id string) error {
	if err := DB.Where("id = ?", id).Delete(&domain.User{}).Error; err != nil {
		return err
	}
	return nil
}
