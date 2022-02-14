package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler interface {
	GetEmployee(c *gin.Context)
	ListEmployee(c *gin.Context)
}

type ginHandler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &ginHandler{service: service}
}

func SetRoute(handler Handler) func(*gin.Engine) *gin.Engine {
	return func(engine *gin.Engine) *gin.Engine {
		routes := map[string]struct {
			fn     func(c *gin.Context)
			method string
		}{
			"/:id": {
				fn:     handler.GetEmployee,
				method: http.MethodGet,
			},
			"/list": {
				fn:     handler.ListEmployee,
				method: http.MethodGet,
			},
		}
		api := engine.Group("/company")
		api.GET(":/id", handler.GetEmployee)
		for path, route := range routes {
			api.Handle(route.method, path, route.fn)
		}
		return engine
	}
}

func (g *ginHandler) GetEmployee(c *gin.Context) {
	entity, err := g.service.FindEmployee(c.Query("id"))
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, entity)
}

func (g *ginHandler) ListEmployee(c *gin.Context) {
	entity, err := g.service.ListEmployee()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, entity)
}
