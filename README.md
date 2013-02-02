gandi
=====


## Overview
gandi is a collection of command line tools for interacting with the [gandi.net API](http://doc.rpc.gandi.net/).
It currently consist of the following tools: `gandi-contact` `gandi-operation` `gandi-domain` `gandi-domain-zone`
`gandi-domain-zone-record` and `gandi-domain-zone-version`. Included is also a wrapper tool `gandi` which can be
used instead of each tool separately (i.e `gandi zone list` will execute `gandi-domain-zone list`).

## Prerequisites
None, binaries are statically linked.
If you want to compile from source you need the [go toolchain](http://golang.org/doc/install)

## Installation
- Save all binaries to a location in your PATH (i.e. `/usr/local/bin/`)

### Downloads
- [gandi-1.0.3-darwin-386.tar.gz](https://drive.google.com/uc?id=0B3X9GlR6EmbnZzBaVTFoZ3h2N2s)
- [gandi-1.0.3-darwin-amd64.tar.gz](https://drive.google.com/uc?id=0B3X9GlR6Embnb25KSGRjSWF1ZUE)
- [gandi-1.0.3-freebsd-386.tar.gz](https://drive.google.com/uc?id=0B3X9GlR6EmbnX2NnTmhmMHNJZTA)
- [gandi-1.0.3-freebsd-amd64.tar.gz](https://drive.google.com/uc?id=0B3X9GlR6EmbnbGU0SkVpS3IyekU)
- [gandi-1.0.3-linux-386.tar.gz](https://drive.google.com/uc?id=0B3X9GlR6EmbnUXFTUDd4WkYtbUE)
- [gandi-1.0.3-linux-amd64.tar.gz](https://drive.google.com/uc?id=0B3X9GlR6EmbnX3NPbDZTUmJWZFk)
- [gandi-1.0.3-linux-arm.tar.gz](https://drive.google.com/uc?id=0B3X9GlR6EmbnUGhmT09UOE5BTVU)
- [gandi-1.0.3-linux-arm5.tar.gz](https://drive.google.com/uc?id=0B3X9GlR6EmbnLXFzbnl2Wi1BNE0)
- [gandi-1.0.3-windows-386.tar.gz](https://drive.google.com/uc?id=0B3X9GlR6EmbnMElzNndjM2V0MW8)
- [gandi-1.0.3-windows-amd64.tar.gz](https://drive.google.com/uc?id=0B3X9GlR6EmbnYjkzekRwVHVuSU0)

## First run
The first time one of the tools are executed it will prompt for your gandi API keys which can be found
[here](https://www.gandi.net/admin/api_key).
The API keys is by default stored in $HOME/.gandi/config (can be overridden with the -c flag).

## Examples
All examples are run against the test platform (OT&E). Remove the --testing flag to perform queries
against the production API.

###### Get contact information about self
    $ gandi contact --testing info
    Firstname: John
    Lastname: Doe
    Email: john.doe@gmail.com
    Address: Foobarstreet 32
    Zipcode: 1337
    City: Foo City
    Country: FB
    Phone: +1.555-0122
    ContactType: 0
    Handle: JD1337-GANDI

###### Check for domain availability
    $ gandi domain --testing available --domain bazqux.com
    Status: available

###### Register domain
    $ gandi domain --testing create --domain bazqux.com --contact JD1337-GANDI --years 1
    DateCreated: 2013-01-27 23:44:10 +0000 UTC
    DateStart: 0001-01-01 00:00:00 +0000 UTC
    DateUpdated: 2013-01-27 23:44:10 +0000 UTC
    Id: 19784
    SessionId: 8101
    Source: JD1337-GANDI
    Step: BILL
    Type: domain_create
    Params:
        auth_id: 8101
        tech: JD1337-GANDI
        duration: 1
        admin: JD1337-GANDI
        tld: com
        domain: bazqux.com
        owner: JD1337-GANDI
        param_type: domain
        ns: [a.dns.gandi-ote.net b.dns.gandi-ote.net c.dns.gandi-ote.net]
        remote_addr: 1.2.3.4
        bill: JD1337-GANDI
    OperationDetails:
        Label: bazqux.com
        ProductAction: create
        ProductName: com
        ProductType: domain
        Quantity: 0

###### Create new zone
    $ gandi zone --testing create --name bazqux.com
    DateUpdated: 2013-01-27 23:46:57 +0000 UTC
    Id: 681917
    Name: bazqux.com
    Public: false
    Version: 1
    Domains: 0
    Owner: JD1337-GANDI
    Versions: 1

###### Create a new zone version since we cant modify the active version
    $ gandi version --testing new --zone 681917
    Created version 2

###### Add A record to zone
    $ gandi record --testing add --zone 681917 --version 2 --name foo --type A --value 10.0.0.100 --ttl 3600
    Id: 3566725
    Name: foo
    Ttl: 3600
    Type: A
    Value: 10.0.0.100

###### Add AAAA record to zone
    $ gandi record --testing add --zone 681917 --version 2 --name foo --type AAAA --value 2001:0db8:85a3:0000:0000:8a2e:0370:7334 --ttl 3600
    Id: 3566726
    Name: foo
    Ttl: 3600
    Type: AAAA
    Value: 2001:db8:85a3::8a2e:370:7334

###### Add CNAME record to zone
    $ gandi record --testing add --zone 681917 --version 2 --name foobar --type CNAME --value foobar.com. --ttl 3600
    Id: 3566727
    Name: foobar
    Ttl: 3600
    Type: CNAME
    Value: foobar.com.

###### Set version 2 active
    $ gandi version --testing set --zone 681917 --version 2
    OK: True

###### List all record on active version
    $ gandi record --testing list --zone 681917
    Id         Name      TTL     Type     Value
    3566725    foo       3600    A        10.0.0.100
    3566726    foo       3600    AAAA     2001:db8:85a3::8a2e:370:7334
    3566727    foobar    3600    CNAME    foobar.com.

###### Set zone on domain
    $ gandi zone --testing set --zone 681917 --domain bazqux.com 
    AuthInfo: ******
    DateCreated: 2013-01-27 23:44:26 +0000 UTC
    DateRegistryCreation: 2013-01-27 22:44:25 +0000 UTC
    DateRegistryEnd: 2014-01-27 22:44:25 +0000 UTC
    DateUpdated: 2013-01-27 23:44:26 +0000 UTC
    Fqdn: bazqux.com
    Id: 2997
    Status: clientTransferProhibited
    Tld: com
    DateDelete: 2014-02-26 12:44:25 +0000 UTC
    DateHoldBegin: 2014-01-27 22:44:25 +0000 UTC
    DateHoldEnd: 2014-02-26 22:44:25 +0000 UTC
    DatePendingDeleteEnd: 2014-04-02 22:44:25 +0000 UTC
    DateRenewBegin: 2012-01-01 00:00:00 +0000 UTC
    DateRestoreEnd: 2014-03-28 22:44:25 +0000 UTC
    Nameservers: a.dns.gandi-ote.net, b.dns.gandi-ote.net, c.dns.gandi-ote.net
    Services: gandidns
    ZoneId: 681917
    Autorenew:
        Active: false
        Id: 0
        ProductId: 0
        ProductTypeId: 0
    Contacts:
        Admin:
            Handle: JD1337-GANDI
            Id: 1337
        Bill:
            Handle: JD1337-GANDI
            Id: 1337
        Owner:
            Handle: JD1337-GANDI
            Id: 1337
        Reseller:
            Id: 0
        Tech:
            Handle: JD1337-GANDI
            Id: 1337
