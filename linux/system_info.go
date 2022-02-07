package linux

import (
	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
	"github.com/pkg/errors"
)

type SysInfo struct {
	MongoProcessStats *MongoProcessStats    `json:"processInfo"`
	CPUStats          *CPUStats             `json:"CPUStats"`
	MemoryStats       *MemoryStats          `json:"memoryStats"`
	DiskStats         *DiskStats            `json:"diskStats"`
	HostInfo          *HostInfo             `json:"hostInfo"`
	HostMemoryInfo    *types.HostMemoryInfo `json:"hostMemoryInfo"`
	VmStatInfo        *types.VMStatInfo     `json:"vmStatInfo"`
	NetworkStats      *NetInfo              `json:"networkStats"`
}

// GetVMStat gets linux vmstat metrics
func GetVMStat() (*types.VMStatInfo, error) {
	// TODO: We may want to pull this code out of go-sysinfo.
	// It's platform specific, and not used by anything else.
	h, err := sysinfo.Host()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read self host information")
	}
	vmstatHandle, ok := h.(types.VMStat)
	if !ok {
		return nil, errors.New("VMStat not available for platform")
	}
	info, err := vmstatHandle.VMStat()
	if err != nil {
		return nil, errors.Wrap(err, "error getting VMStat info")
	}
	return info, nil

}

//GetLoadAverage gets linux vmstat metrics
func GetLoadAverage() (types.LoadAverageInfo, error) {
	// TODO: We may want to pull this code out of go-sysinfo.
	// It's platform specific, and not used by anything else.
	h, _ := sysinfo.Host()
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to read self host information")
	// }
	loadAverageHandle, _ := h.(types.LoadAverage)
	// if !ok {
	// 	return nil, errors.New("LoadAverage not available for platform")
	// }
	info := loadAverageHandle.LoadAverage()

	return info, nil

}

// GetNetworkCounters GetVMStat gets linux vmstat metrics
func GetNetworkCounters() (*types.NetworkCountersInfo, error) {
	// TODO: We may want to pull this code out of go-sysinfo.
	// It's platform specific, and not used by anything else.
	h, err := sysinfo.Host()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read self host information")
	}
	networkCountersHandle, ok := h.(types.NetworkCounters)
	if !ok {
		return nil, errors.New("NetworkCounters not available for platform")
	}
	info, err := networkCountersHandle.NetworkCounters()
	if err != nil {
		return nil, errors.Wrap(err, "error getting NetworkCounters info")
	}
	return info, nil

}

func GetSysInfo() (SysInfo, error) {
	var sysInfo SysInfo
	//var host types.Host
	var err error

	if sysInfo.CPUStats, err = GetCPUStats(); err != nil {
		return sysInfo, err
	}
	if sysInfo.MemoryStats, err = GetMemoryStats(); err != nil {
		return sysInfo, err
	}
	if sysInfo.DiskStats, err = GetDiskInfo(); err != nil {
		return sysInfo, err
	}
	if sysInfo.MongoProcessStats, err = GetMongoProcessInfo(); err != nil {
		return sysInfo, err
	}
	if sysInfo.HostInfo, err = GetHostInfo(); err != nil {
		return sysInfo, err
	}
	//if sysInfo.HostMemoryInfo, err = host.Memory(); err != nil {
	//	return sysInfo, err
	//}
	if sysInfo.VmStatInfo, err = GetVMStat(); err != nil {
		return sysInfo, err
	}
	if sysInfo.NetworkStats, err = GetNetInfo(); err != nil {
		return sysInfo, err
	}

	return sysInfo, nil
}
