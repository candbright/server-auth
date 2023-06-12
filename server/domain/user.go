package domain

type User struct {
	Id          string `gorm:"column:id;primaryKey" json:"id"`
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number"`
	Name        string `gorm:"column:name" json:"name"`
	Password    string `gorm:"column:password" json:"password"`
	Sex         int    `gorm:"column:sex" json:"sex"`
	Email       string `gorm:"column:email" json:"email"`
	CreateAt    string `gorm:"column:create_at" json:"create_at"`
	UpdateAt    string `gorm:"column:update_at" json:"update_at"`
	ExternalIds
}

func (user User) TableName() string {
	return "user"
}
