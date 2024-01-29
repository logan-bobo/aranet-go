package main

import (
	"fmt"
	"os"
	"strings"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	fmt.Println("Hello Aranet 4")
	deviceMac := os.Getenv("ARANET_MAC")
	fmt.Println(deviceMac)

	// Enable BLE interface.
	adapter.Enable()

	// Start scanning.
	println("scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		if strings.Contains(strings.ToLower(device.LocalName()), "aranet") {
			fmt.Println(device.LocalName(), device.Address.String())
		} 
	})

	fmt.Println(err)

	// TODO: split scanning out to seperate function that takes in an adapter and return a struct of device info
}
