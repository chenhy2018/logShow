package handler

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type GetLogCountOutput struct {
	module string 
	count int
}

func GetName(c *gin.Context) {
	//name := c.Param("name")
	//firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s", lastname)
}

func GetLogCount(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Origin", "*")
	//c.Header("Access-Control-Allow-M")

	log.Println("===GetLogCount===")
	c.JSON(http.StatusOK, gin.H{
		"count_log": 5})
}
