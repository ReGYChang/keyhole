package linux

import (
	"github.com/shirou/gopsutil/v3/net"
)

type NetInfo struct {
	IOCountersStat *net.IOCountersStat `json:"ioCountersStat"`
	ConntrackStat  *net.ConntrackStat  `json:"conntrackStat"`
	//ConnectionStat *net.ConnectionStat  `json:"connectionStat"`
}

func GetNetInfo() (*NetInfo, error) {
	var netInfo NetInfo
	var err error
	if netInfo.IOCountersStat, err = GetIOCountersStat(); err != nil {
		return &netInfo, err
	}
	//if netInfo.ConntrackStat, err = GetConntrackStat(); err != nil {
	//	return &netInfo, err
	//}

	return &netInfo, nil
}

func GetIOCountersStat() (*net.IOCountersStat, error) {
	infos, err := net.IOCounters(false)
	return &infos[0], err
}

func GetConntrackStat() (*net.ConntrackStat, error) {
	infos, err := net.ConntrackStats(false)
	return &infos[0], err
}

//func GetConnectionStat() *[]net.ConnectionStat {
//	infos, _ := net.Connections()
//	return &infos
//}
