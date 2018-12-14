package server

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/tockn/tamabus/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	dbx    *sqlx.DB
	engine *gin.Engine
}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
	}
}

func (s *Server) Setup(dbConfPath, env string) error {
	s.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:3000"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept"},
		AllowMethods:     []string{"POST", "PUT", "DELETE"},
	}))

	dbconf, err := NewDBConfigsFromFile(dbConfPath)
	if err != nil {
		return errors.New(fmt.Sprintf("connot open database. %s", err))
	}

	dbx, err := dbconf[env].Open()
	if err != nil {
		return errors.New(fmt.Sprintf("db init failed. %s", err))
	}

	s.dbx = dbx

	s.setRouter()

	return nil
}

func (s *Server) setRouter() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	busController := controllers.BusController{DB: s.dbx, Logger: logger}
	s.engine.GET("/api/bus", busController.GetBuses)
	s.engine.POST("/api/bus", busController.PostGPS)
}

func (s *Server) Run(port string) {
	s.engine.Run(":" + port)
}
