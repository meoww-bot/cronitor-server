package handler

import (
	"cronitor-server/db"
	"cronitor-server/lib"
	"cronitor-server/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	UserName := c.Param("username")

	if UserName == "" {

		c.String(http.StatusOK, "Example: /v3/user/{username}")

	} else {

		user := lib.User{
			Username: UserName,
			ApiKey:   util.Md5encode(UserName),
		}

		_, err := db.AddUser(user)

		if Error(c, err) {
			return //exit
		}

		c.JSON(200, gin.H{
			"message": "success",
			"data":    user,
		})

	}
}

func GetUserInfo(c *gin.Context) {

	Username := c.Param("username")

	if Username == "" {

		c.String(http.StatusOK, "Example: /v3/user/{username}")

	} else {

		user := new(lib.User)

		err := db.GetUser(Username, user)

		if err != nil {
			c.JSON(404, gin.H{
				"error":   true,
				"message": "not found, POST /v3/user/:username to create user",
			})
			return

		}

		c.JSON(200, gin.H{
			"message": "success",
			"data":    user,
		})

	}

}
