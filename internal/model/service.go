package model

import (
	"github.com/diablowu/unit-ctl/internal"
	"github.com/diablowu/unit-ctl/internal/utils"
	"github.com/levigross/grequests"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

type ModelService struct {
	internal.ServiceBaseInfo
}

func NewModelService(accessToken string) (ModelService) {
	ms := ModelService{}
	ms.AccessToken = accessToken
	ms.Endpoint = internal.ManageEndpoint
	return ms
}

func (ms ModelService) ListModel(ctx *kingpin.ParseContext) (error) {
	method := "model/list"

	args := new(ListModelArgs)
	utils.ExtractFlag(ctx.SelectedCommand.Model().Flags, args)
	opts := new(grequests.RequestOptions)
	opts.JSON = args

	if rsp, err := grequests.Post(ms.GetFullUrl(method), opts); err == nil && rsp.Ok {
		list := ListModelResponse{}
		if err := rsp.JSON(&list); err == nil {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ModelID", "ModelVersion", "Desc", "Created", "Status"})
			for _, dm := range list.Result {
				table.Append([]string{dm.ModelID, dm.ModelVersion, dm.ModelDesc, dm.CreateTime, dm.Status})
			}
			table.Render()
		}
	}
	return nil
}
