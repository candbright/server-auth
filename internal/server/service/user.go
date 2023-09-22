package service

import (
	"errors"
	"github.com/candbright/go-core/rest"
	"github.com/candbright/server-auth/internal/server/dao"
	"github.com/candbright/server-auth/internal/server/domain"
	"github.com/candbright/server-auth/internal/server/uuid"
	"github.com/gin-gonic/gin"
)

type RegisterService struct {
	RegisterDao *dao.RegisterDao
	UserDao     *dao.UserDao
}

func NewRegisterService(registerDao *dao.RegisterDao, userDao *dao.UserDao) *RegisterService {
	return &RegisterService{
		registerDao,
		userDao,
	}
}

func (s *RegisterService) AllocateCode() string {
	code := uuid.RandomRegisterCode()
	//TODO send the message
	return code
}

func (s *RegisterService) GetRegisterCode(context *gin.Context) {
	rest.GET(context, func() (interface{}, error) {
		phoneNumber := context.Query("phone_number")
		codeCached, _ := s.RegisterDao.GetRegisterCode(phoneNumber)
		if codeCached != "" {
			return codeCached, nil
		}
		code := s.AllocateCode()
		err := s.RegisterDao.SetRegisterCode(phoneNumber, code)
		if err != nil {
			return nil, err
		}
		return code, nil
	})
}

func (s *RegisterService) RegisterUser(context *gin.Context) {
	rest.POST[domain.User](context, func(receive domain.User) (interface{}, error) {
		phoneNumber := context.Query("phone_number")
		code := context.Query("code")
		codeSaved, err := s.RegisterDao.GetRegisterCode(phoneNumber)
		if err != nil {
			return nil, err
		}
		if code != codeSaved {
			return nil, errors.New("invalid code")
		}
		if receive.Id == "" {
			receive.Id = uuid.UUID()
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
		add, err := s.UserDao.AddUser(receive)
		if err != nil {
			return nil, err
		}
		return add, nil
	})
}

func (s *RegisterService) GetUserInfo(context *gin.Context) {
	rest.GET(context, func() (interface{}, error) {
		var err error
		var user domain.User
		id := context.Query("id")
		if id != "" {
			user, err = s.UserDao.GetUserById(id)
		}
		phoneNumber := context.Query("phone_number")
		if phoneNumber != "" {
			user, err = s.UserDao.GetUserByPhoneNumber(phoneNumber)
		}
		if err != nil {
			return nil, err
		}
		return user, nil
	})
}

func (s *RegisterService) UpdateUserInfo(context *gin.Context) {
	rest.PUT[domain.User](context, func(receive domain.User) (interface{}, error) {
		var err error
		var before domain.User
		var after domain.User
		id := context.Query("id")
		if id != "" {
			after, err = s.UserDao.UpdateUser(id, receive)
			if err != nil {
				return nil, err
			}
		}
		phoneNumber := context.Query("phone_number")
		if phoneNumber != "" {
			before, err = s.UserDao.GetUserByPhoneNumber(phoneNumber)
			if err != nil {
				return nil, err
			}
			after, err = s.UserDao.UpdateUser(before.Id, receive)
			if err != nil {
				return nil, err
			}
		}
		return after, nil
	})
}

func (s *RegisterService) DeleteUserInfo(context *gin.Context) {
	rest.DELETE(context, func() (interface{}, error) {
		id := context.Query("id")
		if id != "" {
			err := s.UserDao.DeleteUser(id)
			if err != nil {
				return nil, err
			}
		}
		phoneNumber := context.Query("phone_number")
		if phoneNumber != "" {
			err := s.UserDao.DeleteUserByPhoneNumber(phoneNumber)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
}

func (s *RegisterService) UserLogin(context *gin.Context) {

}
