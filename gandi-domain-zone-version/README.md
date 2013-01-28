gandi-domain-zone-version
=====


## Usage
    Usage: gandi-domain-zone-version [global options] <verb> [verb options]
    
    Global options:
        -t, --testing Perform queries against the test platform (OT&E)
        -c, --config  Set config path. Defaults to ~/.gandi/config
        -v, --version Print version
        -h, --help    Show this help
    
    Verbs:
        count:
            -z, --zone Zone id (*)
        delete:
            -z, --zone    Zone id (*)
            -v, --version Zone version (*)
        list:
            -z, --zone Zone id (*)
        new:
            -z, --zone    Zone id (*)
            -v, --version Zone version
        set:
            -z, --zone    Zone id (*)
            -v, --version Zone version (*)
    
