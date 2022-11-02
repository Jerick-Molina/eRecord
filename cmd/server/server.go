package server

import (
	"eRecord/internal/db"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

type Server struct {
	router              *gin.Engine
	record              *db.Record
	Status              map[string]string
	Error               string
	AllowAllOrigins     bool
	AllowedHeaders      []string
	AllowedMethods      []string
	AuthorizationHeader string
}

func NewServer(record *db.Record) *Server {
	server := &Server{record: record}
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowAllOrigins: true, AllowHeaders: []string{"Authorization", "projectId", "content-type"},
		AllowMethods: []string{"GET", "POST"},
	}))

	server.router = route
	server.EndpointsHandler()
	defer route.Run("localhost:8080")
	return server
}
func (server *Server) Start(host string) error {
	return server.router.Run(host)
}
func (server *Server) EndpointsHandler() {
	//This is where you init your routes

	api := server.router.Group("/api")
	{
		routeTest_Roles := []string{"Admin"}

		api.POST("/SignIn", server.UserSignIn)
		api.POST("/Create/Account")
		auth := api.Group("/auth", server.AuthorizeToken())
		{
			create := auth.Group("/create")
			{
				createTicketRoleParams := []string{"Admin", "Manager"}
				create.POST("/Ticket", server.RoleAuthorization(createTicketRoleParams, server.CreateTicket))
				create.GET("/Ticket", server.RoleAuthorization(createTicketRoleParams, server.CreateTicketParams))
				createProjectRoleParams := []string{"Admin", "Manager"}
				create.POST("/Project", server.RoleAuthorization(createProjectRoleParams, server.CreateProject))
			}
			auth.POST("/Ticket/Find/User")
			auth.POST("/Ticket/Find/Company")

			projectFindByCompanyRoleParams := []string{"Admin", "Manager"}
			auth.POST("/Project/Find/Company", server.RoleAuthorization(projectFindByCompanyRoleParams, server.FindProjectsByCompanyId))

			projectFindRoleParams := []string{"Admin", "Manager"}
			auth.GET("/Find/Project", server.RoleAuthorization(projectFindRoleParams, server.FindAssociatedToProject))
			auth.POST("/Test", server.RoleAuthorization(routeTest_Roles, server.InitCompany))

			dashboardRoleParams := []string{"Admin", "Manager"}
			auth.GET("/Dashboard", server.RoleAuthorization(dashboardRoleParams, server.FindTicketsAssociatedToCompany))
			//	projectsRoleParams := []string{"Admin", "Manager"}
			auth.GET("/Projects", server.FindProjectsByCompanyId)
			dashboardTicketsRoleParams := []string{"Admin", "Manager"}
			auth.GET("/Dashboard/Tickets", server.RoleAuthorization(dashboardTicketsRoleParams, server.DashboardTickets))
			//	auth.POST("/Test2"), server.InitCompany.Use(server.RoleAuthorization([]string{"Admin", "Owner"})
		}
		//api.POST("/Create_Account", server.CreateAccountRequirements, server.createAccountByCode)
		//	api.POST("/Create_Company", server.CreateAccountWithCompany).Use()
		//	api.POST("/SignIn", server.UserSignIn)

	}
}
