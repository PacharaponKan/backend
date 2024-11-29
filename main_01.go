package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main0() {
	r := gin.Default()

	//Get Method No Path
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "BackEnd Api Ok",
		})
	})

	//--------------------Employee Method------------------------
	//Get Method With Path
	r.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET EMPLOYEE METHOD",
		})
	})

	//Post Method
	r.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST EMPLOYEE METHOD",
		})
	})

	//Put Method
	r.PUT("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT EMPLOYEE METHOD",
		})
	})

	//Delete Method
	r.DELETE("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE EMPLOYEE METHOD",
		})
	})

	//-------------------------------------------------

	//--------------------Customer Method------------------------
	//Get Method With Path
	r.GET("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET CUSTOMER METHOD",
		})
	})

	//Post Method
	r.POST("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST CUSTOMER METHOD",
		})
	})

	//Put Method
	r.PUT("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT CUSTOMER METHOD",
		})
	})

	//Delete Method
	r.DELETE("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE CUSTOMER METHOD",
		})
	})

	//-------------------------------------------------
	r.Run(":8082") // listen and serve on 0.0.0.0:8082
}
