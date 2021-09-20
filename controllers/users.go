package controllers
import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"simple-users-mvc/lib"
)

func GreetMessage(echoContext echo.Context) error {
	return echoContext.JSON(http.StatusOK, "Welcome to RentABook API by Whauzan")
}

func GetAllUsersController(echoContext echo.Context) error {
	users, err := lib.GetAllUsers()
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"stats" : "err",
			"message" : err,
		})
	}
	return echoContext.JSON(http.StatusOK, users)
}

func PostUsersController(echoContext echo.Context) error {
	var userReq RequestUser
	echoContext.Bind(&userReq)
	res, err := lib.PostUsers(userReq.toModel())
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"stats": "err",
			"message": err,
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{} {
		"stats": "success",
		"data": newResponse(*res),
	})
}

func GetUserByIDController(echoContext echo.Context) error {
	ID := echoContext.Param("id")
	intID, _ := strconv.Atoi(ID)
	user, e := lib.GetUserByID(intID)
	if e != nil {
		return echoContext.JSON(http.StatusNotFound, map[string]interface{} {
			"message": e.Error(),
		})
	}
	return echoContext.JSON(http.StatusOK, user)
}

func UpdateUserController(echoContext echo.Context) error {
	ID := echoContext.Param("id")
	intID, _ := strconv.Atoi(ID)

	var userReq RequestUser
	echoContext.Bind(&userReq)
	res, err := lib.UpdateBook(userReq.toModel(),intID)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{} {
			"stats":   "err",
			"messages": err,
		})
	}
	res, _ = lib.GetUserByID(intID)
	return echoContext.JSON(http.StatusOK, res)
}

func DeleteUserController(echoContext echo.Context) error {
	ID := echoContext.Param("id")
	intID, _ := strconv.Atoi(ID)
	_, e := lib.DeleteBook(intID)
	if e != nil {
		return echoContext.JSON(http.StatusNotFound, map[string]interface{} {
			"message": e.Error(),
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{} {
		"message": "deleted",
	})
}