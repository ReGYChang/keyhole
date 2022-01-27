package linux

import (
	"github.com/shirou/gopsutil/v3/host"
)

type HostInfo struct {
	InfoStat *host.InfoStat   `json:"hostInfoStat"`
	UserStat *[]host.UserStat `json:"hostUserStat"`
	//TemperatureStat *[]host.TemperatureStat `json:"hostTemperatureStat"`
}

func GetHostInfo() HostInfo {
	infoStat := GetHostInfoStat()
	userStat := GetHostUserStat()
	//temperatureStat := GetHostTemperatureStat()

	return HostInfo{
		InfoStat: infoStat,
		UserStat: userStat,
		//TemperatureStat: temperatureStat,
	}
}

func GetHostInfoStat() *host.InfoStat {
	infos, _ := host.Info()
	return infos
}

func GetHostUserStat() *[]host.UserStat {
	infos, _ := host.Users()
	return &infos
}

//func GetHostTemperatureStat() *[]host.TemperatureStat {
//	infos, _ := host.SensorsTemperatures()
//	return &infos
//}
