package cli

import (
    "fmt"
    "strconv"
    "github.com/prasmussen/gandi-api/domain/zone/record"
    "github.com/prasmussen/gandi/util"
)

type Record struct {
    record *record.Record
}

func New(d *record.Record) *Record {
    return &Record{d}
}

func (self *Record) Count(zoneId, version int) {
    count, err := self.record.Count(zoneId, version)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Count:", count)
}

func (self *Record) List(zoneId, version int) {
    records, err := self.record.List(zoneId, version)
    if err != nil {
        fmt.Println(err)
        return
    }

    columns := make([]map[string]string, 0)
    order := []string{"Id", "Name", "TTL", "Type", "Value"}
    for _, record := range records {
        columns = append(columns, map[string]string{
            "Id": strconv.Itoa(record.Id),
            "Name": record.Name,
            "TTL": strconv.Itoa(record.Ttl),
            "Type": record.Type,
            "Value": record.Value,
        })
    }

    util.PrintColumns(columns, order, 4)
}

func (self *Record) Add(args record.RecordAdd) {
    info, err := self.record.Add(args)
    if err != nil {
        fmt.Println(err)
        return
    }
    util.PrintStruct(info)
}

func (self *Record) Delete(zoneId, version, recordId int) {
    deleted, err := self.record.Delete(zoneId, version, recordId)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Deleted: %s\n", util.FormatBool(deleted));
}

//func (self *Record) Set(domainName string, id int) {
//    info, err := self.record.Set(domainName, id)
//    if err != nil {
//        fmt.Println(err)
//        return
//    }
//    util.PrintStruct(info)
//}
