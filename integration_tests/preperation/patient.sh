#
# Copyright 2022 DSR Corporation, Denver, Colorado.
# https://www.dsr-corporation.com
#
# This file is part of ssi-medical-prescriptions-demo.
#
# ssi-medical-prescriptions-demo is free software: you can redistribute it
# and/or modify it under the terms of the GNU Affero General Public License
# as published by the Free Software Foundation, either version 3 of the License,
# or (at your option) any later version.
#
# ssi-medical-prescriptions-demo is distributed in the hope that it will be
# useful, but WITHOUT ANY WARRANTY; without even the implied warranty
# of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
# See the GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License along
# with ssi-medical-prescriptions-demo. If not, see <https://www.gnu.org/licenses/>.
#

set -euo pipefail
source integration_tests/preperation/common.sh

userId="p0001"
localKMSPassphrase="4ELnzCgVBcDP"

# Create a wallet profile
echo "Create patient's wallet profile"
response=$(curl --location --request POST 'http://localhost:9082/vcwallet/create-profile' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "'$userId'",
    "localKMSPassphrase": "'$localKMSPassphrase'"
}')
is_ok "$response"

divider

# Open a wallet
echo "Open patient's wallet"
response=$(curl --location --request POST 'http://localhost:9082/vcwallet/open' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "'$userId'",
    "localKMSPassphrase": "'$localKMSPassphrase'"
}')
check_response "$response" "\"token\":"

divider

# Get tokenId from response
tokenId=$(echo "$response" | jq -r '.token')

verification_id="did:cheqd:testnet:z4TRtQjhizgKYVR6tLr42RQmyxfAPZDB#key-8g5iyxylyh"
privateKeyBase58="5SsDgS3P89XRW9wdKKm5UVdLP6EJj4NW2YWzLawXFYikfHfRmKwepyZ4bGGvp8jQ2A5mggk7fLT95YHxxwrC8pgM"

# Add DID method private key to a wallet
echo "Add DID method private key to patient's wallet"
curl --location --request POST 'http://localhost:9082/vcwallet/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "auth": "'$tokenId'",
    "content": {
        "id": "'$verification_id'",
        "type": "ed25519verificationkey2018",
        "privateKeyBase58": "'$privateKeyBase58'"
    },
    "contentType": "key",
    "userID": "'$userId'"
}'
is_ok "$response"

divider

# Close a wallet
echo "Close patient's wallet"
response=$(curl --location --request POST 'http://localhost:9082/vcwallet/close' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "'$userId'",
    "localKMSPassphrase": "'$localKMSPassphrase'"
}')
check_response "$response" "true"

divider

echo "Successfully! Has added patient's DID and key"
