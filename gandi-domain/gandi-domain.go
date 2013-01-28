package main

import (
    "fmt"
    "github.com/voxelbrain/goptions"
    "github.com/prasmussen/gandi/shared"
    "github.com/prasmussen/gandi/gandi-domain/cli"
    api "github.com/prasmussen/gandi-api/domain"
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
        Domain string `goptions:"-d, --domain, obligatory, description='Domain name'"`
    } `goptions:"info"`

    Available struct {
        Domain string `goptions:"-d, --domain, obligatory, description='Domain name'"`
    } `goptions:"available"`

    Create struct {
        Domain string `goptions:"-d, --domain, obligatory, description='Domain name'"`
        Contact string `goptions:"-c, --contact, obligatory, description='Contact handle'"`
        Years int64 `goptions:"-y, --years, obligatory, description='Years to register the domain for'"`
    } `goptions:"create"`
}

func main() {
    opts := &Options{}
    goptions.ParseAndFail(opts)

    // Print version number and exit if the version flag is set
    if opts.Version {
        fmt.Printf("gandi-domain v%s\n", shared.VersionNumber)
        return
    }

    // Get gandi client
    client := shared.NewGandiClient(opts.ConfigPath, opts.Testing)

    // Create api and cli instances
    api := api.New(client)
    domain := cli.New(api)

    switch opts.Verbs {
        case "count":
            domain.Count()

        case "list":
            domain.List()

        case "info":
            domain.Info(opts.Info.Domain)

        case "available":
            domain.Available(opts.Available.Domain)

        case "create":
            args := opts.Create
            domain.Create(args.Domain, args.Contact, args.Years)

        default:
            goptions.PrintHelp()
    }
}
