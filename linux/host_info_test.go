package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHostInfo(t *testing.T) {
	hostInfo := GetHostInfo()

	data, _ := json.MarshalIndent(hostInfo, "", " ")
	fmt.Print(string(data))
}
