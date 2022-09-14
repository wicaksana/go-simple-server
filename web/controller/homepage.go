package controller

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

type LogMessage struct {
	Message string `form:"message" json:"message" xml:"message"  binding:"required"`
}

func Router(r *gin.Engine) {
	r.GET("/", index)
	r.POST("/info/new", newInfoMsgHandler)
	r.POST("/warning/new", newWarningMsgHandler)
	r.POST("/error/new", newErrorMsgHandler)
}

func index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/index.html",
		gin.H{},
	)
}

func newInfoMsgHandler(c *gin.Context) {
	logMessage := &LogMessage{}

	if err := c.ShouldBindJSON(&logMessage); err != nil {
		log.Error().Msg(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Print message as stdout
	log.Info().
		Str("severity", "INFO").
		Msg(logMessage.Message)

	c.JSON(http.StatusOK, gin.H{"status": "INFO message received"})
}

func newWarningMsgHandler(c *gin.Context) {
	logMessage := &LogMessage{}

	if err := c.ShouldBindJSON(&logMessage); err != nil {
		log.Error().Msg(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Print message as stdout
	log.Info().
		Str("severity", "WARNING").
		Msg(logMessage.Message)

	c.JSON(http.StatusOK, gin.H{"status": "WARNING message received"})
}

func newErrorMsgHandler(c *gin.Context) {
	logMessage := &LogMessage{}

	if err := c.ShouldBindJSON(&logMessage); err != nil {
		log.Error().Msg(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Print message as stderr
	log.Error().
		Str("severity", "ERROR").
		Msg(logMessage.Message)

	c.JSON(http.StatusOK, gin.H{"status": "ERROR message received"})
}
