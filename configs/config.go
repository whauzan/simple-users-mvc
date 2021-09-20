package configs
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-users-mvc/models"
)

var DB *gorm.DB

func InitDB() {
	connection := "root:whr1728@tcp(localhost:3306)/Users?parseTime=True"
	var err error
	DB, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	MigrateDB()
	if err != nil {
		panic(err)
	}
}

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}