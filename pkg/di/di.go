package di

import (
	bootserver "myproject/pkg/boot"
	"myproject/pkg/config"
	db "myproject/pkg/database"
	"myproject/pkg/user"
)

func InitiallizeEvent(conf config.Config) (*bootserver.ServerHttp, error) {

	DB := db.ConnectPGDB(conf)

	userRepository := user.NewRepository(DB)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	serverHttp := bootserver.NewServerHttp(*userHandler)

	return serverHttp, nil

}
