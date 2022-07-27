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

DEF_OUTPUT_MODE=json

_check_response() {
    local _result="$1"
    local _expected_string="$2"
    local _mode="${3:-$DEF_OUTPUT_MODE}"

    if [[ "$_mode" == "json" ]]; then
        if [[ -n "$(echo "$_result" | jq | grep "$_expected_string" 2>/dev/null)" ]]; then
            echo true
            return
        fi
    else
        if [[ -n "$(echo "$_result" | grep "$_expected_string" 2>/dev/null)" ]]; then
            echo true
            return
        fi
    fi

    echo false
}

GREEN=""
RESET=""
OK=

check_response() {
    local _result="$1"
    local _expected_string="$2"
    local _mode="${3:-$DEF_OUTPUT_MODE}"

    if [[ "$(_check_response "$_result" "$_expected_string" "$_mode")" != true ]]; then
        echo "${GREEN}ERROR:${RESET} command failed. The expected string: '$_expected_string' not found in the result: $_result"
        exit 1
    fi
}

is_ok() {
    local response="$1"

    if [[ -n "$response" ]]; 
    then 
        echo "Error: $response"
        exit 1
    fi
}

divider() {
  echo ""
  echo "----------------------------------------------------------------"
  echo ""
}
