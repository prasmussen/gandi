gandi-contact
=====


## Usage
    Usage: gandi-contact [global options] <verb> [verb options]
    
    Global options:
        -t, --testing Perform queries against the test platform (OT&E)
        -c, --config  Set config path. Defaults to ~/.gandi/config
        -v, --version Print version
        -h, --help    Show this help
    
    Verbs:
        balance:
        create:
                --firstname   First name (*)
                --lastname    Last name (*)
                --email       Email address (*)
                --password    Password (*)
                --address     Street address (*)
                --zipcode     Zip code (*)
                --city        City (*)
                --country     Country (*)
                --phone       Phone number (*)
                --person      Contact type person (*)
                --company     Contact type company (*)
                --association Contact type association (*)
                --publicbody  Contact type public body (*)
                --reseller    Contact type reseller (*)
        delete:
            -c, --contact Contact handle, defaults to the contact represented by apikey
        info:
            -c, --contact Contact handle, defaults to the contact represented by apikey
    
