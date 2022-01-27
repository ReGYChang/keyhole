package linux

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryStats struct {
	VirtualMemoryStat *mem.VirtualMemoryStat `json:"virtualMemoryStat"`
	SwapMemoryStat    *mem.SwapMemoryStat    `json:"swapMemoryStat"`
}

func GetMemoryStats() MemoryStats {
	virtualMemoryStat := GetVirtualMemoryStat()
	swapMemoryStat := GetSwapMemoryStat()

	return MemoryStats{
		VirtualMemoryStat: virtualMemoryStat,
		SwapMemoryStat:    swapMemoryStat,
	}
}

func GetVirtualMemoryStat() *mem.VirtualMemoryStat {
	infos, _ := mem.VirtualMemory()
	return infos
}

func GetSwapMemoryStat() *mem.SwapMemoryStat {
	infos, _ := mem.SwapMemory()
	return infos
}
