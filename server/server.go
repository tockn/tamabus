package server

import (
	"errors"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func Setup(s *Server, dbConfPath, env string) error {
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
	return nil
}

func (s *Server) Run(port string) {
	s.engine.Run(":" + port)
}
