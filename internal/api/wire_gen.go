// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"github.com/candbright/server-auth/internal/server/dao"
	"github.com/candbright/server-auth/internal/server/db"
	"github.com/candbright/server-auth/internal/server/service"
)

// Injectors from wire.go:

func InitializeRegisterService() (*service.RegisterService, error) {
	registerDao := dao.NewRegisterDao()
	dbDB, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	userDao := dao.NewUserDao(dbDB)
	registerService := service.NewRegisterService(registerDao, userDao)
	return registerService, nil
}
