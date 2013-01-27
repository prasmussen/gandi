package cli

import (
    "fmt"
    "github.com/prasmussen/gandi-api/domain/zone"
    "github.com/prasmussen/gandi/util"
)

type Zone struct {
    zone *zone.Zone
}

func New(d *zone.Zone) *Zone {
    return &Zone{d}
}

func (self *Zone) Count() {
    count, err := self.zone.Count()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Count:", count)
}

func (self *Zone) List() {
    zones, err := self.zone.List()
    if err != nil {
        fmt.Println(err)
        return
    }

    columns := make([]map[string]string, 0)
    order := []string{"Id", "Name", "Version", "Public", "Updated"}
    for _, zone := range zones {
        columns = append(columns, map[string]string{
            "Id": util.Itoa64(zone.Id),
            "Name": zone.Name,
            "Version": util.Itoa64(zone.Version),
            "Public": util.FormatBool(zone.Public),
            "Updated": util.TimeToLocal(zone.DateUpdated),
        })
    }

    util.PrintColumns(columns, order, 4)
}

func (self *Zone) Info(id int64) {
    info, err := self.zone.Info(id)
    if err != nil {
        fmt.Println(err)
        return
    }
    util.PrintStruct(info)
}

func (self *Zone) Create(name string) {
    info, err := self.zone.Create(name)
    if err != nil {
        fmt.Println(err)
        return
    }
    util.PrintStruct(info)
}

func (self *Zone) Delete(id int64) {
    deleted, err := self.zone.Delete(id)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Deleted: %s\n", util.FormatBool(deleted));
}

func (self *Zone) Set(domainName string, id int64) {
    info, err := self.zone.Set(domainName, id)
    if err != nil {
        fmt.Println(err)
        return
    }
    util.PrintStruct(info)
}
