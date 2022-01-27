package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMemoryInfo(t *testing.T) {
	memoryStats := GetMemoryStats()

	data, _ := json.MarshalIndent(memoryStats, "", " ")
	fmt.Print(string(data))
}
