package config

import (
    "fmt"
    "path/filepath"
    "io/ioutil"
    "encoding/json"
    "github.com/prasmussen/gandi/util"
)

var (
    DefaultConfigPath = filepath.Join(util.Homedir(), ".gandi", "config")
)

type Config struct {
    ApiTestKey string
    ApiProdKey string
}

func promptUser() *Config {
    fmt.Printf("Your API keys can be found here: https://www.gandi.net/admin/api_key\n\n")
    return &Config{
        ApiTestKey: util.Prompt("Enter API key for the test system: "),
        ApiProdKey: util.Prompt("Enter API key for the production system: "),
    }
}

func load(fname string) (*Config, error) {
    data, err := ioutil.ReadFile(fname)
    if err != nil {
        return nil, err
    }
    config := &Config{}
    return config, json.Unmarshal(data, config)
}

func save(fname string, config *Config) error {
    data, err := json.MarshalIndent(config, "", "    ")
    if err != nil {
        return err
    }

    if err = util.Mkdir(fname); err != nil {
        return err
    }
    return ioutil.WriteFile(fname, data, 0600)
}

func Load(fpath string) *Config {
    // Set default config path if non is provided
    if fpath == "" {
        fpath = DefaultConfigPath
    }

    // Try to load existing config
    config, err := load(fpath)
    if err != nil {
        // Unable to read existing config, lets start from scracth
        config = promptUser()
        
        // Save new config to file
        err := save(fpath, config)
        if err != nil {
            fmt.Printf("Failed to save config (%s)\n", err)
        }
    }
    return config
}
