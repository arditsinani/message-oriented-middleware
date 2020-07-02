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
		TestS: services.TestS{
			DB: s.DB,
		},
		WSS: services.WSS{
			DB: s.DB,
		},
		PrjS: services.PrjS{
			DB: s.DB,
		},
	}
	//init controllers
	s.Controllers = controllers.Controllers{
		Test: controllers.TestCtrl{
			Config:  s.Config,
			DB:      s.DB,
			Service: srvs.TestS,
		},
		WS: controllers.WSCtrl{
			Config:  s.Config,
			DB:      s.DB,
			Service: srvs.WSS,
		},
		Prj: controllers.PrjCtrl{
			Config:  s.Config,
			DB:      s.DB,
			Service: srvs.PrjS,
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
