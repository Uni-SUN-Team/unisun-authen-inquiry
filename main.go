package main

import (
	"log"
	"unisun/api/auth-processor-api/src"
	"unisun/api/auth-processor-api/src/configs/environment"
	"unisun/api/auth-processor-api/src/constants"
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
