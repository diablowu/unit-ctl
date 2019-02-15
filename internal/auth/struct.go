package auth

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"io/ioutil"
	"log"
)

const (
	DefaultGrantType = "client_credentials"
)

type AuthContext struct {
	Current string                         `json:"current"`
	Cred    map[string]AccessTokenResponse `json:"access_tokens"`
}

func (ac *AuthContext) Token() string {
	return ac.Cred[ac.Current].AccessToken
}

func (ac *AuthContext) Add(name string, token AccessTokenResponse) {

}
func (ac *AuthContext) Save(path string) {
	log.Println("Save config to " + path)
	if bs, err := json.Marshal(*ac); err == nil {
		if err := ioutil.WriteFile(path, bs, 0600); err != nil {
			log.Fatalf("不能保存auth文件, %s", err)
		}

	} else {
		log.Fatalf("不能序列, %s", err)
	}
}

func LoadAuthContext(path string) (*AuthContext) {

	if bs, err := ioutil.ReadFile(path); err == nil {
		authCtx := new(AuthContext)
		if err := json.Unmarshal(bs, authCtx); err == nil {
			return authCtx
		} else {
			log.Fatalln("不能加载auth context")
		}
	}

	return nil
}

type Credential struct {
	Name      string `flag:"name"`
	AppKey    string `flag:"ak"`
	SecretKey string `flag:"sk"`
	GrantType string `flag:"grant"`
}

// 生成access_token获取的query string
func (cred Credential) Auth() *AccessTokenResponse {

	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=%s&client_id=%s&client_secret=%s", cred.GrantType, cred.AppKey, cred.SecretKey)
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
