#!/bin/bash

README_TEMPLATE='{{name}}
=====


## Usage
{{usage}}'

APPS="gandi gandi-contact gandi-domain gandi-domain-zone gandi-domain-zone-record gandi-domain-zone-version gandi-operation"

for NAME in $APPS; do
    USAGE=$($NAME --help 2>&1 | sed -e 's/^/    /')
    README=${README_TEMPLATE//"{{name}}"/$NAME}
    README=${README//"{{usage}}"/$USAGE}

    echo "Making ${NAME}/README.md"
    echo "$README" > "${NAME}/README.md"
done
