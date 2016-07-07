package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Just a echo rsp
// swagger:response someRsp
type someRsp struct {
	// Response message
	//
	// Required: true
	Msg string
	// Optional field
	Word string
}

// Hi api
// swagger:route GET /hi say hi
//
// Say hi to client
//
// Just Say hi
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: someRsp
func (s *Server) echoHanlder(c *gin.Context) {

	c.Request.ParseForm()
	hi := c.Request.Form.Get("hi")
	c.JSON(http.StatusOK, gin.H{"rsp": "hi:" + hi})
}
