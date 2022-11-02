package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var InitCompany string

func (server *Server) InitCompany(ctx *gin.Context) {

	ctx.JSON(http.StatusAccepted, "yee")
}
