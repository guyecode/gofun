package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main(){
	r := gin.Default()

	// 返回JSON
	r.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hi", "status": http.StatusOK})
	})

	//用struct构造JSON
	r.GET("/msg", func(c *gin.Context) {
		var msg struct {
			Name string `json user`
			Message string
			Number int
		}
		msg.Name = "Lena"
		msg.Message = "hi"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/secure", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// Will output  :   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})


	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// The specific definition of protobuf is written in the testdata/protoexample file.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, data)
	})


	r.Run(":8080")
}