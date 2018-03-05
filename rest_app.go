package main

import (
	"goRestExample/api"
)

func main() {

	a := api.App{}
	a.Initialize("root", "admin", "rest_api_example")

	a.Run(":" + a.GetPort())
}
