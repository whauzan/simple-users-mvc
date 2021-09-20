package controllers
import (
	"time"
	"simple-users-mvc/models"
)

type RequestUser struct {
	ID int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Age string `json:"age" form:"age"`
	Sex string `json:"sex" form:"sex"`
	ClientID int `json:"client_id" form:"client_id"`
}

func (req *RequestUser) toModel() models.User {
	return models.User{
		ID: req.ID,
		Name: req.Name,
		Age: req.Age,
		Sex: req.Sex,
		ClientID: req.ClientID,
	}
}

type ResponseUser struct {
	ID int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Age string `json:"age" form:"age"`
	Sex string `json:"sex" form:"sex"`
	ClientID int `json:"client_id" form:"client_id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}

func newResponse(modelUser models.User) ResponseUser {
	return ResponseUser {
		ID: modelUser.ID,
		Name: modelUser.Name,
		Age: modelUser.Age,
		Sex: modelUser.Sex,
		ClientID: modelUser.ClientID,
		CreatedAt: modelUser.CreatedAt,
		UpdatedAt: modelUser.UpdatedAt,
		DeletedAt: modelUser.DeletedAt.Time,
	}
}

func newResponseArray(modelUser []models.User) []ResponseUser {
	var response []ResponseUser
	for _, val := range modelUser {
		response = append(response, newResponse(val))
	}
	return response
}