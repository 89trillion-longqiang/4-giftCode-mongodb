package router

import (
	"giftCode_04/handle"
	"github.com/gin-gonic/gin"
)
func SetUpRount() *gin.Engine  {
	r := gin.Default()
	c1 := r.Group("/giftCode")

	c1.GET("/adminCreatGiftcode", handle.AdminCreatGiftcode)
	c1.GET("/admininquireGiftCode", handle.AdminInquireGiftCode)
	c1.GET("/client", handle.Client)
	c1.GET("/login",handle.Login)
	c1.GET("/VerGiftCode",handle.VerGiftCode)

	return r
}