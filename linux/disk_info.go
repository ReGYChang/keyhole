package linux

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type DiskInfo struct {
	UsageStat      *disk.UsageStat                 `json:"usageStat"`
	PartitionStat  *[]disk.PartitionStat           `json:"partitionStat"`
	IOCountersStat *map[string]disk.IOCountersStat `json:"ioCountersStat"`
}

func GetDiskInfo() DiskInfo {
	usageStat := GetUsageStat()
	partitionStat := GetPartitionStat()
	ioCountersStat := GetDiskIOCountersStat()
	return DiskInfo{
		UsageStat:      usageStat,
		PartitionStat:  partitionStat,
		IOCountersStat: ioCountersStat,
	}
}

func GetUsageStat() *disk.UsageStat {
	usageStat, _ := disk.Usage("/")
	return usageStat
}

func GetPartitionStat() *[]disk.PartitionStat {
	partitionStat, _ := disk.Partitions(true)
	return &partitionStat
}

func GetDiskIOCountersStat() *map[string]disk.IOCountersStat {
	ioCountersStat, _ := disk.IOCounters()
	return &ioCountersStat
}
