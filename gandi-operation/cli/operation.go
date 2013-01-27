package cli

import (
    "fmt"
    "github.com/prasmussen/gandi-api/operation"
    "github.com/prasmussen/gandi/util"
)

type Operation struct {
    operation *operation.Operation
}

func New(o *operation.Operation) *Operation {
    return &Operation{o}
}

func (self *Operation) Count() {
    count, err := self.operation.Count()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Count:", count)
}

func (self *Operation) Info(id int64) {
    info, err := self.operation.Info(id)
    if err != nil {
        fmt.Println(err)
        return
    }
    util.PrintStruct(info)
}

func (self *Operation) Cancel(id int64) {
    ok, err := self.operation.Cancel(id)
    if err != nil {
        fmt.Println(err)
        return
    }

    if ok {
        fmt.Printf("Operation %d has been canceled\n", id)
    } else {
        fmt.Printf("Failed to cancel operation %d\n", id)
    }
}

func (self *Operation) List() {
    operations, err := self.operation.List()
    if err != nil {
        fmt.Println(err)
        return
    }

    columns := make([]map[string]string, 0)
    order := []string{"Id", "Source", "Step", "Type", "Created", "Updated"}
    for _, operation := range operations {
        columns = append(columns, map[string]string{
            "Id": util.Itoa64(operation.Id),
            "Source": operation.Source,
            "Step": operation.Step,
            "Type": operation.Type,
            "Created": util.TimeToLocal(operation.DateCreated),
            "Updated": util.TimeToLocal(operation.DateUpdated),
        })
    }
    util.PrintColumns(columns, order, 4)
}
