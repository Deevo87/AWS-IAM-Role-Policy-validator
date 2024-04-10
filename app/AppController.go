package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type Controller struct {
	router *gin.Engine
}

func NewController(router *gin.Engine) *Controller {
	return &Controller{router: router}
}

func (c *Controller) CreateHtml() {
	c.router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}

func (c *Controller) ReceiveJson() {
	c.router.POST("/uploadJsonFile", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"valid": false, "error": err.Error()})
			return
		}

		extension := filepath.Ext(fileHeader.Filename)
		if extension != ".json" {
			c.JSON(http.StatusOK, gin.H{"valid": false, "error": "File is not JSON"})
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"valid": false, "error": err.Error()})
			return
		}

		defer file.Close()
		validator := NewValidatorService()
		isValid, err := validator.Validate(file)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"valid": isValid, "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"valid": isValid, "error": nil})
	})
}
