package main

import(
	"github.com/gin-gonic/gin"
	"github.com/Qwerci/article/controllers"
	

)



func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", controllers.Home)

	router.Run(":8000")
}