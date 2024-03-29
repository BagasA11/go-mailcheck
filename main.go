package main

import (
	"fmt"

	"os"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Email struct {
	Email string `form:"email" json:"email"`
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

	/*
		//email domain
		if err := checkmail.ValidateHost(req.Email); err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{
				"massage": err.Error(),
			})
			return
		}
	*/

	c.JSON(200, gin.H{
		"massage": "valid",
		"mail":    req.Email,
	})
}

func JsonMail(c *gin.Context) {
	var req Email
	if err := c.ShouldBindJSON(&req); err != nil {
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

	c.JSON(200, gin.H{
		"massage": "valid",
		"mail":    req.Email,
	})
}

func main() {
	//load env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	/*
		router.LoadHTMLGlob("template/*")
		router.GET("/", func(ctx *gin.Context) {
			ctx.HTML(200, "index.html", nil)
		})
		router.POST(fmt.Sprintf("/%s", os.Getenv("ENDPOINT")), email)
	*/
	router.LoadHTMLGlob("template/*")

	g1 := router.Group("/api")
	{
		g1.GET("/", func(ctx *gin.Context) {
			ctx.HTML(200, "index.html", nil)
		})
		//using html form
		g1.POST(fmt.Sprintf("/%s", os.Getenv("ENDPOINT")), email)
		//using json request
		g1.POST(fmt.Sprintf("/%s/json", os.Getenv("ENDPOINT")), JsonMail)
	}

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
