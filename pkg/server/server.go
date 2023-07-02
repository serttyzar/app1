package server

import (
	"app1/internal/app1/endpoint"
	"app1/internal/app1/middleware"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (serv *Server) Start(txt string) error {
	fmt.Println(txt)

	server := echo.New()

	server.Use(middleware.MW)
	server.GET("/", endpoint.New().Action)

	err := server.Start(":8080")

	if err != nil {
		return err
	}

	return nil
}
