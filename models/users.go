package models
import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey" json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Age      int    `json:"age" form:"age"`
	Sex      string `json:"sex" form:"sex"`
	ClientID int    `json:"client_id" form:"client_id"`
}