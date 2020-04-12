package routers

import (
	"net/http"
	"reportGameErr/models"

	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func RegiserService(engi *gin.Engine) {
	engine = engi
	api := engine.Group("/api")
	api.GET("/log", logHandle)
	api.POST("/update", updateHandler)
}

func logHandle(c *gin.Context) {
	module := c.Query("module")
	if len(module) == 0 {
		c.JSON(http.StatusOK, gin.H{"c": 0, "m": "module field is empty", "d": gin.H{}})
		return
	}

	line := c.Query("line")
	if len(line) == 0 {
		c.JSON(http.StatusOK, gin.H{"c": 0, "m": "line field is empty", "d": gin.H{}})
		return
	}

	column := c.Query("column")
	if len(column) == 0 {
		c.JSON(http.StatusOK, gin.H{"c": 0, "m": "column field is empty", "d": gin.H{}})
		return
	}

	info := c.Query("info")
	if len(info) == 0 {
		c.JSON(http.StatusOK, gin.H{"c": 0, "m": "info field is empty", "d": gin.H{}})
		return
	}

	cfg := models.NewLogConfig()
	cfg.SetModule(module)
	cfg.SetLine(line)
	cfg.SetColumn(column)
	cfg.SetInfo(info)

	if er := cfg.Save(c.ClientIP()); er != nil {
		c.JSON(http.StatusOK, gin.H{"c": 0, "m": er.Error(), "d": gin.H{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"c": 1, "m": "ok", "d": gin.H{}})
}

func updateHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"c": 1, "m": "ok", "d": gin.H{}})
}