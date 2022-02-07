package linux

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryStats struct {
	VirtualMemoryStat *mem.VirtualMemoryStat `json:"virtualMemoryStat"`
	SwapMemoryStat    *mem.SwapMemoryStat    `json:"swapMemoryStat"`
}

func GetMemoryStats() (*MemoryStats, error) {
	var memoryStats MemoryStats
	var err error

	if memoryStats.VirtualMemoryStat, err = GetVirtualMemoryStat(); err != nil {
		return &memoryStats, err
	}
	if memoryStats.SwapMemoryStat, err = GetSwapMemoryStat(); err != nil {
		return &memoryStats, err
	}

	return &memoryStats, nil
}

func GetVirtualMemoryStat() (*mem.VirtualMemoryStat, error) {
	infos, err := mem.VirtualMemory()
	return infos, err
}

func GetSwapMemoryStat() (*mem.SwapMemoryStat, error) {
	infos, err := mem.SwapMemory()
	return infos, err
}
