package main

import (
	"goRestExample/api"
	"goRestExample/config"
)

func main() {

	a := api.App{}
	a.Initialize(config.USERNAME, config.PASSWORD, config.DB_NAME)

	a.Run(":" + a.GetPort())
}
