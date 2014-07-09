package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/bgentry/pflag"
)

type Command struct {
	// args does not include the command name
	Run  func(cmd *Command, args []string) error
	Flag flag.FlagSet

	Usage string // first word is the command name
	Short string // `md help` output
	Long  string // `md help cmd` output
}

func (c *Command) printUsage(errExit bool) {
	if c.Runnable() {
		fmt.Printf("Usage: %s %s\n\n", os.Args[0], c.Usage)
	}
	fmt.Println(strings.Trim(c.Long, "\n"))
	if errExit {
		os.Exit(2)
	}
}

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Runnable() bool {
	return c.Run != nil
}

func (c *Command) List() bool {
	return c.Short != ""
}
