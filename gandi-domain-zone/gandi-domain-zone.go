package main

import (
    "fmt"
    "github.com/voxelbrain/goptions"
    "github.com/prasmussen/gandi/shared"
    "github.com/prasmussen/gandi/gandi-domain-zone/cli"
    api "github.com/prasmussen/gandi-api/domain/zone"
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

    Count shared.NoArgs `goptions:"count"`
    List shared.NoArgs `goptions:"list"`

    Info struct {
        Zone int64 `goptions:"-z, --zone, obligatory, description='Zone id'"`
    } `goptions:"info"`

    Create struct {
        Name string `goptions:"-n, --name, obligatory, description='Zone name'"`
    } `goptions:"create"`

    Delete struct {
        Zone int64 `goptions:"-z, --zone, obligatory, description='Zone id'"`
    } `goptions:"delete"`

    Set struct {
        Zone int64 `goptions:"-z, --zone, obligatory, description='Zone id'"`
        Name string `goptions:"-n, --name, obligatory, description='Domain name'"`
    } `goptions:"set"`
}

func main() {
    opts := &Options{}
    goptions.ParseAndFail(opts)

    // Print version number and exit if the version flag is set
    if opts.Version {
        fmt.Printf("gandi-domain-zone v%s\n", VersionNumber)
        return
    }

    // Get gandi client
    client := shared.NewGandiClient(opts.ConfigPath, opts.Testing)

    // Create api and zone instances
    api := api.New(client)
    zone := cli.New(api)

    switch opts.Verbs {
        case "count":
            zone.Count()

        case "list":
            zone.List()

        case "info":
            zone.Info(opts.Info.Zone)

        case "create":
            zone.Create(opts.Create.Name)

        case "delete":
            zone.Delete(opts.Delete.Zone)

        case "set":
            zone.Set(opts.Set.Name, opts.Set.Zone)

        default:
            goptions.PrintHelp()
    }
}

