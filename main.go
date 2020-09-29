package main

import (
	"go-clean/config"
	"go-clean/controller"
	"go-clean/handler"
	"go-clean/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.Connect()

	repo := repository.NewKelasRepository(db)
	entity := controller.NewKelasEntity(repo)
	api := r.Group("/api")
	handler.KelasHandlerFunc(api, entity)

	r.Run()
}
