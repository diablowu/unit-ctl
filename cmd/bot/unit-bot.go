package main

import (
	"github.com/diablowu/unit-ctl/internal/auth"
	"github.com/diablowu/unit-ctl/internal/model"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

const (
	PROG_NAME = "unit-bot"
	INTRO     = "baidu unit bot utility"
)

func main() {
	mainApp := kingpin.New(PROG_NAME, INTRO)

	authCtx := auth.LoadAuthContext(auth.ConfigFilePath)
	modelService := model.NewModelService(authCtx.Token())

	{
		cmd := mainApp.Command("model", "training model")
		cmdList := cmd.Command("list", "list model").Action(modelService.ListModel)
		cmdList.Flag("bot", "bot id").Required().Int()
		cmdList.Flag("pn", "page number").Default("1").Int()
		cmdList.Flag("ps", "page size").Default("15").Int()
		cmdList.Flag("watch", "watch until exit").Default("false").Bool()

	}

	kingpin.MustParse(mainApp.Parse(os.Args[1:]))
}
