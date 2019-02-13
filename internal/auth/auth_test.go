package auth

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestAuthToken(t *testing.T) {

	cred := NewCredential("7AViOnGwX6ixz3OB2h2vkSDQ", "xK7i76QGds6Ne51iof1SO5MTj00yETDO")

	token := cred.Auth()
	spew.Dump(token)
}
