package user

import (
	"net/http"
	"strconv"

	"github.com/SamanNsr/bookstore_users-api/domain/users"
	"github.com/SamanNsr/bookstore_users-api/services"
	"github.com/SamanNsr/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string)(int64, *errors.RestErr)  {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func GetUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}


func UpdateUser(c *gin.Context)  {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	result, err := services.UserService.UpdateUser(user)
	if err !=nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context)  {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UserService.DeleteUser(userId); err !=nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status":"deleted"})
}

func Search(c *gin.Context)  {
	status := c.Query("status")

	users, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

