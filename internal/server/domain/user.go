package domain

import "time"

type User struct {
	Id          string    `gorm:"column:id;primaryKey" json:"id"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	Name        string    `gorm:"column:name" json:"name"`
	Password    string    `gorm:"column:password" json:"password"`
	Sex         string    `gorm:"column:sex" json:"sex"`
	Email       string    `gorm:"column:email" json:"email"`
	CreateAt    time.Time `gorm:"column:create_at" json:"create_at"`
	UpdateAt    time.Time `gorm:"column:update_at" json:"update_at"`
	ExternalIds
}

func (user User) TableName() string {
	return "user"
}
