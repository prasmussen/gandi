#!/bin/bash

# Markdown helpers
HEADER='### Downloads'
ROW_TEMPLATE='- [{{name}}]({{url}})'

# Application name and version
APP_NAME="gandi"
VERSION=$(bin/darwin-amd64/gandi-contact --version | awk '{print $2}' | sed -e 's/v//')

# Remove old archives
VOID=$(rm bin/*tgz 2>&1)

# Print markdown header
echo "$HEADER"

for PLATFORM_PATH in bin/*; do
    PLATFORM=$(basename $PLATFORM_PATH)
    FNAME="${APP_NAME}-${VERSION}-${PLATFORM}.tar.gz"
    FPATH="bin/${FNAME}"

    # Compress files
    tar -czf "${FPATH}" -C "${PLATFORM_PATH}" .

    # Upload file
    URL=$(drive upload -f $FPATH --share | awk '/https/ {print $9}')

    # Render markdown row and print to screen
    ROW=${ROW_TEMPLATE//"{{name}}"/$FNAME}
    ROW=${ROW//"{{url}}"/$URL}
    echo "$ROW"
done
