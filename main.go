package main
import (
	"github.com/whauzan/simple-users-mvc/configs"
	"github.com/whauzan/simple-users-mvc/routes"
)

func main() {
	configs.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
}