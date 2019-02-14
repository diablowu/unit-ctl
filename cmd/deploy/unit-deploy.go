package main

import (
	"github.com/diablowu/unit-ctl/internal/auth"
	"github.com/diablowu/unit-ctl/internal/deployment"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
)

const (
	PROG_NAME = "unit-deploy"
	INTRO     = "baidu unit model deployer"
)

func main() {
	mainApp := kingpin.New(PROG_NAME, INTRO)

	cmdStatus := mainApp.Command("status", "get statue")
	cmdStatusArgs := deployment.GetDmStatusArgs{}
	cmdStatusArgs.BotId = cmdStatus.Flag("bot", "bot id").Required().Int()
	cmdStatusArgs.DmId = cmdStatus.Flag("dm", "deployment id").Required().Int()

	cmdList := mainApp.Command("list", "list deployment")
	cmdListArgs := deployment.ListDmArgs{}
	cmdListArgs.Watch = cmdList.Flag("watch", "watching").Short('w').Default("false").Bool()
	cmdListArgs.BotId = cmdList.Flag("bot", "bot id").Required().Int()
	cmdListArgs.PageNo = cmdList.Flag("pn", "page number").Default("1").Int()
	cmdListArgs.PageSize = cmdList.Flag("ps", "page size").Default("15").Int()

	cmdAdd := mainApp.Command("add", "add deployment")
	cmdAddArgs := deployment.AddDmArgs{}
	cmdAddArgs.BotId = cmdAdd.Flag("bot", "bot id").Required().Int()
	cmdAddArgs.Region = cmdAdd.Flag("region", "bce region").Default("bj").Enum("bj", "su", "gz")
	cmdAddArgs.ModelVersion = cmdAdd.Flag("model-version", "model version in sandbox").String()

	cmdUpdate := mainApp.Command("update", "update deployment")
	cmdUpdateArgs := deployment.AddDmArgs{}
	cmdUpdateArgs.BotId = cmdUpdate.Flag("bot", "bot id").Required().Int()
	cmdUpdateArgs.Region = cmdUpdate.Flag("region", "bce region").Default("bj").Enum("bj", "su", "gz")
	cmdUpdateArgs.ModelVersion = cmdUpdate.Flag("model-version", "model version in sandbox").String()

	cmd := kingpin.MustParse(mainApp.Parse(os.Args[1:]))

	authCtx := auth.LoadAuthContext(auth.ConfigFilePath)

	deployService := deployment.NewDeployService(authCtx.Token())
	switch cmd {
	case cmdStatus.FullCommand():
		{
			deployService.Status(cmdStatusArgs)
		}
	case cmdList.FullCommand():
		{
			if *cmdListArgs.Watch {
				deployService.ListWatch(cmdListArgs)
			} else {
				if _, err := deployService.List(cmdListArgs); err != nil {
					log.Fatal(err)
				}
			}

		}
	case cmdAdd.FullCommand():
		{

		}
	case cmdUpdate.FullCommand():
		{
			if _, err := deployService.Update(cmdUpdateArgs); err != nil {
				log.Fatal(err)
			}

		}
	default:
		log.Fatalln("no such cmd")
	}
}
