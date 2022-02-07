package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCpuInfo(t *testing.T) {
	var cpuStats *CPUStats
	var err error

	if cpuStats, err = GetCPUStats(); err != nil {
		t.Fatal(err)
	}

	data, _ := json.MarshalIndent(cpuStats, "", " ")
	fmt.Print(string(data))
}
