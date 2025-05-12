package express

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestSfRouter(t *testing.T) {
	sfReq := SFRouterReq{
		Language:       "0",
		TrackingType:   "1",
		TrackingNumber: []string{"SF316xxxxxxxxxxx"},
		MethodType:     "1",
		CheckPhoneNo:   "0714",
	}
	msgDataBytes, _ := json.Marshal(sfReq)

	// 获取yaml文件
	//viper.SetConfigFile("./config.yaml")
	//// 读取配置文件
	//if err := viper.ReadInConfig(); err != nil {
	//	t.Fatalf("无法读取配置文件: %v", err)
	//}
	//ccode := viper.GetString("sf.customerCode")
	//checkCode := viper.GetString("sf.checkCode")
	sfService := NewExpressService("xxxxxxxxxxxx", "xxxxxxxxxxxxxxxxxxxxxxx")

	msgData := string(msgDataBytes)
	timestamp := time.Now().Unix()

	msgDigest := sfService.GenerateMsgDigest(msgData, timestamp)
	fmt.Printf("SfSearchRoutes->msgDigest:%+v\n", msgDigest)

	sfService.SfSearchRoutes(msgData, msgDigest, timestamp)
}
