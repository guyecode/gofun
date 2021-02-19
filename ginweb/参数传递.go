package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main()  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// URL参数
	// http://127.0.0.1:8080/user/john
	r.GET("/user/:name", func(c *gin.Context){
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// curl http://127.0.0.1:8080/user/john/run/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message, c.FullPath())
	})

	// GET参数
	// curl http://127.0.0.1:8080/welcome\?firstname\=Gu\&lastname\=Ye
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// POST参数
	// curl -X POST -d "name=guye&message=hello&message=hello" http://127.0.0.1:8080/post\?id\=123
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "1")
		name := c.PostForm("name")
		message := c.DefaultPostForm("message", "nothing happen")
		c.String(http.StatusOK, "id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	// 传递array和map参数
	// curl -X POST -d "names[first]=gu&names[second]=ye" "http://127.0.0.1:8080/map?ids=1234&ids=455"
	r.POST("/map", func(c *gin.Context) {
		ids := c.QueryArray("ids")
		names := c.PostFormMap("names")
		fmt.Printf("ids: %v, names: %v\n", ids, names)
		c.String(http.StatusOK, strings.Join(ids, ", ") + fmt.Sprintf(" first: %s, second: %s", names["first"], names["second"]))
	})

	r.Run()
}
