package main

import (
	"github.com/diablowu/unit-ctl/internal/auth"
	"github.com/diablowu/unit-ctl/internal/deployment"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

const (
	PROG_NAME = "unit-deploy"
	INTRO     = "baidu unit model deployer"
)

func main() {
	mainApp := kingpin.New(PROG_NAME, INTRO)

	authCtx := auth.LoadAuthContext(auth.ConfigFilePath)
	deployService := deployment.NewDeployService(authCtx.Token())

	cmdStatus := mainApp.Command("status", "get statue").Action(deployService.StatusAction)
	cmdStatus.Flag("bot", "bot id").Required().Int()
	cmdStatus.Flag("dm", "deployment id").Required().Int()

	cmdList := mainApp.Command("list", "list deployment").Action(deployService.ListAction)
	cmdList.Flag("watch", "watching").Short('w').Default("false").Bool()
	cmdList.Flag("bot", "bot id").Required().Int()
	cmdList.Flag("pn", "page number").Default("1").Int()
	cmdList.Flag("ps", "page size").Default("15").Int()

	cmdAdd := mainApp.Command("add", "add deployment").Action(deployService.AddAction)
	cmdAdd.Flag("bot", "bot id").Required().Int()
	cmdAdd.Flag("region", "bce region").Default("bj").Enum("bj", "su", "gz")
	cmdAdd.Flag("model-version", "model version in sandbox").String()

	cmdUpdate := mainApp.Command("update", "update deployment").Action(deployService.UpdateAction)
	cmdUpdate.Flag("bot", "bot id").Required().Int()
	cmdUpdate.Flag("region", "bce region").Default("bj").Enum("bj", "su", "gz")
	cmdUpdate.Flag("model-version", "model version in sandbox").String()

	kingpin.MustParse(mainApp.Parse(os.Args[1:]))
}
