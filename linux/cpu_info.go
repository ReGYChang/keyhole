package linux

import (
	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUStats struct {
	CPUInfo  *cpu.InfoStat  `json:"processInfo"`
	CPUTimes *cpu.TimesStat `json:"cpuTimes"`
}

func GetCPUStats() (*CPUStats, error) {
	var cpuStats CPUStats
	if cpuInfo, err := GetCPUInfo(); err != nil {
		return nil, err
	} else {
		cpuStats.CPUInfo = cpuInfo
	}

	if cpuTimesStat, err := GetCPUTime(); err != nil {
		return nil, err
	} else {
		cpuStats.CPUTimes = cpuTimesStat
	}

	return &cpuStats, nil
}

func GetCPUInfo() (*cpu.InfoStat, error) {
	infos, err := cpu.Info()
	return &infos[0], err
}

func GetCPUTime() (*cpu.TimesStat, error) {
	infos, err := cpu.Times(false)
	return &infos[0], err
}
