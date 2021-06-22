package ctrl

import (
	"giftCode_04/handle"
	"github.com/gin-gonic/gin"
)

func AdminCreatGiftcode(c *gin.Context){
	des := c.Query("des")
	GiftNum := c.Query("GN")
	ValidPeriod :=c.Query("VP")
	GiftContent :=c.Query("GC")
	CreatePer := c.Query("CP")
	if len(des)== 0 && len(GiftNum)== 0 && len(ValidPeriod)== 0 && len(GiftContent)== 0 && len(CreatePer)== 0 {
		c.JSON(200, gin.H{
			"condition": "error",
		})
		return
	}
		retMap := handle.HandleAdminCreatGiftcode(des,GiftNum,ValidPeriod,GiftContent,CreatePer)
	c.JSON(200,gin.H{
		"condition":retMap["condition"],
		"GiftCode" : retMap["GiftCode"],
	})
}
func AdminInquireGiftCode(c *gin.Context){
	GiftCode := c.Query("giftCode")
	if len(GiftCode) == 0  {
		c.JSON(200,gin.H{
			"condition":"error",
		})
		return
	}
	retMap,infoMap := handle.HadnleAdminInquireGiftCode(GiftCode)
	c.JSON(200, gin.H{
		"condition": retMap["condition"],
		"GiftCode":  GiftCode,
		"data" : infoMap,
	})

}
func Client(c *gin.Context)  {
	GiftCode := c.Query("giftCode")
	userName := c.Query("usr")

	if len(GiftCode) == 0 && len(userName) == 0{
		c.JSON(200,gin.H{
			"condition":"error",
		})
		return
	}
	ret := handle.HandleClient(GiftCode,userName)
	c.JSON(200,gin.H{
		"condition" :ret["condition"] ,
		"GiftContent" : ret["GiftContent"],
		"GiftCode" : ret["GiftCode"],
	})
}

func Login(c *gin.Context){
	id := c.Query("id")

	if len(id) == 0{
		c.JSON(200,gin.H{
			"condition":"error",
		})
		return
	}

	condition,ret := handle.HandleLogin(id)
	c.JSON(200,gin.H{
		"condition" : condition,
		"data" : ret,
	})
}
func VerGiftCode(c *gin.Context) {
	GiftCode := c.Query("giftCode")
	userName := c.Query("usr")

	if len(GiftCode) == 0 && len(userName) == 0{
		c.JSON(200,gin.H{
			"condition":"error",
		})
		return
	}

	ret ,_ := handle.HandleVerGiftCode(GiftCode,userName)

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")

	c.Data(200,"application/octet-stream",ret)


}