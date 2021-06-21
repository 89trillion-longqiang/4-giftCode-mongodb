package DBoperation

import (
	"context"
	"fmt"

	"giftCode_04/config"
	"giftCode_04/model/gift"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client  = config.InitClient() ///声明一个全局的rdb变量

func HashSet(g gift.Gift) error{
	ctx := context.Background()

	if g.IsEmpty() {
		err := fmt.Errorf("gift struct is empty")
		return err
	}
	if ExistsKey(g.GiftCode) {
		err := fmt.Errorf("GiftCode exists")
		return err
	}
	gmap := map[string]interface{} {
		"giftCode" : g.GiftCode,
		"Description" : g.Description,
		"GiftNum" :g.GiftNum,
		"ValidPeriod" :g.ValidPeriod,
		"GiftContent" :g.GiftContent,
		"CreatePer" :g.CreatePer,
		"CreatTime" :g.CreatTime,
		"AvailableNum":g.AvailableNum,
		"ClaimList":g.ClaimList,
	}
	_ , err := rdb.HSet(ctx,g.GiftCode,gmap).Result()
	if err!=nil {
		return err
	}
	return nil
}

func HashSetMap(gmap map[string]string) error {

	gm := map[string]interface{}{
		"giftCode" : gmap["GiftCode"],
		"Description" : gmap["Description"],
		"GiftNum" :gmap["GiftNum"],
		"ValidPeriod" :gmap["ValidPeriod"],
		"GiftContent" :gmap["GiftContent"],
		"CreatePer" :gmap["CreatePer"],
		"CreatTime" :gmap["CreatTime"],
		"AvailableNum":gmap["AvailableNum"],
		"ClaimList":gmap["ClaimList"],
	}
	ctx := context.Background()

	_ , err := rdb.HSet(ctx,gmap["GiftCode"],gmap).Result()
	if err!=nil {
		return err
	}
	fmt.Println(gm)
	return nil
}
///HGETALL
func HashGetAll(giftCode string) ( a map[string]string, err error){
	ctx := context.Background()

	return rdb.HGetAll(ctx,giftCode).Result()
}
///判断key值是否已经存在
func ExistsKey(key string) bool {
	ctx := context.Background()

	n, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	if n > 0 {
		return true
	} else {
		return false
	}
}