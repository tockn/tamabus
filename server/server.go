package server

import (
	"errors"

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
		Engine: gin.Default(),
	}
}

func Setup(s *Server, dbConfPath, env string) error {
	s.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:3000"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept"},
		AllowMethods:     []string{"PUT", "DELETE"},
	}))

	dbconf, err := NewDBConfigsFromFile(dbConfPath)
	if err != nil {
		return errors.New("connot open database. %s", err)
	}

	dbx, err := dbconf.Open(env)
	if err != nil {
		return errors.New("db init failed. %s", err)
	}

	s.dbx = dbx

}

func (s *Server) Run(port string) {
	s.Engine.Run(":" + port)
}
