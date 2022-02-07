package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMemoryInfo(t *testing.T) {
	var memoryStats *MemoryStats
	var err error
	if memoryStats, err = GetMemoryStats(); err != nil {
		t.Fatal(err)
	}

	data, _ := json.MarshalIndent(memoryStats, "", " ")
	fmt.Print(string(data))
}
