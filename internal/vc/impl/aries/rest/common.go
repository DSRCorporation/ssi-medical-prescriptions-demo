/*
  Copyright 2022 DSR Corporation, Denver, Colorado.
  https://www.dsr-corporation.com

  This file is part of ssi-medical-prescriptions-demo.

  ssi-medical-prescriptions-demo is free software: you can redistribute it
  and/or modify it under the terms of the GNU Affero General Public License
  as published by the Free Software Foundation, either version 3 of the License,
  or (at your option) any later version.

  ssi-medical-prescriptions-demo is distributed in the hope that it will be
  useful, but WITHOUT ANY WARRANTY; without even the implied warranty
  of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
  See the GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License along
  with ssi-medical-prescriptions-demo. If not, see <https://www.gnu.org/licenses/>.
*/

package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

func CreateOOBInvitation(client *resty.Client, endpoint string) (invitation json.RawMessage, err error) {
	resp, err := client.R().
		SetBody(struct {
			Label string `json:"label"`
		}{Label: "Issuer"}).
		Post(endpoint + "/outofband/create-invitation")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.Body(), nil
	} else {
		return nil, errors.New(string(resp.Body()))
	}
}

func AcceptOOBRequest(client *resty.Client, endpoint string, connectionId string) (connection domain.Connection, err error) {
	resp, err := client.R().
		SetPathParam("connectionId", connectionId).
		Post(endpoint + "/{connectionId}/accept-request")

	if err != nil {
		return domain.Connection{}, err
	}

	if resp.StatusCode() == http.StatusOK {
		var res map[string]interface{}
		resp, err := client.R().
			SetPathParam("connectionId", connectionId).
			SetResult(res).
			Get(endpoint + "/{connectionId}")

		if err != nil {
			return domain.Connection{}, err
		}

		if resp.StatusCode() == http.StatusOK {
			inviterDID := res["result"].(map[string]interface{})["MyDID"].(string)
			inviteeDID := res["result"].(map[string]interface{})["TheirDID"].(string)

			conn := domain.Connection{
				InviterDID:   inviterDID,
				InviteeDID:   inviteeDID,
				ConnectionId: connectionId,
			}

			return conn, nil
		} else {
			return domain.Connection{}, errors.New(string(resp.Body()))
		}

	} else {
		return domain.Connection{}, errors.New(string(resp.Body()))
	}
}
