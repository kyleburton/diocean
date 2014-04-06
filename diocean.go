package main

import (
  "flag"
  "os"
  "fmt"
  "strings"
  "io/ioutil"
  "encoding/json"
  "github.com/kyleburton/diocean-go"
)

var Client *diocean.DioceanClient

type ConfigType map[string]string

var Config ConfigType

type CmdlineOptionsStruct struct {
	Verbose       bool
	WaitForEvents bool
}

var CmdlineOptions CmdlineOptionsStruct

type RouteHandler func (*Route)

type Route struct {
	Pattern  []string
	Params   map[string]string
	Args     []string
	Handler  RouteHandler
	HelpText *string
}

var RoutingTable []*Route

func InitRoutingTable() {
	RoutingTable = make([]*Route, 0)

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"sizes", "ls"},
		Params:  make(map[string]string),
		Handler: DropletSizesLs,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "ls", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsLsDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "show", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsLsDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "reboot", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsRebootDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "power-cycle", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsPowerCycleDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "shut-down", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsShutDownDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "shutdown", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsShutDownDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "power-off", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsPowerOffDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "poweroff", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsPowerOffDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "power-on", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsPowerOnDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "poweron", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsPowerOnDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "password-reset", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsPasswordResetDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "resize", ":droplet_id", ":size"},
		Params:  make(map[string]string),
		Handler: DoDropletsResizeDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "snapshot", ":droplet_id", ":name"},
		Params:  make(map[string]string),
		Handler: DoDropletsSnapshotDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "snapshot", ":droplet_id"},
		Params:  make(map[string]string),
		Handler: DoDropletsSnapshotDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "new", ":name", ":size", ":image", ":region", ":ssh_key_ids", ":private_networking", ":backups_enabled"},
		Params:  make(map[string]string),
		Handler: DoDropletsNewDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "destroy", ":droplet_id", ":scrub_data"},
		Params:  make(map[string]string),
		Handler: DoDropletsDestroyDroplet,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"droplets", "ls"},
		Params:  make(map[string]string),
		Handler: DoDropletsLs,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"images", "ls"},
		Params:  make(map[string]string),
		Handler: DoImagesLs})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"images", "show", ":image_id"},
		Params:  make(map[string]string),
		Handler: DoImageShow,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"images", "destroy", ":image_id"},
		Params:  make(map[string]string),
		Handler: DoImageDestroy,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"events", "show", ":event_id"},
		Params:  make(map[string]string),
		Handler: DoEventShow,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"events", "wait", ":event_id"},
		Params:  make(map[string]string),
		Handler: DoEventWait,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"regions", "ls"},
		Params:  make(map[string]string),
		Handler: DoRegionsLs,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"ssh-keys", "ls"},
		Params:  make(map[string]string),
		Handler: DoSshKeysLs,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"ssh", "fix-known-hosts"},
		Params:  make(map[string]string),
		Handler: DoSshFixKnownHosts,
	})

	RoutingTable = append(RoutingTable, &Route{
		Pattern: []string{"help"},
		Params:  make(map[string]string),
		Handler: ShowGeneralHelp,
	})

}

func ShowGeneralHelp(route *Route) {
	fmt.Printf("diocean <command> [arg1 [arg2 ..]] \n")
	fmt.Printf("  Commands:\n")
	for _, route := range RoutingTable {
		fmt.Printf("    %s\n", strings.Join(route.Pattern, "\t"))
		if route.HelpText != nil {
			fmt.Printf("\n")
			fmt.Print(route.HelpText)
			fmt.Printf("\n")
		}
	}
}

func RouteMatches(route *Route, args []string) (*Route, bool) {
	if CmdlineOptions.Verbose {
		fmt.Fprintf(os.Stderr, "Route: %s args: %s\n", route, args)
	}
	if len(args) < len(route.Pattern) {
		return nil, false
	}
	var res *Route = &Route{
		Pattern: route.Pattern,
		Params:  make(map[string]string),
		Handler: route.Handler,
	}

	for idx, part := range route.Pattern {
		arg := args[idx]
		res.Args = args[idx:]
		if CmdlineOptions.Verbose {
			fmt.Fprintf(os.Stderr, "  part:%s arg:%s rest:%s\n", part, arg, res.Args)
		}
		if strings.HasPrefix(part, ":") {
			res.Params[part[1:]] = arg
			continue
		}

		if part == arg {
			continue
		}

		if CmdlineOptions.Verbose {
			fmt.Fprintf(os.Stderr, "  ran out of parts at idx=%d, no match\n", idx)
		}
		return nil, false
	}

	return res, true
}

