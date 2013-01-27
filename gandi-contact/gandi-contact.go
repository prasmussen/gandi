package main

import (
    "fmt"
    "github.com/voxelbrain/goptions"
    "github.com/prasmussen/gandi/shared"
    "github.com/prasmussen/gandi/gandi-contact/cli"
    api "github.com/prasmussen/gandi-api/contact"
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

    Balance shared.NoArgs `goptions:"balance"`

    Info struct {
        Contact string `goptions:"-c, --contact, description='Contact handle, defaults to the contact represented by apikey'"`
    } `goptions:"info"`

    Delete struct {
        Contact string `goptions:"-c, --contact, description='Contact handle, defaults to the contact represented by apikey'"`
    } `goptions:"delete"`

    Create api.ContactCreate `goptions:"create"`
}

func main() {
    opts := &Options{}
    goptions.ParseAndFail(opts)

    // Print version number and exit if the version flag is set
    if opts.Version {
        fmt.Printf("gandi-contact v%s\n", VersionNumber)
        return
    }

    // Get gandi client
    client := shared.NewGandiClient(opts.ConfigPath, opts.Testing)

    // Create api and cli instances
    api := api.New(client)
    contact := cli.New(api)

    switch opts.Verbs {
        case "balance":
            contact.Balance()

        case "info":
            contact.Info(opts.Info.Contact)

        case "create":
            contact.Create(opts.Create)

        case "delete":
            contact.Delete(opts.Delete.Contact)

        default:
            goptions.PrintHelp()
    }
}
