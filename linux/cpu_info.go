package linux

import (
	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUStats struct {
	CPUInfo  *cpu.InfoStat  `json:"processInfo"`
	CPUTimes *cpu.TimesStat `json:"cpuTimes"`
}

func GetCPUStats() CPUStats {
	cpuInfo := GetCPUInfo()
	cpuTimesStat := GetCPUTime()

	return CPUStats{
		CPUInfo:  cpuInfo,
		CPUTimes: cpuTimesStat,
	}
}

func GetCPUInfo() *cpu.InfoStat {
	infos, _ := cpu.Info()
	return &infos[0]
}

func GetCPUTime() *cpu.TimesStat {
	infos, _ := cpu.Times(false)
	return &infos[0]
}
