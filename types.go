package express

type SFRouterReq struct {
	Language       string   `json:"language"`
	TrackingType   string   `json:"trackingType"`
	TrackingNumber []string `json:"trackingNumber"`
	MethodType     string   `json:"methodType"`
	CheckPhoneNo   string   `json:"checkPhoneNo"`
}

type SFRouterResp struct {
	ApiErrorMsg   string       `json:"apiErrorMsg"`
	ApiResponseID string       `json:"apiResponseID"`
	ApiResultCode string       `json:"apiResultCode"`
	ApiResultData string       `json:"apiResultData"`
	ApiResultMsg  SFResultData `json:"apiResultMsg"`
}

type (
	SFResultData struct {
		Success   bool        `json:"success"`
		ErrorCode string      `json:"errorCode"`
		ErrorMsg  interface{} `json:"errorMsg"`
		MsgData   MsgData     `json:"msgData"`
	}

	MsgData struct {
		RouteResps []RouteResps `json:"routeResps"`
	}
	RouteResps struct {
		MailNo string   `json:"mailNo"`
		Routes []Routes `json:"routes"`
	}

	Routes struct {
		AcceptAddress       string `json:"acceptAddress"`
		FirstStatusCode     string `json:"firstStatusCode"`
		SecondaryStatusName string `json:"secondaryStatusName"`
		AcceptTime          string `json:"acceptTime"`
		Remark              string `json:"remark"`
		OpCode              string `json:"opCode"`
		SecondaryStatusCode string `json:"secondaryStatusCode"`
		FirstStatusName     string `json:"firstStatusName"`
	}
)
