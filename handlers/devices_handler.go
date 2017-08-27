package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/patternMiner/bazelsauce/context"
)

func DevicesHandler (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	setAccessControlResponseHeaders(w, req)
	devices := []Device{}
	for device := range context.DeviceList {
		deviceRecord := context.DeviceMap[device]
		devices = append(devices, Device{Id: deviceRecord[0], Description: deviceRecord[1]})
	}
	data, err := json.Marshal(Response{Items: devices})
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(data))
}
