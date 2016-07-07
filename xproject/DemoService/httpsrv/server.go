package httpserver

import (
	"github.com/jinzhu/gorm"
	"jasonghw/goframework"
	"jasonghw/goframework/log"
)

type Server struct {
	framework.Frameworker
	db *gorm.DB
	// gingine *gin.Engine
}

func New(configFile string, si framework.ServiceInfo) (*Server, error) {
	fwk, err := framework.Default(&si, configFile, nil)
	if err != nil {
		return nil, err
	}

	s := &Server{
		Frameworker: fwk,
	}

	err = fwk.RegisterGinRoute(s.registerRoute)
	if err != nil {
		return nil, err
	}

	s.db = fwk.GetMysql()
	return s, nil
}

func (s *Server) Serve(addr string) {
	err := s.Frameworker.Serve(addr)
	if err != nil {
		log.Error("s.Serve(%s) error:%v", addr, err)
	}
}
