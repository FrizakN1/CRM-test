package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*.html")
	router.Static("assets/", "assets")
	router.GET("/", index)
	router.GET("/about", index2)
	router.GET("/settings", index1)
	router.GET("/login", login)
	_ = router.Run("127.0.0.1:8080")
}

func index(c *gin.Context) {
	c.HTML(200, "index", nil)
}

func index1(c *gin.Context) {
	c.HTML(200, "index1", nil)
}

func index2(c *gin.Context) {
	c.HTML(200, "index2", nil)
}

func login(c *gin.Context) {
	c.HTML(200, "login", nil)
}
