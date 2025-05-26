package express

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/linbe-ff/express-go-sdk/consts"
	"time"
)

type kuaiDi100 struct {
	Key      string // 客户授权key
	Secret   string // 秘钥
	Customer string
	UserId   string
}

func NewKuaiDi100(Key, Secret, Customer, UserId string) *kuaiDi100 {
	return &kuaiDi100{
		Key:      Key,
		Secret:   Secret,
		Customer: Customer,
		UserId:   UserId,
	}
}

func (k *kuaiDi100) AddressResolution(req *AddressResolutionParam) (*AddressResolutionRes, error) {

	param := AddressResolutionParam{
		Content: req.Content,
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	res, err := k.DoRequest(t, paramStr, consts.ADDRESS_RESOLUTION_URL)

	if err != nil {
		return nil, errors.New("请求失败")
	}

	if res == "" {
		return nil, errors.New("返回內容为空")
	}

	// 解析到AddressResolutionRes
	var addressResolutionRes AddressResolutionRes
	err = json.Unmarshal([]byte(res), &addressResolutionRes)
	if err != nil {
		return nil, err
	}
	return &addressResolutionRes, nil
}

type AddressResolutionParam struct {
	Content  string
	Image    string
	ImageUrl string
	PdfUrl   string
	HtmlUrl  string
}

type (
	AddressResolutionRes struct {
		Code                  int                   `json:"code"`
		AddressResolutionData AddressResolutionData `json:"data"`
		Message               string                `json:"message"`
		Time                  int                   `json:"time"`
		Success               bool                  `json:"success"`
	}

	AddressResolutionData struct {
		TaskId string                    `json:"taskId"`
		Result []AddressResolutionResult `json:"result"`
	}

	AddressResolutionResult struct {
		Content string               `json:"content"`
		Mobile  []string             `json:"mobile"`
		Name    string               `json:"name"`
		Address string               `json:"address"`
		Xzq     AddressResolutionXzq `json:"xzq"`
	}

	AddressResolutionXzq struct {
		FullName   string `json:"fullName"`
		Province   string `json:"province"`
		City       string `json:"city"`
		District   string `json:"district"`
		SubArea    string `json:"subArea"`
		ParentCode string `json:"parentCode"`
		Code       string `json:"code"`
		Level      int    `json:"level"`
	}
)
