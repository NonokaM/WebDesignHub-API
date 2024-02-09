package main

import (
	"fmt"

	"github.com/NonokaM/WebDesignHub-API/db"
	"github.com/NonokaM/WebDesignHub-API/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{})
}
