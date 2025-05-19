package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/y4ns3/sarkor-test/internal/rest/handler"
)

type Server struct {
	gin     *gin.Engine
	handler *handler.ProductHandler
}

func NewServer(h *handler.ProductHandler) *Server {
	engine := gin.Default()
	return &Server{
		handler: h,
		gin:     engine,
	}
}

func (s *Server) Run(url string) error {

	s.gin.Use(gin.Recovery())
	s.gin.Use(gin.Logger())
	productGroup := s.gin.Group("/products")
	{
		productGroup.GET("/:id", s.handler.GetProduct)
		productGroup.GET("/", s.handler.GetProducts)
		productGroup.POST("/", s.handler.CreateProduct)
		productGroup.PUT("/:id", s.handler.UpdateProduct)
		productGroup.DELETE("/:id", s.handler.DeleteProduct)
	}
	return s.gin.Run(url)
}
