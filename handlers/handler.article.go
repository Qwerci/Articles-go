package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Qwerci/article/controllers"
)

func ShowIndexPage(c *gin.Context){
	articles := controllers.GetAllArticles()

	// Call the HTML method of the context to render a template
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
			"payload": articles,
		})
}