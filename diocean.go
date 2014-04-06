package main

import (
  "flag"
  "os"
  "fmt"
  "github.com/kyleburton/diocean"
)

func main() {
	flag.BoolVar(&diocean.CmdlineOptions.Verbose, "v", false, "Verbose")
	flag.BoolVar(&diocean.CmdlineOptions.WaitForEvent, "w", false, "For commands that return an event_id, wait for the event to complete.")
	InitRoutingTable()
	flag.Parse()
	route := FindMatchingRoute(flag.Args())
	if CmdlineOptions.Verbose {
		fmt.Fprintf(os.Stderr, "Args: %s\n", flag.Args())
	}
	if CmdlineOptions.Verbose {
		fmt.Fprintf(os.Stderr, "Route: %s\n", route)
	}

	if route == nil {
		fmt.Fprintf(os.Stderr, "Error: unrecognized command: %s\n", flag.Args())
		ShowGeneralHelp(route)
		os.Exit(1)
	}

	if !InitConfig() {
		fmt.Fprintf(os.Stderr, "Invalid or Missing configuration file.\n")
		os.Exit(1)
	}
	if CmdlineOptions.Verbose {
		fmt.Fprintf(os.Stderr, "Config: %s\n", Config)
	}

	if route != nil {
		if CmdlineOptions.Verbose {
			fmt.Fprintf(os.Stderr, "Calling route: %s\n", route)
		}
		route.Handler(route)
	}
}
