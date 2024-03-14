package main

import (
	"fmt"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

type Email struct {
	Email string `form:"email"`
}

func email(c *gin.Context) {
	var req Email
	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//email format
	if err := checkmail.ValidateFormat(req.Email); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"massage": err.Error(),
		})
		return
	}

	//email domain
	if err := checkmail.ValidateHost(req.Email); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"massage": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"massage": "valid",
		"mail":    req.Email,
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})
	router.POST("/email", email)
	router.Run(":8080")
}
