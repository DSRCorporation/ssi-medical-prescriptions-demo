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

userId="v0001"
localKMSPassphrase="Np6VR4Yg6PPL"

# Create a wallet
echo "Create pharmacy's wallet"
response=$(curl --location --request POST 'http://localhost:10082/vcwallet/create-profile' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "'$userId'",
    "localKMSPassphrase": "'$localKMSPassphrase'"
}')
is_ok "$response"

divider

echo "Open pharmacy's wallet"
response=$(curl --location --request POST 'http://localhost:10082/vcwallet/open' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "'$userId'",
    "localKMSPassphrase": "'$localKMSPassphrase'"
}')
is_ok "$response"

divider

# Close a wallet
echo "Close pharmacy's wallet"
response=$(curl --location --request POST 'http://localhost:10082/vcwallet/close' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "'$userId'",
    "localKMSPassphrase": "'$localKMSPassphrase'"
}')
check_response "$response" "true"

divider

echo "Successfully! Has opened pharmacy's wallet"
