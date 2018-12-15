package main

import (
	"fmt"
	"github.com/go-apibox/api"
	"os"
)

func main() {
	app, err := api.NewApp()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	app.Run(apiRoutes)
}
