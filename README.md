diocean
=======

Digital Ocean API Command Line Client

    diocean <command> [arg1 [arg2 ..]] 
      Commands:
        sizes  ls
        droplets  ls  :dropletId
        droplets  show  :dropletId
        droplets  reboot  :droplet_id
        droplets  power-cycle  :droplet_id
        droplets  shut-down  :droplet_id
        droplets  shutdown  :droplet_id
        droplets  power-off  :droplet_id
        droplets  poweroff  :droplet_id
        droplets  power-on  :droplet_id
        droplets  poweron  :droplet_id
        droplets  password-reset  :droplet_id
        droplets  resize  :droplet_id  :size
        droplets  snapshot  :droplet_id  :name
        droplets  snapshot  :droplet_id
        droplets  new  :name  :size  :image  :region  :ssh_key_ids  :private_networking  :backups_enabled
        droplets  destroy  :droplet_id  :scrub_data
        droplets  ls
        images  ls
        images  show  :image_id
        images  destroy  :image_id
        events  show  :event_id
        events  wait  :event_id
        regions  ls
        ssh-keys  ls
        ssh  fix-known-hosts
        help

# Roadmap / *TODO*

Documentation: both a basic manual and help text for the application (link back to the on-line API documentation).

Factor the HTTP Api into a re-useable library, separate the command line interface and formatting of results into a separate module.  The command line interface deals with configuration, input and output.  The light-weight api abstracts the HTTP interface.

Support json output in addition to tab delimited output.  This would be useful in conjunction with tools like [jq](http://stedolan.github.io/jq/).

### Command Line Completion

- DONE bash wrapper
- DONE completion for route patterns
- DONE parameter expansion (eg: "droplets new test1 <TAB>" should list the available sizes since that is the next parameter)
- DONE implement caching to speed up completion


### API Support

- DONE Support a -wait flag for all operations that return an event id

- Droplets
    - DONE Show All Active Droplets
    - DONE New Droplet
    - DONE Show Droplet (sans backups and snapshots)
    - *TODO* Show Droplet: Backups
    - *TODO* Show Droplet: Snapshots
    - DONE Reboot Droplet
    - DONE Power Cycle Droplet
    - DONE Shut Down Droplet
    - DONE Power Off
    - DONE Power On
    - DONE Reset Root Password
    - DONE Resize Droplet
    - DONE Take a Snapshot
    - *TODO* Restore Droplet
    - *TODO* Rebuild Droplet
    - *TODO* Rename Droplet
    - DONE Destroy Droplet

- Regions
    - DONE All Regions

- Images
    - DONE All Images
    - DONE Show Image
    - DONE Destroy Image
    - *TODO* Transfer Image

- SSH Keys
    - *TODO* All SSH Keys
    - *TODO* Add SSH Keys
    - *TODO* Show SSH Keys
    - *TODO* Edit SSH Keys
    - *TODO* Destroy SSH Keys

- Sizes
-- DONE All Sizes

- Domains
    - *TODO* All Domains
    - *TODO* New Domain
    - *TODO* Domain Show
    - *TODO* Destroy Domain
    - *TODO* All Domain Records
    - *TODO* New Domain Record
    - *TODO* Show Domain Record
    - *TODO* Edit Domain Record
    - *TODO* Destroy Domain Record

- Events
    - DONE Show Event
    - DONE Wait For Event (percentage=100)

### Tests

- _IN PROGRESS_ Route.CompletionsFor
- _IN PROGRESS_ FindCompletionWords
- *TODO* ParameterCompletions
- *TODO* AppendUnique
- *TODO* STripColonPrefix
- *TODO* ReadFromDiskCache
- *TODO* UseDiskCache

## References

- http://blog.equanimity.nl/blog/2013/05/29/a-beginners-guide-to-erlang/
- https://developers.digitalocean.com/sizes/
- https://developers.digitalocean.com/images/



