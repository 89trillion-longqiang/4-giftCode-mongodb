package handle

import "github.com/gin-gonic/gin"

func AdminCreatGiftcode(c *gin.Context){
	des := c.Query("des")
	GiftNum := c.Query("GN")
	ValidPeriod :=c.Query("VP")
	GiftContent :=c.Query("GC")
	CreatePer := c.Query("CP")
	retMap := HandleAdminCreatGiftcode(des,GiftNum,ValidPeriod,GiftContent,CreatePer)
	c.JSON(200,gin.H{
		"condition":retMap["condition"],
		"GiftCode" : retMap["GiftCode"],
	})
}
func AdminInquireGiftCode(c *gin.Context){
	GiftCode := c.Query("giftCode")
	retMap,infoMap := HadnleAdminInquireGiftCode(GiftCode)
	c.JSON(200, gin.H{
		"condition": retMap["condition"],
		"GiftCode":  GiftCode,
		"data" : infoMap,
	})

}
func Client(c *gin.Context)  {
	GiftCode := c.Query("giftCode")
	userName := c.Query("usr")
	ret := HandleClient(GiftCode,userName)
	c.JSON(200,gin.H{
		"condition" :ret["condition"] ,
		"GiftContent" : ret["GiftContent"],
		"GiftCode" : ret["GiftCode"],
	})
}

func Login(c *gin.Context){
	id := c.Query("id")

	condition,ret := HandleLogin(id)
	c.JSON(200,gin.H{
		"condition" : condition,
		"data" : ret,
	})
}
func VerGiftCode(c *gin.Context) {
	GiftCode := c.Query("giftCode")
	userName := c.Query("usr")
	ret ,_ := HandleVerGiftCode(GiftCode,userName)

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")

	c.Data(200,"application/octet-stream",ret)


}