func FindMatchingRoute(args []string) *Route {
	for _, route := range RoutingTable {
		res, matched := RouteMatches(route, args)
		if matched {
			return res
		}
	}
	return nil
}


func InitConfig() bool {
	file, e := ioutil.ReadFile(os.Getenv("HOME") + "/.digitalocean.json")
	if e != nil {
		fmt.Fprintf(os.Stderr, "File error: %v\n", e)
		return false
	}

	json.Unmarshal(file, &Config)

	if _, ok := Config["ClientId"]; !ok {
		fmt.Fprintf(os.Stderr, "Error: No ClienId in configuration file!\n", e)
		return false
	}

	if _, ok := Config["ApiKey"]; !ok {
		fmt.Fprintf(os.Stderr, "Error: No ApiKey in configuration file!\n", e)
		return false
	}

	return true
}

////////////////////////////////////////////////////////////////////////////////
func DropletSizesLs (route *Route) {
  Client.DropletSizesLs()
}

func DoDropletsLsDroplet (route *Route) {
  Client.DoDropletsLsDroplet(route.Params["droplet_id"])
}

func DoDropletsRebootDroplet (route *Route) {
  Client.DoDropletsRebootDroplet(route.Params["droplet_id"])
}

func DoDropletsPowerCycleDroplet (route *Route) {
  Client.DoDropletsPowerCycleDroplet(route.Params["droplet_id"])
}

func DoDropletsShutDownDroplet (route *Route) {
  Client.DoDropletsShutDownDroplet(route.Params["droplet_id"])
}

func DoDropletsPowerOffDroplet (route *Route) {
  Client.DoDropletsPowerOffDroplet(route.Params["droplet_id"])
}

func DoDropletsPowerOnDroplet (route *Route) {
  Client.DoDropletsPowerOnDroplet(route.Params["droplet_id"])
}

func DoDropletsPasswordResetDroplet (route *Route) {
  Client.DoDropletsPasswordResetDroplet(route.Params["droplet_id"])
}

func DoDropletsResizeDroplet (route *Route) {
  Client.DoDropletsResizeDroplet(route.Params["droplet_id"], route.Params["size"])
}

func DoDropletsSnapshotDroplet (route *Route) {
  Client.DoDropletsSnapshotDroplet(route.Params["droplet_id"], route.Params["name"])
}

func DoDropletsNewDroplet (route *Route) {
  Client.DoDropletsNewDroplet(
    route.Params["name"],
    route.Params["size"],
    route.Params["image"],
    route.Params["region"],
    route.Params["ssh_key_ids"],
    route.Params["private_networking"],
    route.Params["backups_enabled"],
  )
}

func DoDropletsDestroyDroplet (route *Route) {
  Client.DoDropletsDestroyDroplet(route.Params["droplet_id"], route.Params["scrub_data"])
}

func DoDropletsLs (route *Route) {
  Client.DoDropletsLs()
}

func DoImagesLs (route *Route) {
  Client.DoImagesLs()
}

func DoImageShow (route *Route) {
  Client.DoImageShow(route.Params["image_id"])
}


func DoImageDestroy (route *Route) {
  Client.DoImageDestroy(route.Params["image_id"])
}

func DoEventShow (route *Route) {
  Client.DoEventShow(route.Params["event_id"])
}

func DoEventWait (route *Route) {
  Client.DoEventWait(route.Params["event_id"])
}

func DoRegionsLs (route *Route) {
  Client.DoRegionsLs()
}

func DoSshKeysLs (route *Route) {
  Client.DoSshKeysLs()
}

func DoSshFixKnownHosts (route *Route) {
  Client.DoSshFixKnownHosts()
}



////////////////////////////////////////////////////////////////////////////////


func main() {
	flag.BoolVar(&CmdlineOptions.Verbose, "v", false, "Verbose")
	flag.BoolVar(&CmdlineOptions.WaitForEvents, "w", false, "For commands that return an event_id, wait for the event to complete.")
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

  Client = &diocean.DioceanClient{
    ClientId:      Config["ClientId"],
    ApiKey:        Config["ApiKey"],
    Verbose:       CmdlineOptions.Verbose,
    WaitForEvents: CmdlineOptions.WaitForEvents,
  }

	if route != nil {
		if CmdlineOptions.Verbose {
			fmt.Fprintf(os.Stderr, "Calling route: %s\n", route)
		}
		route.Handler(route)
	}
}
