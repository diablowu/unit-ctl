package deployment

import (
	"errors"
	"fmt"
	"github.com/diablowu/unit-ctl/internal"
	"github.com/levigross/grequests"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type DeployService struct {
	internal.ServiceBaseInfo
}

func NewDeployService(accessToken string) (DeployService) {
	ds := DeployService{}
	ds.AccessToken = accessToken
	ds.Endpoint = internal.ManageEndpoint
	return ds
}

func (ds DeployService) Status(args GetDmStatusArgs) {
	method := "deployment/getStatus"
	reqOpts := new(grequests.RequestOptions)
	reqOpts.JSON = args
	if rsp, err := grequests.Post(ds.GetFullUrl(method), reqOpts); err == nil && rsp.Ok {
		status := GetDmStatusResponse{}
		if err := rsp.JSON(&status); err == nil {
			if status.Success() {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"DeploymentID", "Region", "ModelVersion", "Created", "Status"})
				table.Append([]string{strconv.Itoa(status.Result.DeploymentID), status.Result.Region, status.Result.ModelVersion, status.Result.CreateTime, status.Result.DeployStatus})
				table.Render()
			} else {
				log.Fatalln(status.ErrorReason())
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func (ds DeployService) Add(args AddDmArgs) {
}

func (ds DeployService) Update(args AddDmArgs) (*UpdateDmResponse, error) {

	if *args.ModelVersion == "" {
		args.ModelVersion = nil
	}

	method := "deployment/updateModelVersion"
	reqOpts := new(grequests.RequestOptions)
	reqOpts.JSON = args
	if rsp, err := grequests.Post(ds.GetFullUrl(method), reqOpts); err == nil && rsp.Ok {
		dmUpdate := UpdateDmResponse{}
		if err := rsp.JSON(&dmUpdate); err == nil {
			if dmUpdate.Success() {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"BotID", "DeploymentID", "Region", "ModelVersion"})
				mv := "latest"
				if args.ModelVersion != nil {
					mv = *args.ModelVersion
				}

				table.Append([]string{strconv.Itoa(*args.BotId), dmUpdate.Result.DeploymentID, *args.Region, mv})
				table.Render()
				return &dmUpdate, nil
			} else {
				return nil, errors.New(dmUpdate.ErrorMsg)
			}

		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (ds DeployService) ListWatch(args ListDmArgs) {
	ds.List(args)
	tm := time.NewTicker(time.Second * 5)
	go func() {
		for t := range tm.C {
			fmt.Println("Watching at ", t)
			ds.List(args)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
		{
			os.Exit(0)
		}
	}

}
//func (ds DeployService) ListDm(ctx *kingpin.ParseContext) error {
//	ctx.SelectedCommand.Model().Flags[0]
//}

func (ds DeployService) List(args ListDmArgs) (*ListDmResponse, error) {
	method := "deployment/list"
	reqOpts := new(grequests.RequestOptions)
	reqOpts.JSON = args
	if rsp, err := grequests.Post(ds.GetFullUrl(method), reqOpts); err == nil && rsp.Ok {
		dmList := ListDmResponse{}
		if err := rsp.JSON(&dmList); err == nil {

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"DeploymentID", "Region", "ModelVersion", "Created", "Status"})
			for _, dm := range dmList.Result.Records {
				table.Append([]string{strconv.Itoa(dm.DeploymentID), dm.Region, dm.ModelVersion, dm.CreateTime, dm.DeployStatus})
			}
			table.Render()

			return &dmList, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

}
