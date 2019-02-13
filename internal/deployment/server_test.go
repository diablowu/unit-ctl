package deployment

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDeployService_Add(t *testing.T) {
	a := AddDmArgs{}
	ss := 2434
	bb := "bj"
	a.BotId = &ss
	a.Region = &bb

	bs, _ := json.Marshal(a)

	fmt.Println(string(bs))

}
