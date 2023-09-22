package api

import (
	"github.com/candbright/go-log/log"
	"github.com/candbright/server-auth/internal/server/service"
	"time"
)

var RegisterService *service.RegisterService

func initServices() {
	var err error
	for {
		RegisterService, err = InitializeRegisterService()
		if err == nil {
			break
		}
		log.Debug(err)
		time.Sleep(10 * time.Second)
	}
}
