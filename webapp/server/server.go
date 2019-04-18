package server

import (
	"errors"
	"fmt"
	"github.com/tockn/tamabus/webapp/models"
	"log"
	"os"
	"time"

	"github.com/tockn/tamabus/webapp/controllers"

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

	s.engine.StaticFile( "/home.html", "./frontend/html/home.html")
	s.engine.StaticFile( "/access.html", "./frontend/html/access.html")
	s.engine.StaticFile( "/congestion.html", "./frontend/html/congestion.html")
	s.engine.StaticFile( "/contact.html", "./frontend/html/contact.html")
	s.engine.StaticFile( "/timetable.html", "./frontend/html/timetable.html")

	s.engine.Static("/js", "./frontend/js")
	s.engine.Static("/images", "./frontend/images")
	s.engine.Static("/css", "./frontend/css")

	busController := controllers.BusController{DB: s.dbx, Logger: logger}
	s.engine.GET("/api/bus", busController.GetBuses)
	s.engine.POST("/api/bus", busController.PostGPS)
	s.engine.POST("/api/bus/image", busController.PostImage)

}

func (s *Server) Run(port string) error {
	go truncater(s.dbx)
	return s.engine.Run(":" + port)
}

func truncater (db *sqlx.DB) {
	for range time.Tick(24 * time.Hour) {
		log.Println("Truncate images!")
		models.TruncateImage(db)
	}
}
