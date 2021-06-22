package router

import (
	"giftCode_04/ctrl"
	"github.com/gin-gonic/gin"
)
func SetUpRount() *gin.Engine  {
	r := gin.Default()
	c1 := r.Group("/giftCode")

	c1.GET("/adminCreatGiftcode", ctrl.AdminCreatGiftcode)
	c1.GET("/admininquireGiftCode", ctrl.AdminInquireGiftCode)
	c1.GET("/client", ctrl.Client)
	c1.GET("/login", ctrl.Login)
	c1.GET("/VerGiftCode", ctrl.VerGiftCode)

	return r
}