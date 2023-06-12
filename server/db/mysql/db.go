package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"piano-server/config"
	"piano-server/server/domain"
)

type DB struct {
	*gorm.DB
}

func (DB *DB) InitTables() error {
	err := DB.DB.AutoMigrate(&domain.User{})
	if err != nil {
		return err
	}
	return nil
}

func NewDB() (*DB, error) {
	var (
		ip       = config.Get("db.mysql.ip")
		port     = config.GetInt("db.mysql.port")
		userName = config.Get("db.mysql.username")
		password = config.Get("db.mysql.password")
		dbName   = config.Get("db.mysql.db")
		params   = config.Get("db.mysql.params")
		logPath  = config.Get("db.mysql.log")
	)
	var ssh string
	if ip == "" || port <= 0 {
		ssh = ""
	} else {
		ssh = fmt.Sprintf("tcp(%s:%d)", ip, port)
	}
	dsn := fmt.Sprintf("%s:%s@%s/%s%s", userName, password, ssh, dbName, params)
	writer, err := os.Create(logPath)
	if err != nil {
		return nil, err
	}
	dbConn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,  // default size for string fields
		DisableDatetimePrecision:  true, // disable datetime precision, witch not supported before MySQL 5.6
		DontSupportRenameIndex:    true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true, // `change` when rename column, rename column not supported before MuSQL 8, MariaDB
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: logger.New(
			log.New(io.MultiWriter(os.Stdout, writer), "", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
				Colorful: false,
			},
		),
	})
	if err != nil {
		return nil, err
	}
	instance := &DB{dbConn}
	err = instance.InitTables()
	if err != nil {
		return nil, err
	}
	return instance, nil
}
