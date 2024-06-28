package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"polaris-trpc/edu/imau/pb"
	"strconv"
	"trpc.group/trpc-go/trpc-go"
	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
)

func main() {
	r := gin.Default()

	r.GET("/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "id参数错误",
				"id":      id,
			})
			return
		}

		trpc.NewServer()

		resp, err := pb.NewUserServiceClientProxy().GetUserById(trpc.BackgroundContext(), &pb.UserRequest{Id: int32(id)})
		if err != nil {
			log.Println("调用失败", err)
			c.JSON(500, gin.H{
				"message": "err",
				"type":    err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "ok",
			"id":      id,
			"data":    resp,
		})
	})

	address := fmt.Sprintf(":%d", 8099)

	err := r.Run(address)

	if err != nil {
		log.Fatalln("http服务启动失败!")
		return
	}
}
