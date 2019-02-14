package auth

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"path"
	"reflect"
)

var ConfigFilePath string

func init() {
	if cacheHome, err := os.UserCacheDir(); err == nil {
		configDir := path.Join(cacheHome, "unit-ctl")
		if _, err := os.Stat(configDir); err != nil && os.IsNotExist(err) {
			if err := os.Mkdir(configDir, 0700); err != nil {
				log.Fatalf("没办法初始化配置目录, %s", err)
			}
		}
		ConfigFilePath = path.Join(configDir, "config.json")
	} else {
		log.Fatalf("没找到UserCacheDir %s", err)
	}

}

type AuthService struct {
}

func (as AuthService) Add(ctx *kingpin.ParseContext) error {

	cred := new(Credential)

	setCredField(ctx.SelectedCommand.Model().Flags, cred)

	token := cred.Auth()
	spew.Dump(token)

	ac := AuthContext{}
	ac.Cred = make(map[string]AccessTokenResponse, 1)
	ac.Cred[cred.Name] = *token
	ac.Current = cred.Name
	ac.Save(ConfigFilePath)
	return nil
}

func setCredField(flags []*kingpin.FlagModel, target *Credential) {
	flagMap := make(map[string]*kingpin.FlagModel, 0)
	for _, flagValue := range flags {
		flagMap[flagValue.Name] = flagValue
	}

	tv := reflect.ValueOf(target).Elem()
	tvt := tv.Type()
	fn := tv.NumField()
	for i := 0; i < fn; i++ {
		ft := tv.Field(i)
		ft.SetString(flagMap[tvt.Field(i).Tag.Get("fn")].Value.String())
	}

}
func (as AuthService) List(ctx *kingpin.ParseContext) error {
	return nil
}

func (as AuthService) Current(ctx *kingpin.ParseContext) error {
	return nil
}

func (as AuthService) SetCurrent(ctx *kingpin.ParseContext) error {
	return nil
}

func (as AuthService) Delete(ctx *kingpin.ParseContext) error {
	return errors.New("Not Supported!!!")
}
