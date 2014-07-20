package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/modouwifi/md/api"
)

var cmdDevices = cli.Command{
	Name:  "devices",
	Usage: "Show Devices Info",
	Subcommands: []cli.Command{
		{
			Name:   "cables",
			Usage:  "Show cables info",
			Action: runDevicesCables,
		},
		{
			Name:   "wifis",
			Usage:  "Show wifis info",
			Action: runDevicesWifis,
		},
		{
			Name:   "disk",
			Usage:  "Show disk info",
			Action: runDevicesDisk,
		},
	},
}

func runDevicesCables(c *cli.Context) {
	req, err := client.NewRequest("GET", "/devices/cables")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.DevicesCables
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runDevicesWifis(c *cli.Context) {
	req, err := client.NewRequest("GET", "/devices/wifis")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.DevicesWifis
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runDevicesDisk(c *cli.Context) {
	req, err := client.NewRequest("GET", "/devices/disk")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.DeviceDisk
	err = req.ToJSON(&result)
	fmt.Println(result)
}
