package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestProcessInfo(t *testing.T) {
	processInfos, err := GetMongoProcessInfo()
	if err != nil {
		t.Fatal(err)
	}

	data, _ := json.MarshalIndent(processInfos, "", " ")
	fmt.Print(string(data))
}

func TestProcessCpuUsage(t *testing.T) {
	p, err := GetMongoProcess()
	if err != nil {
		t.Fatal(err)
	}

	info, err := GetCPUPercent(p)
	fmt.Sprintf("CPU Usage:%d\n", info)
}

func TestProcessMemoryUsage(t *testing.T) {
	p, err := GetMongoProcess()
	if err != nil {
		t.Fatal(err)
	}

	info, err := GetMemoryPercent(p)
	fmt.Sprintf("Memory Usage:%d\n", info)
}
