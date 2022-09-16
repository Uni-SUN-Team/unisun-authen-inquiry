package main

import (
	"log"
	"unisun/api/unisun-authen-management-information/src"
	"unisun/api/unisun-authen-management-information/src/configs/environment"
	"unisun/api/unisun-authen-management-information/src/constants"
)

func main() {
	err := environment.LoadEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	r := src.App()
	port := environment.ENV.App.Port
	if port == "" {
		r.Run(":" + constants.PORT)
	} else {
		r.Run(":" + port)
	}
}
