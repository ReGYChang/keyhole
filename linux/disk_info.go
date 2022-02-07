package linux

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type DiskStats struct {
	UsageStat      *disk.UsageStat                 `json:"usageStat"`
	PartitionStat  *[]disk.PartitionStat           `json:"partitionStat"`
	IOCountersStat *map[string]disk.IOCountersStat `json:"ioCountersStat"`
}

func GetDiskInfo() (*DiskStats, error) {
	var diskStats DiskStats
	var err error
	if diskStats.UsageStat, err = GetUsageStat(); err != nil {
		return &diskStats, err
	}
	if diskStats.PartitionStat, err = GetPartitionStat(); err != nil {
		return &diskStats, err
	}
	if diskStats.IOCountersStat, err = GetDiskIOCountersStat(); err != nil {
		return &diskStats, err
	}
	return &diskStats, err
}

func GetUsageStat() (*disk.UsageStat, error) {
	usageStat, err := disk.Usage("/")
	return usageStat, err
}

func GetPartitionStat() (*[]disk.PartitionStat, error) {
	partitionStat, err := disk.Partitions(true)
	return &partitionStat, err
}

func GetDiskIOCountersStat() (*map[string]disk.IOCountersStat, error) {
	ioCountersStat, err := disk.IOCounters()
	return &ioCountersStat, err
}
