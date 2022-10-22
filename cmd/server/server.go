package server

import (
	"eRecord/internal/db"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

type Server struct {
	router          *gin.Engine
	record          *db.Record
	Status          map[string]string
	Error           string
	AllowAllOrigins bool
	AllowedHeaders  []string
	AllowedMethods  []string
}

func NewServer(record *db.Record) *Server {
	server := &Server{record: record}
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", "projectId", "content-type"},
		AllowMethods:    []string{"GET", "POST"},
	}))

	server.router = route
	server.EndpointsHandler()
	defer route.Run("localhost:8080")
	return server
}
func (server *Server) Start(host string) error {
	return server.router.Run(host)
}
func (servers *Server) EndpointsHandler() {
	//This is where you init your routes

	api := servers.router.Group("/api")
	{

		auth_Api := api.Group("/auth")
		{
			auth_Api.POST("/SignIn")
		}
		//api.POST("/Create_Account", servers.CreateAccountRequirements, servers.createAccountByCode)
		api.POST("/Create_Company", servers.CreateAccountWithCompany).Next()
		api.POST("/SignIn", servers.UserSignIn)
	}
}
