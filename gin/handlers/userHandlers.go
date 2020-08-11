package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hubertme/gin-demo/datasource"
	"github.com/hubertme/gin-demo/result"
	"net/http"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	var res []datasource.User
	for _, val := range datasource.AllUsers {
		res = append(res, val)
	}

	c.JSON(http.StatusOK, result.Success(res))
}

func GetUserById(c *gin.Context) {
	idInt64, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, "Need int as id", err.Error()))
		return
	}
	id := int(idInt64)

	userData := datasource.AllUsers[id]
	if userData.ID == 0 {
		c.JSON(http.StatusNotFound, result.Error(404, "User not found", nil))
		return
	}

	c.JSON(http.StatusOK, result.Success(userData))
}

func PostNewUser(c *gin.Context) {
	user := datasource.User{}
	err := c.BindJSON(&user)
	userId := user.ID

	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, "Bad formatted code", err.Error()))
		return
	} else if userId <= 0 {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, "Require id > 0", nil))
		return
	} else if datasource.AllUsers[userId].ID != 0 {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, fmt.Sprintf("User with ID %v is already exist", user.ID), nil))
		return
	}

	datasource.AllUsers[userId] = user
	c.JSON(http.StatusOK, result.Success(user))
}

func UpdateUserById(c *gin.Context) {

}

func DeleteUserById(c *gin.Context) {
	idInt64, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, "Need int as id", err.Error()))
		return
	}
	id := int(idInt64)

	datasource.AllUsers[id] = datasource.User{}
	c.JSON(http.StatusGone, result.Success(nil))
}

func DeleteAllUsers(c *gin.Context) {
	datasource.AllUsers = map[int]datasource.User{}

	c.JSON(http.StatusGone, result.Success(nil))
}
