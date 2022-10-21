package server

import (
	"eRecord/internal/db"
	"eRecord/internal/security"
	"eRecord/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateAccountWithCompany(ctx *gin.Context) {
	var req db.CreateAccountWithCompanyParams
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Err")
		return
	}

	token, err := server.record.CreateStarterAccountTx(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusNoContent, fmt.Sprintln(err))
		return
	}

	ctx.JSON(http.StatusOK, token)
	return
}

type userSignInParams struct {
	Email    string
	Password string
}

func (server *Server) UserSignIn(ctx *gin.Context) {
	var req userSignInParams
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Err")
		return
	}
	req.Password = util.HashPassword(req.Password)
	acc, err := server.record.SignInValidation(ctx, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintln("Invalid Credidentials"))
		return
	}

	token, err := security.CreateAccessToken(acc.Id, acc.Role, acc.CompanyId)
	if err != nil {
		ctx.JSON(http.StatusNoContent, fmt.Sprintln(err))
		return
	}
	ctx.JSON(http.StatusOK, token)
	return
}

func (server *Server) CreateAccountByInviteCode(ctx *gin.Context) {
	var req db.CreateAccountWithCompanyToken

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Err")
		return
	}

	token, err := server.record.CreateAccountAndJoinCompanyTx(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
	return
}
