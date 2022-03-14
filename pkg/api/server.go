package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Serve(){
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "templates/assets")
	router.Static("/images", "templates/images")
	base := router.Group("/")
	base.GET("/", getIndex)
	base.GET("/elements", getElements)

	apiV1 := router.Group("/api/v1/")
	{
		apiV1.GET("person", getPerson)
		apiV1.GET("person/:id", getPersonById)
		apiV1.POST("person", addPerson)
		apiV1.PUT("person/:id", updatePerson)
		apiV1.DELETE("person/:id", deletePerson)
		apiV1.OPTIONS("options", options)
	}
	log.Panic(router.Run(":9090"))
}

func getIndex(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", nil)
}

func getElements(c *gin.Context)  {
	c.HTML(http.StatusOK, "elements.html", nil)
}

func options(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message":"options"})
}

func addPerson(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message":"addPerson"})
}

func getPerson(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message":"getPerson"})
}

func getPersonById(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message":"getPersonById:"+c.Param("id")})
}

func updatePerson(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message":"updatePerson:"+c.Param("id")})
}

func deletePerson(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message":"deletePerson:"+c.Param("id")})
}
