package cli

import (
    "fmt"
    "github.com/prasmussen/gandi-api/domain/zone/version"
    "github.com/prasmussen/gandi/util"
)


type Version struct {
    version *version.Version
}

func New(d *version.Version) *Version {
    return &Version{d}
}

func (self *Version) Count(zoneId int64) {
    count, err := self.version.Count(zoneId)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Count:", count)
}

func (self *Version) List(zoneId int64) {
    versions, err := self.version.List(zoneId)
    if err != nil {
        fmt.Println(err)
        return
    }

    columns := make([]map[string]string, 0)
    order := []string{"Id", "Date Created"}
    for _, version := range versions {
        columns = append(columns, map[string]string{
            "Id": util.Itoa64(version.Id),
            "Date Created": util.TimeToLocal(version.DateCreated),
        })
    }

    util.PrintColumns(columns, order, 4)
}

func (self *Version) New(zoneId, version int64) {
    version, err := self.version.New(zoneId, version)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Created version %d\n", version)
}

func (self *Version) Delete(zoneId, version int64) {
    deleted, err := self.version.Delete(zoneId, version)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Deleted: %s\n", util.FormatBool(deleted));
}

func (self *Version) Set(zoneId, version int64) {
    ok, err := self.version.Set(zoneId, version)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("OK: %s\n", util.FormatBool(ok));
}
