package handler

import (
	"github.com/gin-gonic/gin"
	service2 "github.com/shinshin8/golang-grpc-client/service"
	"log"
	"net/http"
)

type Handler interface {
	GetEmployee(c *gin.Context)
	ListEmployee(c *gin.Context)
}

type ginHandler struct {
	service service2.Service
}

func NewHandler(service service2.Service) Handler {
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
		for path, route := range routes {
			api.Handle(route.method, path, route.fn)
		}
		return engine
	}
}

func (g *ginHandler) GetEmployee(c *gin.Context) {
	entity, err := g.service.FindEmployee(c.Param("id"))
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
