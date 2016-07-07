// Package classification User API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
// 	   Contact: liudanking<liudanking@example.com> http://somewhere.com
//
//     Consumes:
//     - application/x-www-form-urlencoded
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"jasonghw/goframework"
	"jasonghw/goframework/log"
	"jasonghw/xproject/hr/hrsrv/httpsrv"
	"net/http"
	_ "net/http/pprof"
)

var (
	gConfig     *ServiceConfig
	gConfigFile *string
)

var (
	// the variables will be set value at compile time from golang build ldflags
	g_Service   = "HR"
	g_Version   = "version"
	g_BuildDate = "build_date"
	g_CommitID  = "commit_id"
)

func initConfig() (*ServiceConfig, error) {
	gConfigFile = flag.String("c", "hr.conf", "config file")
	flag.Parse()
	return loadConfig(*gConfigFile)
}

func startPprof() {
	go func() {
		log.Error("%v", http.ListenAndServe("localhost:6060", nil))
	}()
}

func startHttpServer(configFile string, addr string) (*gin.Engine, error) {
	si := framework.ServiceInfo{
		Service:   g_Service,
		Version:   g_Version,
		BuildDate: g_BuildDate,
		CommitID:  g_CommitID,
	}

	httpServer, err := httpserver.New(configFile, si)
	if err != nil {
		return nil, err
	}

	httpServer.Serve(addr)
	return httpServer.GetGingine(), nil
}

func main() {
	var err error
	gConfig, err = initConfig()
	if err != nil {
		fmt.Printf("initConfig error:%v", err)
		return
	}

	startPprof()

	_, err = startHttpServer(*gConfigFile, gConfig.Custom.HttpListenAddr)
	if err != nil {
		fmt.Printf("HttpServer error:%v", err)
		return
	}

	log.Debug("exit")
}
