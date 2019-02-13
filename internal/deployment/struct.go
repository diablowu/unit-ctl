package deployment

import "github.com/diablowu/unit-ctl/internal"

type GetDmStatusArgs struct {
	BotId *int `json:"botId"`
	DmId  *int `json:"deploymentId"`
}

type ListDmArgs struct {
	BotId    *int  `json:"botId"`
	PageNo   *int  `json:"pageNo"`
	PageSize *int  `json:"pageSize"`
	Watch    *bool `json:"-"`
}

type AddDmArgs struct {
	BotId        *int    `json:"botId"`
	Region       *string `json:"region"`
	ModelVersion *string `json:"modelVersion,omitempty"`
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
