package express

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

/*
官网：https://www.apispace.com
*/
type apiSpaceService struct {
	token string
	url   string
}

type APISpaceReq struct {
	CpCode    string `json:"cpCode"`    // 快递公司编码
	MailNo    string `json:"mailNo"`    // 运单号
	Tel       string `json:"tel"`       // 收件人电话
	OrderType string `json:"orderType"` // 降序升序
}

type APISpaceResp struct {
	TraceId        string         `json:"traceId"`
	TraceId1       string         `json:"trace_id"`
	Success        bool           `json:"success"`
	LogisticsTrace LogisticsTrace `json:"logisticsTrace"`
}

type LogisticsTrace struct {
	TheLastTime              string                     `json:"theLastTime"`
	CpCode                   string                     `json:"cpCode"` // 快递公司编码
	CpUrl                    string                     `json:"cpUrl"`  // 快递公司官网
	TakeTime                 string                     `json:"takeTime"`
	LogisticsStatusDesc      string                     `json:"logisticsStatusDesc"`
	LogisticsTraceDetailList []LogisticsTraceDetailList `json:"logisticsTraceDetailList"`
	MailNo                   string                     `json:"mailNo"`
	TheLastMessage           string                     `json:"theLastMessage"`       // 最后一条物流信息
	CpMobile                 string                     `json:"cpMobile"`             // 快递公司电话
	LogisticsCompanyName     string                     `json:"logisticsCompanyName"` // 快递公司名称
	Courier                  string                     `json:"courier"`              // 快递员
	CourierPhone             string                     `json:"courierPhone"`         // 快递员电话
	LogisticsStatus          string                     `json:"logisticsStatus"`      // 物流状态
}

type LogisticsTraceDetailList struct {
	AreaCode           string `json:"areaCode"`
	AreaName           string `json:"areaName"`
	SubLogisticsStatus string `json:"subLogisticsStatus"`
	Time               int64  `json:"time"`
	LogisticsStatus    string `json:"logisticsStatus"`
	Desc               string `json:"desc"`
	Courier            string `json:"courier,omitempty"`
	CourierPhone       string `json:"courierPhone,omitempty"`
}

func NewAPISpaceService(token string) *apiSpaceService {
	return &apiSpaceService{
		token: token,
		url:   "https://eolink.o.apispace.com/wlgj1/paidtobuy_api/trace_search",
	}
}

func (s *apiSpaceService) SearchRoutes(ctx context.Context, req *APISpaceReq) (*APISpaceResp, error) {
	if req == nil {
		return nil, errors.New("input cannot be nil")
	}
	if req.CpCode == "" {
		return nil, errors.New("CpCode cannot be empty")
	}
	if req.MailNo == "" {
		return nil, errors.New("MailNo cannot be empty")
	}

	if req.OrderType == "" {
		req.OrderType = "desc" // 降序 最新的会在上面
	}

	jsonData, _ := json.Marshal(req)
	// 转成buffer
	body := bytes.NewBuffer(jsonData)

	client := http.Client{}
	request, _ := http.NewRequest("POST", s.url, body)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-APISpace-Token", s.token)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// 解析 response
	var resp APISpaceResp
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
