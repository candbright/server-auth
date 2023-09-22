//go:build wireinject

package api

import (
	"github.com/candbright/server-auth/internal/server/dao"
	"github.com/candbright/server-auth/internal/server/db"
	"github.com/candbright/server-auth/internal/server/service"
	"github.com/google/wire"
)

func InitializeRegisterService() (*service.RegisterService, error) {
	wire.Build(db.GetDB, service.NewRegisterService, dao.NewRegisterDao, dao.NewUserDao)
	return nil, nil
}
