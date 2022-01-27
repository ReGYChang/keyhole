package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDiskInfo(t *testing.T) {
	diskInfo := GetDiskInfo()

	data, _ := json.MarshalIndent(diskInfo, "", " ")
	fmt.Print(string(data))
}
