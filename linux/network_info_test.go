package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNetInfo(t *testing.T) {
	netInfos := GetNetInfo()

	data, _ := json.MarshalIndent(netInfos, "", " ")
	fmt.Print(string(data))
}
