package express

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
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
	Key       string `json:"key"`       // apiKey
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

type ExpressCompanyRes struct {
	ExpressCompanyList []ExpressCompany `json:"expressCompanyList"`
	Success            bool             `json:"success"`
	TraceID            string           `json:"traceId"` // 注意：表格中是"traceld"但可能是拼写错误，应为"traceId"
}

type ExpressCompany struct {
	CompanyName string `json:"companyName"` // 快递公司名称，示例："圆通快递"
	CpCode      string `json:"cpCode"`      // 快递公司编码，示例："YTO"
}

func NewAPISpaceService(token string) *apiSpaceService {
	return &apiSpaceService{
		token: token,
		url:   "https://eolink.o.apispace.com/wlgj1/paidtobuy_api",
	}
}

// SearchRoutes 搜索物流信息
func (s *apiSpaceService) SearchRoutes(ctx context.Context, req *APISpaceReq) (*APISpaceResp, error) {

	api := "/trace_search"

	if req == nil {
		return nil, errors.New("input cannot be nil")
	}
	if req.MailNo == "" {
		return nil, errors.New("MailNo cannot be empty")
	}
	if req.CpCode == "" {
		company, err := s.MailDiscern(ctx, req.MailNo, req.Key)
		if err != nil {
			return nil, err
		}
		if company == nil {
			return nil, errors.New("mail discern failed")
		}
		if company.ExpressCompanyList == nil {
			return nil, errors.New("mailNo false")
		}
		if len(company.ExpressCompanyList) > 0 {
			req.CpCode = company.ExpressCompanyList[0].CpCode
		}
	}

	// 再次确认
	if req.CpCode == "" {
		return nil, errors.New("CpCode cannot be empty")
	}

	if req.OrderType == "" {
		req.OrderType = "desc" // 降序 最新的会在上面
	}

	jsonData, _ := json.Marshal(req)
	// 转成buffer
	body := bytes.NewBuffer(jsonData)

	client := http.Client{}
	request, _ := http.NewRequest("POST", s.url+api, body)
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

// MailDiscern 识别物流单号所属的快递公司(免费)
// @company 快递公司信息 (如果ExpressCompanyList返回为nil，则可能说明单号有问题)
func (s *apiSpaceService) MailDiscern(ctx context.Context, mailNo, key string) (company *ExpressCompanyRes, err error) {
	api := "/mail_discern"
	if mailNo == "" {
		return nil, errors.New("MailNo cannot be empty")
	}
	data := map[string]interface{}{"mailNo": mailNo}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", s.url+api, bytes.NewBuffer(jsonData))

	req.Header.Add("X-APISpace-Token", key)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	uData := &ExpressCompanyRes{}
	json.Unmarshal(all, &uData)

	return uData, nil
}
