package server

import (
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/internal/controllers"
	"mom/services/ms-extractor/internal/db"
	"mom/services/ms-extractor/internal/services"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config      *config.Config
	DB          *db.DB
	Controllers controllers.Controllers
}

func (s *Server) Run() {
	//init services
	srvs := services.Services{
		TestService: services.TestService{
			DB: s.DB,
		},
	}
	//init controllers
	s.Controllers = controllers.Controllers{
		Test: controllers.TestController{
			Config: s.Config,
			DB: s.DB,
			Service: srvs.TestService,
		},
	}
	ginEngine := gin.Default()
	//init routes
	s.initRoutes(ginEngine)
	//run server
	ginEngine.Run(s.Config.Server.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func New(conf *config.Config, db *db.DB) {
	web := Server{Config: conf, DB: db}
	web.Run()
}
