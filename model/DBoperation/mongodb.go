package DBoperation

import (
	"context"
	"fmt"
	"log"

	"giftCode_04/config"
	"giftCode_04/model/userInfo"
	"go.mongodb.org/mongo-driver/bson"

)
var client ,collection= config.InitMongodb()




func InsertMongo(info userInfo.UserInfo,col string)bool{
	_ ,err := collection.Collection(col).InsertOne(context.TODO(),info)
	if err != nil{
		return false
	}
	return true
}

func FindMongo(key string,value string,col string) (userInfo.UserInfo, error){
	filter := bson.D{{key,value}}
	var result userInfo.UserInfo
	err := collection.Collection(col).FindOne(context.TODO(),filter).Decode(&result)
	if result.Uid == "" && err != nil {
		return result,err
	}
	fmt.Printf("Found a single document: %+v\n", result)
	return result,err
}

func UpdataMongo(keyID string,valueID string,giftKey string,gift string)bool{

	filter := bson.D{{keyID,valueID}}
	update :=  bson.D{
		{"$set", bson.D{
			{giftKey, gift},
		}},
	}

	result ,err := collection.Collection("info").UpdateOne(context.TODO(),filter,update)
	if err != nil {     log.Fatal(err) }
	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return false
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
	return true
}

func ExistId(id string,col string)  bool{

	result,_:= FindMongo("Uid",id,col)
	if result.Uid == ""{
		return false
	}
	return true
}
// DisConnection 断开连接
func DisConnection(){
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}