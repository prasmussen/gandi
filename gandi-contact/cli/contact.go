package cli

import (
    "fmt"
    "github.com/prasmussen/gandi-api/contact"
    "github.com/prasmussen/gandi/util"
)


type Contact struct {
    contact *contact.Contact
}

func New(d *contact.Contact) *Contact {
    return &Contact{d}
}

func (self *Contact) Balance() {
    info, err := self.contact.Balance()
    if err != nil {
        fmt.Println(err)
        return
    }
    util.PrintStruct(info)
}

func (self *Contact) Info(handle string) {
    info, err := self.contact.Info(handle)
    if err != nil {
        fmt.Println(err)
        return
    }
    util.PrintStruct(info)
}

func (self *Contact) Delete(handle string) {
    ok, err := self.contact.Delete(handle)
    if err != nil {
        fmt.Println(err)
        return
    }
   
    if ok { 
        fmt.Printf("Deleted %s\n", handle)
    } else {
        fmt.Printf("Failed to delete %s\n", handle)
    }
}

func (self *Contact) Create(options contact.ContactCreate) {
    info, err := self.contact.Create(options)
    if err != nil {
        fmt.Println(err)
        return
    }
    util.PrintStruct(info)
}
