package router

import (
	"CRM-test/database"
	"CRM-test/structures"
	"CRM-test/utils"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Initialization() *gin.Engine {
	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secretWord"))
	router.Use(sessions.Sessions("session", store))

	router.LoadHTMLGlob("template/*.html")
	router.Static("assets/", "assets")

	router.GET("/", index)
	router.GET("/about", index2)
	router.GET("/settings", index1)
	router.GET("/login", login)
	router.POST("/login", loginPOST)

	return router
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

func loginPOST(c *gin.Context) {
	var data structures.User

	e := c.BindJSON(&data)

	if e != nil {
		utils.Logger.Println(e)
		return
	}

	data.Password, e = utils.Encrypt(data.Password)

	if e != nil {
		utils.Logger.Println(e)
		return
	}

	res := database.Link.QueryRow(`SELECT "id" FROM "User" WHERE "password" = $1 AND "login" = $2`, data.Password, data.Login)

	userID := 0

	e = res.Scan(&userID)

	if e != nil {
		utils.Logger.Println(e)
		return
	}

	if userID > 0 {
		c.JSON(200, true)
	}
}

func login(c *gin.Context) {
	c.HTML(200, "login", nil)
}
