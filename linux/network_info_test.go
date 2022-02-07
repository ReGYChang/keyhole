package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNetInfo(t *testing.T) {
	var err error
	var netInfos *NetInfo
	if netInfos, err = GetNetInfo(); err != nil {
		t.Fatal(err)
	}

	data, _ := json.MarshalIndent(netInfos, "", " ")
	fmt.Print(string(data))
}
