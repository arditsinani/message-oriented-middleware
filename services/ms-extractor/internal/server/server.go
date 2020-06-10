package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"mom/services/ms-extractor/config"
	"mom/services/ms-extractor/internal/controllers"
)

type Server struct {
	Config		config.Config
	Mongo 		*mongo.Client
	Controllers controllers.Controllers
}

func (s *Server) Run() {
	//init controllers
	s.Controllers = controllers.Controllers{Test: controllers.TestController{Config: s.Config, Mongo: s.Mongo}}
	ginEngine := gin.Default()
	//init routes
	s.initRoutes(ginEngine)
	//run server
	ginEngine.Run(s.Config.Server.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func New(conf config.Config,client *mongo.Client) {
	web := Server{Config: conf, Mongo: client}
	web.Run()
}
