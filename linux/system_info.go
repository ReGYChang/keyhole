package linux

import (
	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
	"github.com/pkg/errors"
)

type SysInfo struct {
	ProcessInfo    types.ProcessInfo     `json:"process_info"`
	CPUTimes       types.CPUTimes        `json:"cpu_times"`
	MemoryInfo     types.MemoryInfo      `json:"memory_info"`
	UserInfo       types.UserInfo        `json:"user_info"`
	HostInfo       types.HostInfo        `json:"host_info"`
	HostMemoryInfo *types.HostMemoryInfo `json:"host_memory_info"`
	VmStatInfo     *types.VMStatInfo     `json:"vm_stat_info"`
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
	var process types.Process
	var host types.Host
	var err error
	if process, err = sysinfo.Self(); err != nil {
		return sysInfo, err
	}

	if sysInfo.ProcessInfo, err = process.Info(); err != nil {
		return sysInfo, err
	}
	if sysInfo.CPUTimes, err = process.CPUTime(); err != nil {
		return sysInfo, err
	}
	if sysInfo.MemoryInfo, err = process.Memory(); err != nil {
		return sysInfo, err
	}
	if sysInfo.UserInfo, err = process.User(); err != nil {
		return sysInfo, err
	}

	if host, err = sysinfo.Host(); err != nil {
		return sysInfo, err
	}
	sysInfo.HostInfo = host.Info()
	if sysInfo.HostMemoryInfo, err = host.Memory(); err != nil {
		return sysInfo, err
	}

	if sysInfo.VmStatInfo, err = GetVMStat(); err != nil {
		return sysInfo, err
	}

	return sysInfo, nil
}
