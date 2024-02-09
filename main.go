package main

import (
	"github.com/NonokaM/WebDesignHub-API/controller"
	"github.com/NonokaM/WebDesignHub-API/db"
	"github.com/NonokaM/WebDesignHub-API/repository"
	"github.com/NonokaM/WebDesignHub-API/router"
	"github.com/NonokaM/WebDesignHub-API/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
