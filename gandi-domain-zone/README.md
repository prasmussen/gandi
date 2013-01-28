gandi-domain-zone
=====


## Usage
    Usage: gandi-domain-zone [global options] <verb> [verb options]
    
    Global options:
        -t, --testing Perform queries against the test platform (OT&E)
        -c, --config  Set config path. Defaults to ~/.gandi/config
        -v, --version Print version
        -h, --help    Show this help
    
    Verbs:
        count:
        create:
            -n, --name Zone name (*)
        delete:
            -z, --zone Zone id (*)
        info:
            -z, --zone Zone id (*)
        list:
        set:
            -z, --zone Zone id (*)
            -n, --name Domain name (*)
    
