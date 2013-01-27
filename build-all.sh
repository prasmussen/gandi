#!/bin/bash

# Load crosscompile environment
source /Users/pii/scripts/golang-crosscompile/crosscompile.bash

PLATFORMS="darwin/386 darwin/amd64 freebsd/386 freebsd/amd64 linux/386 linux/amd64 linux/arm linux/arm5 windows/386 windows/amd64"
APPS="gandi gandi-contact gandi-domain gandi-domain-zone gandi-domain-zone-record gandi-domain-zone-version gandi-operation"

# Remove old binaries
rm -rvf bin/*

for APP_NAME in $APPS; do
    # Build binary for each platform in parallel
    for PLATFORM in $PLATFORMS; do
        GOOS=${PLATFORM%/*}
        GOARCH=${PLATFORM#*/}
        BIN_NAME=$APP_NAME
        BIN_PATH="bin/${GOOS}-${GOARCH}"

        if [ $GOOS == "windows" ]; then
            BIN_NAME="${APP_NAME}.exe"
        fi

        if [ $GOARCH == "arm5" ]; then
            export GOARM=5
            GOARCH="arm"
        else
            unset GOARM
        fi

        mkdir -p $BIN_PATH
        BUILD_CMD="go-${GOOS}-${GOARCH} build -o ${BIN_PATH}/${BIN_NAME} ${APP_NAME}/${APP_NAME}.go"

        echo "Building ${BIN_PATH}/${BIN_NAME}"
        $BUILD_CMD &
    done

    # Wait for builds to complete
    for job in $(jobs -p); do
        wait $job
    done
    echo "------------------------------------"
    echo
done

echo "All done"
