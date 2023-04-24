package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Qwerci/article/models"
)

func Home(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
	})
}
func GetAllArticles()[]models.Article{
	return models.ArticleList
}
