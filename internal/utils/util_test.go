package utils

import (
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/alecthomas/kingpin.v2"
	"testing"
)

type Flags struct {
	Name      string  `flag:"name"`
	Age       int     `flag:"age"`
	Score     float64 `flag:"score"`
	Watch     bool    `flag:"watch"`
	ExtraInfo string
}

func TestExtractFlag(t *testing.T) {

	app := kingpin.New("test-app", "test-app")
	cmd := app.Command("test", "test func")
	cmd.Flag("name", "name").Required().String()
	cmd.Flag("age", "age").Required().Int()
	cmd.Flag("score", "score").Required().Float64()
	cmd.Flag("watch", "watch").Default("false").Bool()

	cmd.Action(func(ctx *kingpin.ParseContext) error {

		v := new(Flags)
		ExtractFlag(ctx.SelectedCommand.Model().Flags, v)
		spew.Dump(v)
		return nil
	})

	args := []string{"test", "--name", "Name", "--age=31", "--score", "341.223","--watch"}
	if _, err := app.Parse(args); err != nil {
		t.Log(err)
	}

}
