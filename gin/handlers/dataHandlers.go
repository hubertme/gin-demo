package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hubertme/gin-demo/datasource"
	"github.com/hubertme/gin-demo/result"
	"net/http"
	"strconv"
)

func GetAllTestUsers(c *gin.Context) {
	var res []datasource.User
	for _, val := range datasource.AllUsers {
		res = append(res, val)
	}

	c.JSON(http.StatusOK, result.Success(res))
}

func GetTestUserById(c *gin.Context) {
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

func PostNewTestUser(c *gin.Context) {
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

func updateUserData(oldUser datasource.User, newUser datasource.User) datasource.User {
	if newUser.FullName != "" {
		oldUser.FullName = newUser.FullName
	}

	if newUser.PhoneNumber != "" {
		oldUser.PhoneNumber = newUser.PhoneNumber
	}

	return oldUser
}

func UpdateTestUserById(c *gin.Context) {
	idInt64, err := strconv.ParseInt(c.Param("id"), 10, 64)
	id := int(idInt64)

	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, "Need int as id", err.Error()))
		return
	} else if id <= 0 {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, "Require id > 0", nil))
		return
	} else if datasource.AllUsers[id].ID == 0 {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, fmt.Sprintf("User with ID %v does not found", id), nil))
		return
	}

	user := datasource.User{}
	parseErr := c.BindJSON(&user)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, "Require user data", parseErr.Error()))
		return
	}
	datasource.AllUsers[id] = updateUserData(datasource.AllUsers[id], user)

	c.JSON(http.StatusOK, result.Success(datasource.AllUsers[id]))
}

func DeleteTestUserById(c *gin.Context) {
	idInt64, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, "Need int as id", err.Error()))
		return
	}
	id := int(idInt64)

	delete(datasource.AllUsers, id)
	c.JSON(http.StatusGone, result.Success(nil))
}

func DeleteAllTestUsers(c *gin.Context) {
	datasource.AllUsers = map[int]datasource.User{}

	c.JSON(http.StatusGone, result.Success(nil))
}
