package db

import (
	"github.com/candbright/server-auth/internal/server/db/mysql"
	"github.com/candbright/server-auth/internal/server/db/options"
	"github.com/candbright/server-auth/internal/server/db/redis"
	"github.com/candbright/server-auth/internal/server/domain"
)

type Register interface {
	GetRegisterCode(phoneNumber string) (string, error)
	SetRegisterCode(phoneNumber string, code string) error
}

type User interface {
	ListUsers() ([]domain.User, error)
	GetUser(filter ...options.Option) (domain.User, error)
	AddUser(data domain.User) error
	UpdateUser(id string, user domain.User) error
	DeleteUser(id string) error
}

type DB interface {
	User
	Register
}

type db struct {
	User
	Register
}

func NewDB() (DB, error) {
	mysqlDB, err := mysql.NewDB()
	if err != nil {
		return nil, err
	}
	redisDB, err := redis.NewDB()
	if err != nil {
		return nil, err
	}
	return &db{
		User:     mysqlDB,
		Register: redisDB,
	}, nil
}
