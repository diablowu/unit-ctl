package auth

import (
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"testing"
)

func TestAuthToken(t *testing.T) {

	cred := new(Credential)

	rv := reflect.ValueOf(cred)

	t.Log(rv.IsNil())
	t.Log(rv.Kind())
	t.Log(rv.Elem().CanSet())

	rve := rv.Elem()
	f := rve.FieldByName("AppKey")
	spew.Dump(f)
	t.Log(f.CanSet())

	f.SetString("XXXXXXXXXXXX")

	spew.Dump(cred)

}
