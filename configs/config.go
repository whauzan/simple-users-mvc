package configs
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-users-mvc/models"
)

var DB *gorm.DB

func InitDB() {
	connectionString := "root:whr1728@tcp(0.0.0.0:3306)/Users?parseTime=True"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}