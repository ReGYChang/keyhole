package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCpuInfo(t *testing.T) {
	cpuStats := GetCPUStats()

	data, _ := json.MarshalIndent(cpuStats, "", " ")
	fmt.Print(string(data))
}
