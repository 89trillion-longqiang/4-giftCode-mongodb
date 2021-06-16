package main

import (
	"giftCode_04/config"
	"giftCode_04/router"
)

func main()  {
	config.InitClient()
	config.InitMongodb()
	r:=router.SetUpRount()
	r.Run(":8080")
}
