package service

import (
	"errors"
	"github.com/candbright/go-core/rest"
	"github.com/gin-gonic/gin"
	"piano-server/server/domain"
	"piano-server/server/internal"
)

func AllocateCode() string {
	code := internal.RandomRegisterCode()
	//TODO send the message
	return code
}

func GetRegisterCode(context *gin.Context) {
	rest.GET(context, func() (interface{}, error) {
		phoneNumber := context.Query("phone_number")
		code := AllocateCode()
		err := registerDao.SetRegisterCode(phoneNumber, code)
		if err != nil {
			return nil, err
		}
		return code, nil
	})
}

func RegisterUser(context *gin.Context) {
	rest.POST[domain.User](context, func(receive domain.User) (interface{}, error) {
		phoneNumber := context.Query("phone_number")
		code := context.Query("code")
		codeSaved, err := registerDao.GetRegisterCode(phoneNumber)
		if err != nil {
			return nil, err
		}
		if code != codeSaved {
			return nil, errors.New("invalid code")
		}
		if receive.Id == "" {
			receive.Id = internal.UUID()
		}
		if receive.PhoneNumber == "" {
			return nil, errors.New("please set the phone number")
		}
		if receive.Name == "" {
			receive.Name = "pianist0"
		}
		if receive.Password == "" {
			receive.Password = "1q2w3e4R!"
		}
		add, err := userDao.AddUser(receive)
		if err != nil {
			return nil, err
		}
		return add, nil
	})
}

func GetUserInfo(context *gin.Context) {
	rest.GET(context, func() (interface{}, error) {
		var err error
		var user domain.User
		id := context.Query("id")
		if id != "" {
			user, err = userDao.GetUser(id)
		}
		phoneNumber := context.Query("phone_number")
		if phoneNumber != "" {
			user, err = userDao.GetUserByPhoneNumber(phoneNumber)
		}
		if err != nil {
			return nil, err
		}
		return user, nil
	})
}

func UpdateUserInfo(context *gin.Context) {
	rest.PUT[domain.User](context, func(receive domain.User) (interface{}, error) {
		var err error
		var before domain.User
		var after domain.User
		id := context.Query("id")
		if id != "" {
			after, err = userDao.UpdateUser(id, receive)
			if err != nil {
				return nil, err
			}
		}
		phoneNumber := context.Query("phone_number")
		if phoneNumber != "" {
			before, err = userDao.GetUserByPhoneNumber(phoneNumber)
			if err != nil {
				return nil, err
			}
			after, err = userDao.UpdateUser(before.Id, receive)
			if err != nil {
				return nil, err
			}
		}
		return after, nil
	})
}

func DeleteUserInfo(context *gin.Context) {
	rest.DELETE(context, func() (interface{}, error) {
		id := context.Query("id")
		if id != "" {
			err := userDao.DeleteUser(id)
			if err != nil {
				return nil, err
			}
		}
		phoneNumber := context.Query("phone_number")
		if phoneNumber != "" {
			err := userDao.DeleteUserByPhoneNumber(phoneNumber)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
}

func UserLogin(context *gin.Context) {

}
