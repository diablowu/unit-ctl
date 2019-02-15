package main

import (
	"github.com/diablowu/unit-ctl/internal/auth"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

const (
	PROG_NAME = "unit-autj"
	INTRO     = "baidu unit auth utils"
)

func main() {
	mainApp := kingpin.New(PROG_NAME, INTRO)

	authService := auth.AuthService{}

	mainApp.Command("list", "list context").Action(authService.List)

	mainApp.Command("current", "get current context").Action(authService.Current)

	mainApp.Command("set-current", "set current context").Action(authService.SetCurrent).
		Flag("name", "new context name").Required().String()

	{
		cmd := mainApp.Command("add", "add a access token context").Action(authService.Add)
		cmd.Flag("name", "context name").Required().String()
		cmd.Flag("ak", "app key").Required().String()
		cmd.Flag("sk", "secret key").Required().String()
		cmd.Flag("grant", "grant type").Default(auth.DefaultGrantType).String()
	}

	mainApp.Command("delete", "delete context").Action(authService.Delete)

	kingpin.MustParse(mainApp.Parse(os.Args[1:]))

}
