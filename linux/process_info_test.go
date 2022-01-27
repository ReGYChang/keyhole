package linux

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestProcessInfo(t *testing.T) {
	processInfos := GetMongoProcessInfo()

	data, _ := json.MarshalIndent(processInfos, "", " ")
	fmt.Print(string(data))
}
