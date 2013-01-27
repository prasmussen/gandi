#!/bin/bash

# Find handle associated to api key
HANDLE=$(gandi contact --testing info | awk '/^Handle:/ {print $2}')
echo "Got handle: ${HANDLE}"



# Find available domain
DOMAIN=""
for i in {0..9}; do
    DOMAIN="foobar${RANDOM}.com"
    STATUS=$(gandi domain --testing available -n ${DOMAIN} | awk '{print $2}')
    if [[ "$STATUS" == "available" ]]; then
        break
    fi

    if [ $i == 9 ]; then
        echo "Could not find available domain"
        exit 1
    fi
done

echo "Found available domain: ${DOMAIN}"



# Register domain
OPERATION_ID=$(gandi domain --testing create -n ${DOMAIN} -c ${HANDLE} -y 1 | awk '/^Id:/ {print $2}')



# Wait for operation to complete
echo "Waiting for domain registation to complete (id: ${OPERATION_ID})"
for i in {0..9}; do
    STATUS=$(gandi operation --testing info -i ${OPERATION_ID} | awk '/^Step:/ {print $2}')
    if [[ -z $STATUS ]]; then
        echo "Got invalid operation response from operation: ${OPERATION_ID}"
        exit 1
    fi

    if [[ $STATUS == "DONE" ]]; then
        break
    fi

    if [ $i == 9 ]; then
        echo "Unable to complete operation: ${OPERATION_ID}"
        exit 1
    fi

    sleep $i
done

echo "Domain registration complete"



# Create new zone
ZONE_ID=$(gandi zone --testing create -n ${DOMAIN} | awk '/^Id:/ {print $2}')

if [[ -z $ZONE_ID ]]; then
    echo "Unable to create zone"
    exit 1
fi

echo "Created new zone: ${ZONE_ID}"



# Create a new version since we cant modify the active version
ZONE_VERSION=$(gandi version --testing new -i $ZONE_ID | awk '{print $3}')

if [[ -z $ZONE_VERSION ]]; then
    echo "Unable to create new version of zone"
    exit 1
fi



# Add some records to zone
echo "Adding records to version: ${ZONE_VERSION}"
VOID=$(gandi record --testing add -i ${ZONE_ID} -v ${ZONE_VERSION} -n baz -t A -V 10.0.0.100 -T 3600)
VOID=$(gandi record --testing add -i ${ZONE_ID} -v ${ZONE_VERSION} -n baz -t AAAA -V 2001:0db8:85a3:0000:0000:8a2e:0370:7334 -T 3600)
VOID=$(gandi record --testing add -i ${ZONE_ID} -v ${ZONE_VERSION} -n qux -t CNAME -V bazqux.com. -T 3600)



# Set version active
echo "Activating version: ${ZONE_VERSION}"
STATUS=$(gandi version --testing set -i ${ZONE_ID} -v ${ZONE_VERSION})

if [[ "$STATUS" != "OK: True" ]]; then
    echo "Failed to set active version (${STATUS})"
    exit 1
fi



# Display records on active zone
echo "These are the active records for ${DOMAIN}"
echo
gandi record --testing list -i ${ZONE_ID}
