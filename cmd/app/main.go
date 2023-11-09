package main

import (
	"log"

	"github.com/prcryx/monster_mall/config"
	"github.com/prcryx/monster_mall/internal/server"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	apiServer := server.NewAPIServer(config.ServerAddr, config.ServerPort, &gorm.DB{})
	apiServer.Run()
}
