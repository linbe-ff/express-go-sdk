package express

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type expressService struct {
	CustomerCode string // 顾客编码
	CheckCode    string // 校验码
}

func NewExpressService(customerCode, checkCode string) *expressService {
	return &expressService{
		CustomerCode: customerCode,
		CheckCode:    checkCode,
	}
}

func (t *expressService) GenerateMsgDigest(msgData string, timestamp int64) string {
	str := url.QueryEscape(msgData + strconv.FormatInt(timestamp, 10) + t.CheckCode)
	fmt.Printf("generateMsgDigest->str:%+v\n", str)

	h := md5.New()
	h.Write([]byte(str))
	msgDigest := base64.StdEncoding.EncodeToString(h.Sum(nil))
	fmt.Printf("generateMsgDigest->msgDigest:%+v\n", msgDigest)

	return msgDigest
}

func (t *expressService) SfSearchRoutes(msgData string, msgDigest string, timestamp int64) (*SFRouterResp, error) {

	var data = url.Values{}
	data.Add("partnerID", t.CustomerCode)
	data.Add("requestID", GenerateToken())
	data.Add("serviceCode", "EXP_RECE_SEARCH_ROUTES")
	data.Add("timestamp", strconv.FormatInt(timestamp, 10))
	data.Add("msgDigest", msgDigest)
	data.Add("msgData", url.QueryEscape(msgData))

	dataStr, _ := url.QueryUnescape(data.Encode())
	fmt.Printf("ExpressService->SfSearchRoutes->dataStr:%+v\n", dataStr)

	request, _ := http.NewRequest("POST", "https://bspgw.sf-express.com/std/service", strings.NewReader(dataStr))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)

	var result = &SFRouterResp{}

	json.Unmarshal(bytes, &result)
	json.Unmarshal([]byte(result.ApiResultData), &result.ApiResultMsg)

	if result.ApiErrorMsg != "" {
		return result, errors.New(result.ApiErrorMsg)
	}
	return result, nil
}
