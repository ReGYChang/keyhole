package linux

import (
	"github.com/shirou/gopsutil/v3/process"
)

type MongoProcessInfo struct {
	//RlimitStat     *[]process.RlimitStat    `json:"rlimitStat"`
	//OpenFileStat   *[]process.OpenFilesStat `json:"openFileStat"`
	MemoryInfoStat *process.MemoryInfoStat `json:"memoryInfoStat"`
	IOCountersStat *process.IOCountersStat `json:"ioCountersStat"`
	PageFaultsStat *process.PageFaultsStat `json:"pageFaultsStat"`
}

func GetMongoProcessInfo() MongoProcessInfo {
	p := GetMongoProcess()
	return MongoProcessInfo{
		//RlimitStat:     GetRlimitStat(p),
		//OpenFileStat:   GetOpenFileStat(p),
		MemoryInfoStat: GetMemoryInfoStat(p),
		IOCountersStat: GetIOCounterStat(p),
		PageFaultsStat: GetPageFaultsStats(p),
	}
}

func GetMongoProcess() *process.Process {
	infos, _ := process.Processes()
	for _, p := range infos {
		if n, _ := p.Name(); n == "mongod" {
			return p
		}
	}
	return nil
}

func GetRlimitStat(p *process.Process) *[]process.RlimitStat {
	infos, _ := p.Rlimit()
	return &infos
}

func GetOpenFileStat(p *process.Process) *[]process.OpenFilesStat {
	infos, _ := p.OpenFiles()
	return &infos
}

func GetMemoryInfoStat(p *process.Process) *process.MemoryInfoStat {
	infos, _ := p.MemoryInfo()
	return infos
}

func GetIOCounterStat(p *process.Process) *process.IOCountersStat {
	infos, _ := p.IOCounters()
	return infos
}

func GetPageFaultsStats(p *process.Process) *process.PageFaultsStat {
	infos, _ := p.PageFaults()
	return infos
}
