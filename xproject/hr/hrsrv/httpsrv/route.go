package httpserver

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) registerRoute(engine *gin.Engine) error {
	engine.GET("/echo", s.echoHanlder)
	engine.GET("/groupcount", s.groupCountHanlder)
	return nil
}
