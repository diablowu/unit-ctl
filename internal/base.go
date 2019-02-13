package internal

import "fmt"

const (
	ManageEndpoint = "https://aip.baidubce.com/rpc/2.0/unit"
)

type BaseCliArgs struct {
	Action string
}

type ResponseBase struct {
	ErrorMsg  string `json:"error_msg"`
	ErrorCode int    `json:"error_code"`
}

func (rsp ResponseBase) Success() bool {
	return rsp.ErrorCode == 0
}

func (rsp ResponseBase) ErrorReason() string {
	return fmt.Sprintf("code:%d, reason:%s", rsp.ErrorCode, rsp.ErrorMsg)
}

type ServiceBaseInfo struct {
	AccessToken string
	Endpoint    string
}

func (bs ServiceBaseInfo) GetFullUrl(method string) string {
	return fmt.Sprintf("%s/%s?access_token=%s", bs.Endpoint, method, bs.AccessToken)
}
