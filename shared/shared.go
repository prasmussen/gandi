package shared

import (
    "os"
    "fmt"
    "github.com/prasmussen/gandi/config"
    "github.com/prasmussen/gandi-api/client"
)

const (
    VersionNumber = "1.0.2"
)

type NoArgs struct {}

func NewGandiClient(configPath string, testing bool) *client.Client {
    // Load config    
    cfg := config.Load(configPath)

    var apiKey string
    var systemType client.SystemType

    // Use test system and api key if the Testing flag was provided
    if testing {
        apiKey = cfg.ApiTestKey
        systemType = client.Testing
    } else {
        apiKey = cfg.ApiProdKey
        systemType = client.Production
    }

    // Create gandi client
    c, err := client.New(apiKey, systemType)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return c
}
