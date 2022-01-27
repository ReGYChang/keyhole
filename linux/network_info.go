package linux

import (
	"github.com/shirou/gopsutil/v3/net"
)

type NetInfo struct {
	IOCountersStat *net.IOCountersStat `json:"ioCountersStat"`
	ConntrackStat  *net.ConntrackStat  `json:"conntrackStat"`
	//ConnectionStat *net.ConnectionStat  `json:"connectionStat"`
}

func GetNetInfo() NetInfo {
	ioCountersStat := GetIOCountersStat()
	//conntrackStat := GetConntrackStat()

	return NetInfo{
		IOCountersStat: ioCountersStat,
		//ConntrackStat:  conntrackStat,
	}
}

func GetIOCountersStat() *net.IOCountersStat {
	infos, _ := net.IOCounters(false)
	return &infos[0]
}

func GetConntrackStat() *net.ConntrackStat {
	infos, _ := net.ConntrackStats(false)
	return &infos[0]
}

//func GetConnectionStat() *[]net.ConnectionStat {
//	infos, _ := net.Connections()
//	return &infos
//}
