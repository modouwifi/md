package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bgentry/speakeasy"
	"github.com/codegangsta/cli"
	"github.com/modouwifi/md/api"
)

var cmdLogin = cli.Command{
	Name:   "login",
	Usage:  "log in of modou",
	Action: runLogin,
}

func runLogin(c *cli.Context) {
	// NOTE: gopass doesn't support multi-byte chars on Windows
	password, err := readPassword("Enter password: ")
	switch {
	case err == nil:
	case err.Error() == "unexpected newline":
		fmt.Println("password is required.")
	default:
		fmt.Println(err.Error())
	}

	err = attemptLogin(password)
}

var cmdLogout = cli.Command{
	Name:   "logout",
	Usage:  "log out of modou",
	Action: runLogout,
}

func runLogout(c *cli.Context) {
}

func readPassword(prompt string) (password string, err error) {
	if acceptPasswordFromStdin && !IsTerminal(os.Stdin) {
		_, err = fmt.Scanln(&password)
		return
	}
	// NOTE: speakeasy may not support multi-byte chars on Windows
	return speakeasy.Ask("Enter password: ")
}

func attemptLogin(password string) error {
	req, err := client.NewRequest("POST", "/auth/login")
	if err != nil {
		return err
	}
	body := struct {
		Password string `json:"password"`
	}{
		Password: password,
	}
	req.Body(body)
	var result api.ResponseMessage
	err = req.ToJSON(&result)
	if err != nil {
		return err
	}
	if result.Code != 0 {
		fmt.Println(result.Msg)
	} else {
		res := req.GetRes()
		rawCookie := res.Header.Get("Set-Cookie")
		if rawCookie != "" {
			strs := strings.Split(rawCookie, ";")
			cookie := strings.Split(strs[0], "=")
			path := strings.Split(strs[1], "=")[1]
			config.Cookie = &http.Cookie{
				Name:  cookie[0],
				Value: cookie[1],
				Path:  path,
			}
			config.Password = password
			err = config.Write()
			return err
		}
	}
	return nil
}
