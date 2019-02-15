package model

import "github.com/diablowu/unit-ctl/internal"

type ListModelArgs struct {
	BotId    int  `json:"botId" flag:"bot"`
	PageNo   int  `json:"pageNo" flag:"pn"`
	PageSize int  `json:"pageSize" flag:"ps"`
	Watch    bool `json:"-" flag:"watch"`
}

type ListModelResponse struct {
	internal.ResponseBase
	Result []struct {
		ModelID      string `json:"modelId"`
		CreateTime   string `json:"createTime"`
		ModelVersion string `json:"modelVersion"`
		ModelDesc    string `json:"modelDesc"`
		UpdateTime   string `json:"updateTime"`
		Status       string `json:"status"`
	} `json:"result"`
}
