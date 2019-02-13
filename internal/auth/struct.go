package auth

import (
	"fmt"
	"github.com/levigross/grequests"
	"log"
)

const (
	DefaultGrantType = "client_credentials"
)

type Credential struct {
	AppKey    *string
	SecretKey *string
	GrantType *string
}

// 生成access_token获取的query string
func (cred Credential) Auth() *AccessTokenResponse {

	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=%s&client_id=%s&client_secret=%s", *cred.GrantType, *cred.AppKey, *cred.SecretKey)
	if resp, err := grequests.Get(url, nil); err == nil && resp.Ok {
		token := AccessTokenResponse{}
		if err := resp.JSON(&token); err == nil {
			return &token
		} else {
			log.Fatalf("解析access token数据错误, %s", err)
		}
	} else {
		log.Fatalf("请求access token失败, %s", err)
	}

	return nil

}

// unit access token 结构
type AccessTokenResponse struct {
	RefreshToken     string `json:"refresh_token" bson:"refresh_token"`
	ExpiresIn        int64  `json:"expires_in" bson:"expires_in"`
	Scope            string `json:"scope" bson:"scope"`
	SessionKey       string `json:"session_key" bson:"session_key"`
	AccessToken      string `json:"access_token" bson:"access_token"`
	SessionSecret    string `json:"session_secret" bson:"session_secret"`
	Error            string `json:"error" bson:"error"`
	ErrorDescription string `json:"error_description" bson:"error_description"`
}
