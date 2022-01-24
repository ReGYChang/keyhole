package linux

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	sysinf, err := GetSysInfo()
	// vmstatdata, err := GetVMStat()
	networkCountersdata, err := GetNetworkCounters()
	// loadAveragedata, err := GetLoadAverage()

	sysInfo, err := json.MarshalIndent(&sysinf, "", "  ")
	networkCountersInfo, err := json.MarshalIndent(&networkCountersdata, "", "  ")
	// loadAverageInfo, err := json.MarshalIndent(&loadAveragedata, "", "  ")

	// vmstatInfo_data, err := json.MarshalIndent(&vmstatInfo, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(sysInfo))
	fmt.Println(string(networkCountersInfo))
	fmt.Println(sysInfo)
	fmt.Println(networkCountersInfo)
	fmt.Println(sysinf)
	fmt.Printf("Datatype of i : %T\n", &sysinf)

	// fmt.Println(string(loadAverageInfo))
	// fmt.Println(string(vmstatInfo_data))
}
