package cli

import (
    "fmt"
    "time"
    "strings"
    "github.com/prasmussen/gandi/util"
    "github.com/prasmussen/gandi-api/domain"
)

type Domain struct {
    domain *domain.Domain
}

func New(d *domain.Domain) *Domain {
    return &Domain{d}
}

func (self *Domain) Available(name string) {
    // Try three times with a 1 second delay to get domain status
    for i := 0; i < 3; i++ {
        status, err := self.domain.Available(name)
        if err != nil {
            fmt.Println(err)
            return
        }

        if status == "pending" {
            time.Sleep(time.Second * 1)
        } else {
            fmt.Println("Status:", status)
            return
        }
    }
    fmt.Println("Error: Could not get availability status within the given time, please try again")
}

func (self *Domain) Count() {
    count, err := self.domain.Count()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Count:", count)
}

func (self *Domain) List() {
    domains, err := self.domain.List()
    if err != nil {
        fmt.Println(err)
        return
    }

    columns := make([]map[string]string, 0)
    order := []string{"Name", "Created", "Expires", "Status"}
    for _, domain := range domains {
        columns = append(columns, map[string]string{
            "Name": domain.Fqdn,
            "Created": util.TimeToLocal(domain.DateCreated),
            "Expires": util.TimeToLocal(domain.DateRegistryEnd),
            "Status": strings.Join(domain.Status, ", "),
        })
    }

    util.PrintColumns(columns, order, 4)
}

func (self *Domain) Info(name string) {
    info, err := self.domain.Info(name)
    if err != nil {
        fmt.Println(err)
        return
    }
    util.PrintStruct(info)
}

func (self *Domain) Create(name, contact string, years int) {
    info, err := self.domain.Create(name, contact, years)
    if err != nil {
        fmt.Println(err)
        return
    }
    util.PrintStruct(info)
}
