package linux

import (
	"github.com/shirou/gopsutil/v3/process"
)

type MongoProcessStats struct {
	//RlimitStat     *[]process.RlimitStat    `json:"rlimitStat"`
	//OpenFileStat   *[]process.OpenFilesStat `json:"openFileStat"`
	MemoryInfoStat *process.MemoryInfoStat `json:"memoryInfoStat"`
	IOCountersStat *process.IOCountersStat `json:"ioCountersStat"`
	PageFaultsStat *process.PageFaultsStat `json:"pageFaultsStat"`
	CPUPercent     float64                 `json:"cpuPercent"`
	MemoryPercent  float32                 `json:"memoryPercent"`
}

func GetMongoProcessInfo() (*MongoProcessStats, error) {
	var p *process.Process
	var err error
	var mongoProcessStats MongoProcessStats

	if p, err = GetMongoProcess(); err != nil {
		return &mongoProcessStats, err
	}
	if mongoProcessStats.MemoryInfoStat, err = GetMemoryInfoStat(p); err != nil {
		return &mongoProcessStats, err
	}
	if mongoProcessStats.IOCountersStat, err = GetIOCounterStat(p); err != nil {
		return &mongoProcessStats, err
	}
	if mongoProcessStats.PageFaultsStat, err = GetPageFaultsStats(p); err != nil {
		return &mongoProcessStats, err
	}
	if mongoProcessStats.CPUPercent, err = GetCPUPercent(p); err != nil {
		return &mongoProcessStats, err
	}
	if mongoProcessStats.MemoryPercent, err = GetMemoryPercent(p); err != nil {
		return &mongoProcessStats, err
	}
	return &mongoProcessStats, nil
}

func GetMongoProcess() (*process.Process, error) {
	infos, err := process.Processes()
	for _, p := range infos {
		if n, _ := p.Name(); n == "mongod" {
			return p, nil
		}
	}
	return nil, err
}

func GetRlimitStat(p *process.Process) (*[]process.RlimitStat, error) {
	infos, err := p.Rlimit()
	return &infos, err
}

func GetOpenFileStat(p *process.Process) (*[]process.OpenFilesStat, error) {
	infos, err := p.OpenFiles()
	return &infos, err
}

func GetMemoryInfoStat(p *process.Process) (*process.MemoryInfoStat, error) {
	infos, err := p.MemoryInfo()
	return infos, err
}

func GetIOCounterStat(p *process.Process) (*process.IOCountersStat, error) {
	infos, err := p.IOCounters()
	return infos, err
}

func GetPageFaultsStats(p *process.Process) (*process.PageFaultsStat, error) {
	infos, err := p.PageFaults()
	return infos, err
}

func GetCPUPercent(p *process.Process) (float64, error) {
	info, err := p.CPUPercent()
	return info, err
}

func GetMemoryPercent(p *process.Process) (float32, error) {
	info, err := p.MemoryPercent()
	return info, err
}
