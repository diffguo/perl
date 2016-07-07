package httpserver

import (
	"github.com/gin-gonic/gin"
	"jasonghw/xproject/hr/hrsrv/logic"
	"net/http"
)

func (s *Server) groupCountHanlder(c *gin.Context) {
	c.Request.ParseForm()
	uid := c.Request.Form.Get("uid")
	count, err := logic.GetUidGroupCount(s.db, uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"count": count})
	}
}
