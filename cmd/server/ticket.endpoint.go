package server

import (
	"eRecord/internal/db"
	"eRecord/internal/security"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateTicket(ctx *gin.Context) {
	var tkt db.Ticket
	if err := ctx.BindJSON(&tkt); err != nil {
		ctx.JSON(http.StatusNotAcceptable, err)
		return
	}

	claims, err := security.GetJwtMap(ctx)
	t := fmt.Sprint(claims["userId"])
	userId, err := strconv.ParseInt(t, 32, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c := fmt.Sprint(claims["companyId"])
	companyId, err := strconv.ParseInt(c, 32, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := server.record.CreateTicketAssignedToProject(ctx, tkt, int(companyId), int(userId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, "OK")
	return
}
func (server *Server) FindTicketsAssociatedToCompany(ctx *gin.Context) {

	claims, err := security.GetJwtMap(ctx)
	t := fmt.Sprint(claims["companyId"])
	companyId, err := strconv.ParseInt(t, 32, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tkts, err := server.record.FindTicketByAssociatedCompanyTx(ctx, int(companyId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, tkts)
	return
}
func (server *Server) DashboardTickets(ctx *gin.Context) {

	claims, err := security.GetJwtMap(ctx)
	t := fmt.Sprint(claims["companyId"])
	companyId, err := strconv.ParseInt(t, 32, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err, array := server.record.FindTicketsDashboardTx(ctx, int(companyId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, array)
	return
}

func (server *Server) CreateTicketParams(ctx *gin.Context) {

	claims, err := security.GetJwtMap(ctx)
	t := fmt.Sprint(claims["companyId"])
	companyId, err := strconv.ParseInt(t, 32, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err, array := server.record.CreateTicketDashboardTx(ctx, int(companyId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(array)
	ctx.JSON(http.StatusOK, array)
	return
}

func (server *Server) FindAssociatedToProject(ctx *gin.Context) {

	projectIdHeader := ctx.GetHeader("projectId")

	projectId, err := strconv.ParseInt(projectIdHeader, 32, 64)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, err)
		return
	}
	claims, err := security.GetJwtMap(ctx)
	t := fmt.Sprint(claims["companyId"])
	companyId, err := strconv.ParseInt(t, 32, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err, array := server.record.FindSingleProjectTx(ctx, int(companyId), int(projectId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, array)
	return
}

// func (server *Server) FindTicketAssociatedToUser(ctx *gin.Context) {

// 	if err := ctx.BindJSON(&args); err != nil {
// 		ctx.JSON(http.StatusNotAcceptable, err)
// 		return
// 	}

// }
