package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHostInfo(t *testing.T) {
	var hostInfo *HostInfo
	var err error
	if hostInfo, err = GetHostInfo(); err != nil {
		t.Fatal(err)
	}

	data, _ := json.MarshalIndent(hostInfo, "", " ")
	fmt.Print(string(data))
}
