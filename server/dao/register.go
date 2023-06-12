package dao

import (
	"piano-server/server/db"
)

type RegisterDao struct {
}

func NewRegisterDao() *RegisterDao {
	return &RegisterDao{}
}

func (dao *RegisterDao) GetRegisterCode(phoneNumber string) (string, error) {
	ins, err := db.GetDB()
	if err != nil {
		return "", err
	}
	get, err := ins.GetRegisterCode(phoneNumber)
	if err != nil {
		return "", err
	}
	return get, nil
}

func (dao *RegisterDao) SetRegisterCode(phoneNumber string, code string) error {
	ins, err := db.GetDB()
	if err != nil {
		return err
	}
	err = ins.SetRegisterCode(phoneNumber, code)
	if err != nil {
		return err
	}
	return nil
}
