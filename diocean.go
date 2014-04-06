package main

import (
  "flag"
  "os"
  "fmt"
  "github.com/kyleburton/diocean-go"
)

func main() {
	flag.BoolVar(&diocean.CmdlineOptions.Verbose, "v", false, "Verbose")
	flag.BoolVar(&diocean.CmdlineOptions.WaitForEvent, "w", false, "For commands that return an event_id, wait for the event to complete.")
	diocean.InitRoutingTable()
	flag.Parse()
	route := diocean.FindMatchingRoute(flag.Args())
	if diocean.CmdlineOptions.Verbose {
		fmt.Fprintf(os.Stderr, "Args: %s\n", flag.Args())
	}
	if diocean.CmdlineOptions.Verbose {
		fmt.Fprintf(os.Stderr, "Route: %s\n", route)
	}

	if route == nil {
		fmt.Fprintf(os.Stderr, "Error: unrecognized command: %s\n", flag.Args())
		diocean.ShowGeneralHelp(route)
		os.Exit(1)
	}

	if !diocean.InitConfig() {
		fmt.Fprintf(os.Stderr, "Invalid or Missing configuration file.\n")
		os.Exit(1)
	}
	if diocean.CmdlineOptions.Verbose {
		fmt.Fprintf(os.Stderr, "Config: %s\n", diocean.Config)
	}

	if route != nil {
		if diocean.CmdlineOptions.Verbose {
			fmt.Fprintf(os.Stderr, "Calling route: %s\n", route)
		}
		route.Handler(route)
	}
}
