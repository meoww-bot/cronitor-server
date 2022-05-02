package handler

import (
	"cronitor-server/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {

	env := util.CheckEnv()

	c.String(http.StatusOK, "hello world\n")

	c.String(http.StatusOK, env)
}
