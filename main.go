package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	Version = "0.0.0"
	apiURL  = "http://modouwifi.net/api"
	mdAgent = "md/" + Version + " (" + runtime.GOOS + "; " + runtime.GOARCH + ")"
	stdin   = bufio.NewReader(os.Stdin)
)

// Running `md help` will list commands in this order.
var commands = []*Command{
	cmdLogin,
	cmdLogout,

	cmdVersion,
	cmdHelp,
}

func main() {
	log.SetFlags(0)

	args := os.Args[1:]

	if len(args) < 1 {
		usage()
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() {
				cmd.printUsage(false)
			}
			if err := cmd.Flag.Parse(args[1:]); err != nil {
				os.Exit(2)
			}
			if err := cmd.Run(cmd, cmd.Flag.Args()); err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown command: %s\n", args[0])
	usage()
}
