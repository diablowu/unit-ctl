package deployment

import "github.com/diablowu/unit-ctl/internal"

type GetDmStatusArgs struct {
	BotId int `json:"botId" flag:"bot"`
	DmId  int `json:"deploymentId" flag:"dm"`
}

type ListDmArgs struct {
	BotId    int  `json:"botId" flag:"bot"`
	PageNo   int  `json:"pageNo" flag:"pn"`
	PageSize int  `json:"pageSize" flag:"ps"`
	Watch    bool `json:"-" flag:"watch"`
}

type AddDmArgs struct {
	BotId        int    `json:"botId" flag:"bot"`
	Region       string `json:"region" flag:"region"`
	ModelVersion string `json:"modelVersion,omitempty" flag:"model-version"`
}

type ListDmResponse struct {
	internal.ResponseBase
	Result struct {
		Records []struct {
			CreateTime   string `json:"createTime"`
			ModelVersion string `json:"modelVersion"`
			DeploymentID int    `json:"deploymentId"`
			UpdateTime   string `json:"updateTime"`
			Region       string `json:"region"`
			DeployStatus string `json:"deployStatus"`
		} `json:"records"`
		TotalCount  int `json:"totalCount"`
		CurrentPage int `json:"currentPage"`
	} `json:"result"`
}

type UpdateDmResponse struct {
	internal.ResponseBase
	Result struct {
		DeploymentID string `json:"deploymentId"`
	} `json:"result"`
}

type GetDmStatusResponse struct {
	internal.ResponseBase
	Result struct {
		CreateTime   string `json:"createTime"`
		ModelVersion string `json:"modelVersion"`
		DeploymentID int    `json:"deploymentId"`
		UpdateTime   string `json:"updateTime"`
		Region       string `json:"region"`
		DeployStatus string `json:"deployStatus"`
	} `json:"result"`
}
