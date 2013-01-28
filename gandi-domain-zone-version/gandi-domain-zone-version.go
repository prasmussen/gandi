package main

import (
    "fmt"
    "github.com/voxelbrain/goptions"
    "github.com/prasmussen/gandi/shared"
    "github.com/prasmussen/gandi/gandi-domain-zone-version/cli"
    api "github.com/prasmussen/gandi-api/domain/zone/version"
)

const (
    VersionNumber = "1.0.0"
)

type Options struct {
    Testing bool `goptions:"-t, --testing, description='Perform queries against the test platform (OT&E)'"`
    ConfigPath string `goptions:"-c, --config, description='Set config path. Defaults to ~/.gandi/config'"`
    Version bool `goptions:"-v, --version, description='Print version'"`
    goptions.Help `goptions:"-h, --help, description='Show this help'"`

    goptions.Verbs

    Count struct {
        Zone int64 `goptions:"-z, --zone, obligatory, description='Zone id'"`
    } `goptions:"count"`

    List struct {
        Zone int64 `goptions:"-z, --zone, obligatory, description='Zone id'"`
    } `goptions:"list"`

    Delete struct {
        Zone int64 `goptions:"-z, --zone, obligatory, description='Zone id'"`
        Version int64 `goptions:"-v, --version, obligatory, description='Zone version'"`
    } `goptions:"delete"`

    Set struct {
        Zone int64 `goptions:"-z, --zone, obligatory, description='Zone id'"`
        Version int64 `goptions:"-v, --version, obligatory, description='Zone version'"`
    } `goptions:"set"`

    New struct {
        Zone int64 `goptions:"-z, --zone, obligatory, description='Zone id'"`
        Version int64 `goptions:"-v, --version, description='Zone version'"`
    } `goptions:"new"`
}

func main() {
    opts := &Options{}
    goptions.ParseAndFail(opts)

    // Print version number and exit if the version flag is set
    if opts.Version {
        fmt.Printf("gandi-domain-zone-version v%s\n", VersionNumber)
        return
    }

    // Get gandi client
    client := shared.NewGandiClient(opts.ConfigPath, opts.Testing)

    // Create api and zone instances
    api := api.New(client)
    version := cli.New(api)

    switch opts.Verbs {
        case "count":
            version.Count(opts.Count.Zone)

        case "list":
            version.List(opts.List.Zone)

        case "new":
            version.New(opts.New.Zone, opts.New.Version)

        case "delete":
            version.Delete(opts.Delete.Zone, opts.Delete.Version)

        case "set":
            version.Set(opts.Set.Zone, opts.Set.Version)

        default:
            goptions.PrintHelp()
    }
}

