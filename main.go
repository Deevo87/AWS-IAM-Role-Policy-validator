package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zadanie_remitly/app"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/templates", "./templates")
	controller := app.NewController(router)
	controller.CreateHtml()
	controller.ReceiveJson()

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error during starting app: ", err)
		return
	}

}
