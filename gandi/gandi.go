package main

import (
    "os"
    "os/exec"
    "fmt"
    "strings"
)

var commands = map[string]string{
    "contact": "gandi-contact",
    "operation": "gandi-operation",
    "domain": "gandi-domain",
    "zone": "gandi-domain-zone",
    "record": "gandi-domain-zone-record",
    "version": "gandi-domain-zone-version",
}

func printHelp() {
    fmt.Printf("Usage: gandi <command> [args ...]\n\n")
    fmt.Println("Available commands:")
    for name, _ := range commands {
        fmt.Printf("  %s\n", name)
    }
    os.Exit(0)
}

func findBin(cmd string) string {
    binaries := make([]string, 0)
    for name, bin := range commands {
        if strings.Index(name, cmd) == 0 {
            binaries = append(binaries, bin)
        }
    }
    
    // Return empty string if no or more than one matches was found
    if len(binaries) != 1 {
        return ""
    }

    // Find path of binary
    binName := binaries[0]
    binPath, err := exec.LookPath(binName)
    if err != nil {
        fmt.Println("Error: Could not find binary:", binName)
        os.Exit(1)
    }

    return binPath
}

func main() {
    args := os.Args[1:]

    // Print help and exit if no command is given
    if len(args) < 1 {
        printHelp()
    }

    // Find matching binary, or print help if none is found
    binPath := findBin(strings.ToLower(args[0]))
    if binPath == "" {
        printHelp()
    }

    // Execute command
    cmd := exec.Command(binPath, args[1:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}
