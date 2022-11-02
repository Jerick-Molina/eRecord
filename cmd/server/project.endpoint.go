package server

import (
	"eRecord/internal/db"
	"eRecord/internal/security"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateProject(ctx *gin.Context) {
	var args db.CreateProjectTxParams

	if err := ctx.BindJSON(&args); err != nil {
		ctx.JSON(http.StatusNotAcceptable, err)
		return
	}
	claims, err := security.GetJwtMap(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	t := fmt.Sprint(claims["companyId"])
	companyId, err := strconv.ParseInt(t, 32, 64)

	args.AssociatedCompany = int(companyId)
	if err := server.record.CreateProjectTx(ctx, args); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, "Project created")

}

func (server *Server) FindProjectsByCompanyId(ctx *gin.Context) {

	claims, err := security.GetJwtMap(ctx)
	t := fmt.Sprint(claims["companyId"])
	companyId, err := strconv.ParseInt(t, 32, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(companyId)
	projects, err := server.record.FindProjectByAssociatedCompanyTx(ctx, int(companyId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(projects)
	ctx.JSON(http.StatusOK, projects)
	return
}

func (server *Server) SingleProjectDashboard(ctx *gin.Context) {

}
