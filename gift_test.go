package main

import (
	"fmt"
	"giftCode_04/handle"
	"testing"

	"giftCode_04/util"
)

func Test_GetRandCode(t *testing.T){

	retString := util.GetRandCode(8)
	if len(retString) != 8{
		fmt.Println("Test_GetRandCode error")
	}else {
		fmt.Println("Test_GetRandCode pass")
	}
}
func Test_HandleLogin(t *testing.T){

	retS,retMap := handle.HandleLogin("nccKwM9O")

	if retMap["Gold"] == "600" && retS == "success"{
		fmt.Println("Test_HandleLogin pass")
	}else {
		fmt.Println("Test_HandleLogin error")
	}


}
