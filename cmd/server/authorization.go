package server

import (
	"eRecord/internal/security"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) RoleAuthorization(ValidRoles []string, fn func(*gin.Context)) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		claims, err := security.GetJwtMap(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		role := fmt.Sprint(claims["role"])

		for i := 0; i < len(ValidRoles); i++ {
			if ValidRoles[i] == role || role == "Owner" {
				fn(ctx)
				return
			}
		}

		ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
}

func (server *Server) AuthorizeToken() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var token string

		token = ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, "No token found")
			ctx.Abort()
		}

		claims, err := security.TokenReader(token)

		if err != nil {

			ctx.JSON(http.StatusUnauthorized, err.Error())
			ctx.Abort()

		}

		//Sets users claims
		ctx.Set("claims", claims)

		ctx.Next()
	}
}
