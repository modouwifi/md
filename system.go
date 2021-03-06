package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/modouwifi/md/api"
)

var cmdSystem = cli.Command{
	Name:  "system",
	Usage: "Operate system",
	//Action: runSystem,
	Subcommands: []cli.Command{
		{
			Name:   "current",
			Usage:  "Display currently version",
			Action: runSystemGetVersionInfo,
		},
		{
			Name:   "check",
			Usage:  "Check latest version",
			Action: runSystemCheckLatestVersion,
		},
		{
			Name:   "reboot",
			Usage:  "Reboot router",
			Action: runSystemReboot,
		},
		{
			Name:   "safe",
			Usage:  "Enter safe mode",
			Action: runSystemSafe,
		},
		{
			Name:   "reset",
			Usage:  "Reset system",
			Action: runSystemReset,
		},
		{
			Name:   "backlight",
			Usage:  "Control the backlight",
			Action: runSystemBacklight,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "action, a", Value: "lock", Usage: "Control the blacklight, lock|release|wakeup"},
			},
		},
		{
			Name:   "time",
			Usage:  "Get/set system time",
			Action: runSystemTime,
		},
		{
			Name:   "cable",
			Usage:  "Get system cable connections",
			Action: runSystemCable,
		},
	},
}

func runSystemGetVersionInfo(c *cli.Context) {
	req, err := client.NewRequest("GET", "/system/get_version_info")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.SystemVersion
	err = req.ToJSON(&result)
	fmt.Println(result.Code)
	fmt.Println(result.Track)
	fmt.Println(result.Version1)
	fmt.Println(result.Version2)
}

func runSystemCheckLatestVersion(c *cli.Context) {
	req, err := client.NewRequest("GET", "/system/check_remote_version_upgrade")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.SystemLatestVersion
	err = req.ToJSON(&result)
	fmt.Println(result)
	fmt.Println(result.Filename)
	fmt.Println(result.Version)
	fmt.Println(result.Releasenote)
}

func runSystemReboot(c *cli.Context) {
	req, err := client.NewRequest("GET", "/system/reboot")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.ResponseMessage
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runSystemSafe(c *cli.Context) {
	req, err := client.NewRequest("GET", "/system/safe_reboot")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.ResponseMessage
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runSystemReset(c *cli.Context) {
	req, err := client.NewRequest("GET", "/system/reset_config")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.ResponseMessage
	err = req.ToJSON(&result)
	fmt.Println(result)
}

var actionsOfBacklight = []string{"lock", "release", "wakeup"}

func runSystemBacklight(c *cli.Context) {
	var a = c.String("action")
	var b bool
	for _, i := range actionsOfBacklight {
		b = a == i
		if b {
			break
		}
	}
	if b {
		url := fmt.Sprintf("/system/%s_backlight", a)
		req, err := client.NewRequest("GET", url)
		if err != nil {
			fmt.Println(err)
		}
		req.SetCookie(config.Cookie)
		var result api.ResponseMessage
		err = req.ToJSON(&result)
		fmt.Println(result)
	}
}

func runSystemTime(c *cli.Context) {
	req, err := client.NewRequest("GET", "/system/get_time")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.SystemTime
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runSystemCable(c *cli.Context) {
	req, err := client.NewRequest("GET", "/system/get_cable_connection")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.SystemCable
	err = req.ToJSON(&result)
	fmt.Println(result)
}
