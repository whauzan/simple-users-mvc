package main
import (
	"simple-users-mvc/configs"
	"simple-users-mvc/routes"
)

func main() {
	configs.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
}