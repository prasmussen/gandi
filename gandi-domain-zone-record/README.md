gandi-domain-zone-record
=====


## Usage
    Usage: gandi-domain-zone-record [global options] <verb> [verb options]
    
    Global options:
        -t, --testing Perform queries against the test platform (OT&E)
        -c, --config  Set config path. Defaults to ~/.gandi/config
        -v, --version Print version
        -h, --help    Show this help
    
    Verbs:
        add:
            -z, --zone    Zone id (*)
            -v, --version Zone version (*)
            -n, --name    Record name. Relative name, may contain leading wildcard. @ for empty name (*)
            -t, --type    Record type (*)
            -V, --value   Value for record. Semantics depends on the record type. (*)
            -T, --ttl     Time to live, in seconds, between 5 minutes and 30 days
        count:
            -z, --zone    Zone id (*)
            -v, --version Zone version
        delete:
            -z, --zone    Zone id (*)
            -v, --version Zone version (*)
            -r, --record  Record id (*)
        list:
            -z, --zone    Zone id (*)
            -v, --version Zone version
    
