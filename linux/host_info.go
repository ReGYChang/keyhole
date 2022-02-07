package linux

import (
	"github.com/shirou/gopsutil/v3/host"
)

type HostInfo struct {
	InfoStat *host.InfoStat   `json:"hostInfoStat"`
	UserStat *[]host.UserStat `json:"hostUserStat"`
	//TemperatureStat *[]host.TemperatureStat `json:"hostTemperatureStat"`
}

func GetHostInfo() (*HostInfo, error) {
	var hostInfo HostInfo
	var err error
	if hostInfo.InfoStat, err = GetHostInfoStat(); err != nil {
		return &hostInfo, err
	}
	if hostInfo.UserStat, err = GetHostUserStat(); err != nil {
		return &hostInfo, err
	}
	//temperatureStat := GetHostTemperatureStat()

	return &hostInfo, err
}

func GetHostInfoStat() (*host.InfoStat, error) {
	infos, err := host.Info()
	return infos, err
}

func GetHostUserStat() (*[]host.UserStat, error) {
	infos, err := host.Users()
	return &infos, err
}

//func GetHostTemperatureStat() *[]host.TemperatureStat {
//	infos, _ := host.SensorsTemperatures()
//	return &infos
//}
