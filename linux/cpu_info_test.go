package linux

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"testing"
	"time"
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

func TestCpuCount(t *testing.T) {
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)

	fmt.Printf("physical count:%d logical count:%d\n", physicalCnt, logicalCnt)
}

func TestCpuUsage(t *testing.T) {
	totalPercent, _ := cpu.Percent(time.Second, false)
	perPercents, _ := cpu.Percent(time.Second, true)

	fmt.Printf("total percent:%v per percents:%v", totalPercent, perPercents)
}

func TestCpuLoad(t *testing.T) {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}
