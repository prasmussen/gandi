package main

import (
    "fmt"
    "github.com/voxelbrain/goptions"
    "github.com/prasmussen/gandi/shared"
    "github.com/prasmussen/gandi/gandi-operation/cli"
    api "github.com/prasmussen/gandi-api/operation"
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
        Id int `goptions:"-i, --id, obligatory, description='Operation id'"`
    } `goptions:"info"`

    Cancel struct {
        Id int `goptions:"-i, --id, obligatory, description='Operation id'"`
    } `goptions:"cancel"`
}

func main() {
    opts := &Options{}
    goptions.ParseAndFail(opts)

    // Print version number and exit if the version flag is set
    if opts.Version {
        fmt.Printf("gandi-operation v%s\n", VersionNumber)
        return
    }

    // Get gandi client
    client := shared.NewGandiClient(opts.ConfigPath, opts.Testing)

    // Create api and operation instances
    api := api.New(client)
    operation := cli.New(api)

    switch opts.Verbs {
        case "count":
            operation.Count()

        case "list":
            operation.List()

        case "info":
            operation.Info(opts.Info.Id)

        case "cancel":
            operation.Cancel(opts.Cancel.Id)

        default:
            goptions.PrintHelp()
    }
}
