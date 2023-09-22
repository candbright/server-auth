package db

import (
	"github.com/candbright/go-log/log"
	"github.com/candbright/server-auth/internal/config"
)

var instance DB

func GetDB() (DB, error) {
	if instance != nil {
		return instance, nil
	}
	newDB, err := NewDB()
	if err != nil {
		log.Debug(err)
		for i := 0; i < config.GetInt("db.retry"); i++ {
			newDB, err = NewDB()
			if err == nil {
				instance = newDB
				return instance, nil
			}
		}
		return nil, err
	}
	instance = newDB
	return instance, nil
}
