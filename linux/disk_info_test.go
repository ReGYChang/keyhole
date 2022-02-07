package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDiskInfo(t *testing.T) {
	var err error
	var diskInfo *DiskStats

	if diskInfo, err = GetDiskInfo(); err != nil {
		t.Fatal(err)
	}

	data, _ := json.MarshalIndent(diskInfo, "", " ")
	fmt.Print(string(data))
}
