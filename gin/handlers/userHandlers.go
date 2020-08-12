package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hubertme/gin-demo/database"
	"github.com/hubertme/gin-demo/datasource"
	"github.com/hubertme/gin-demo/result"
	"net/http"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	res := []datasource.User{}

	results, err := database.DB.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusBadRequest, result.DevError(err.Error()))
		return
	}

	for results.Next() {
		var user datasource.User

		err = results.Scan(&user.ID, &user.FullName, &user.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, result.DevError(err.Error()))
			return
		}

		res = append(res, user)
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

	statement, sqlErr := database.DB.Prepare("SELECT * FROM users WHERE id = ?")
	if sqlErr != nil {
		c.JSON(http.StatusBadRequest, result.DevError(sqlErr.Error()))
		return
	}

	results, qErr := statement.Query(id)
	if qErr != nil {
		c.JSON(http.StatusInternalServerError, result.DevError(qErr.Error()))
		return
	}

	for results.Next() {
		var user datasource.User

		err = results.Scan(&user.ID, &user.FullName, &user.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, result.DevError(err.Error()))
			panic(err.Error())
		}

		c.JSON(http.StatusOK, result.Success(user))
		return
	}

	c.JSON(http.StatusOK, result.Success(nil))
}

func PostNewUser(c *gin.Context) {
	user := datasource.User{}
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, "Bad formatted code", err.Error()))
		return
	}

	statement, sqlErr := database.DB.Prepare("INSERT INTO users (full_name, phone_number) VALUE (?, ?)")
	if sqlErr != nil {
		c.JSON(http.StatusBadRequest, result.DevError(sqlErr.Error()))
		return
	}

	results, qErr := statement.Query(user.FullName, user.PhoneNumber)
	if qErr != nil {
		c.JSON(http.StatusInternalServerError, result.DevError(qErr.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success(results))
}

func DeleteUserById(c *gin.Context) {
	idInt64, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, "Need int as id", err.Error()))
		return
	}
	id := int(idInt64)

	statement, sqlErr := database.DB.Prepare("DELETE FROM users WHERE id = ?")
	if sqlErr != nil {
		c.JSON(http.StatusBadRequest, result.DevError(sqlErr.Error()))
		return
	}

	results, qErr := statement.Query(id)
	if qErr != nil {
		c.JSON(http.StatusInternalServerError, result.DevError(qErr.Error()))
		return
	}

	c.JSON(http.StatusGone, result.Success(results))
}

func DeleteAllUsers(c *gin.Context) {
	results, qErr := database.DB.Query("DELETE FROM users")
	if qErr != nil {
		c.JSON(http.StatusInternalServerError, result.DevError(qErr.Error()))
		return
	}

	c.JSON(http.StatusGone, result.Success(results))
}

func UpdateUserById(c *gin.Context) {
	idInt64, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error(result.BAD_FORMATTED_REQUEST, "Need int as id", err.Error()))
		return
	}
	id := int(idInt64)

	user := datasource.User{}
	err = c.BindJSON(&user)
	user.ID = id

	statement, sqlErr := database.DB.Prepare("UPDATE users SET full_name = IF(LENGTH(?) = 0, full_name, ?), phone_number = IF(LENGTH(?) = 0, phone_number, ?) WHERE id = ?")
	if sqlErr != nil {
		c.JSON(http.StatusBadRequest, result.DevError(sqlErr.Error()))
		return
	}

	results, qErr := statement.Query(user.FullName, user.FullName, user.PhoneNumber, user.PhoneNumber, user.ID)
	if qErr != nil {
		c.JSON(http.StatusInternalServerError, result.DevError(qErr.Error()))
		return
	}

	c.JSON(http.StatusAccepted, result.Success(results))
}
