package main

import (
	"GoStudy/WebFrameworks/Echo/api"
	"GoStudy/WebFrameworks/Echo/api/helper/logging"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	// TODO: init viper, logger
	startServer()
}

func startServer() {
	e := echo.New()

	logging.InitLogging()
	fmt.Println("Starting Server at localhost:8080 ...")
	e.HideBanner = true
	api.Init(e)
	err := e.Start(":8080")

	if err != nil {
		fmt.Println("Error while starting server : ", err)
	}

}
