package main

import (
	"fmt"
	"os"

	"github.com/bgentry/speakeasy"
	"github.com/heroku/hk/term"
)

var cmdLogin = &Command{
	Run:   runLogin,
	Usage: "login",
	Short: "log in of modou",
	Long: `
Login with geek mode's password.

Example:

    $ md login
    Enter password:

`,
}

func runLogin(cmd *Command, args []string) error {
	if len(args) != 0 {
		cmd.printUsage(false)
		os.Exit(2)
	}

	// NOTE: gopass doesn't support multi-byte chars on Windows
	password, err := readPassword("Enter password: ")
	switch {
	case err == nil:
	case err.Error() == "unexpected newline":
		printFatal("password is required.")
	default:
		printFatal(err.Error())
	}

	err = attemptLogin(password)
	return nil
}

func readPassword(prompt string) (password string, err error) {
	if acceptPasswordFromStdin && !term.IsTerminal(os.Stdin) {
		_, err = fmt.Scanln(&password)
		return
	}
	// NOTE: speakeasy may not support multi-byte chars on Windows
	return speakeasy.Ask("Enter password: ")
}

func attemptLogin(password string) (err error) {
	req, err := client.NewRequest("POST", "/auth/login")
	if err != nil {
		return err
	}
	var data struct {
		Password string `json:"password"`
	}
	data.Password = password
	req.Body(data)
	resp, err := req.Do()
  if err != nil {
    return err
  }
  fmt.Println(resp)
	return nil
}

var cmdLogout = &Command{
	Run:   runLogout,
	Usage: "logout",
	Short: "log out of modou",
	Long: `
Log out

Example:

    $ md logout
    Logged out
`,
}

func runLogout(cmd *Command, args []string) error {
	if len(args) != 0 {
		cmd.printUsage(false)
		os.Exit(2)
	}
	return nil
}
