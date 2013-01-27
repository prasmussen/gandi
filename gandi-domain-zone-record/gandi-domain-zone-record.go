package main

import (
    "fmt"
    "github.com/voxelbrain/goptions"
    "github.com/prasmussen/gandi/shared"
    "github.com/prasmussen/gandi/gandi-domain-zone-record/cli"
    api "github.com/prasmussen/gandi-api/domain/zone/record"
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

    Add api.RecordAdd `goptions:"add"`

    Count struct {
        Id int64 `goptions:"-i, --id, obligatory, description='Zone id'"`
        Version int64 `goptions:"-v, --version, description='Zone version'"`
    } `goptions:"count"`

    List struct {
        Id int64 `goptions:"-i, --id, obligatory, description='Zone id'"`
        Version int64 `goptions:"-v, --version, description='Zone version'"`
    } `goptions:"list"`

    Delete struct {
        Id int64 `goptions:"-i, --id, obligatory, description='Zone id'"`
        Version int64 `goptions:"-v, --version, obligatory, description='Zone version'"`
        Record int64 `goptions:"-r, --record, obligatory, description='Record id'"`
    } `goptions:"delete"`
}

func main() {
    opts := &Options{}
    goptions.ParseAndFail(opts)

    // Print version number and exit if the version flag is set
    if opts.Version {
        fmt.Printf("gandi-domain-zone-record v%s\n", VersionNumber)
        return
    }

    // Get gandi client
    client := shared.NewGandiClient(opts.ConfigPath, opts.Testing)

    // Create api and zone instances
    api := api.New(client)
    record := cli.New(api)

    switch opts.Verbs {
        case "count":
            record.Count(opts.Count.Id, opts.Count.Version)

        case "list":
            record.List(opts.List.Id, opts.List.Version)

        case "add":
            record.Add(opts.Add)

        case "delete":
            args := opts.Delete
            record.Delete(args.Id, args.Version, args.Record)

        default:
            goptions.PrintHelp()
    }
}
