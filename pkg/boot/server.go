package bootserver

import (
	"myproject/pkg/config"
	"myproject/pkg/user"

	"github.com/labstack/echo/v4"
)

type ServerHttp struct {
	engine *echo.Echo
}

func NewServerHttp(userHandler user.Handler) *ServerHttp {
	engine := echo.New()
	userHandler.MountRoutes(engine)

	return &ServerHttp{engine}
}
func (s *ServerHttp) Start(conf config.Config) {
	s.engine.Start(conf.Host + ":" + conf.ServerPort)
}
