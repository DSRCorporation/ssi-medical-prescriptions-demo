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

func CreateOOBInvitation(client *resty.Client) (invitation json.RawMessage, err error) {
	var res struct {
		Invitation json.RawMessage `json:"invitation"`
	}

	resp, err := client.R().
		SetResult(&res).
		SetBody(struct {
			Label string `json:"label"`
		}{Label: "Issuer"}).
		Post("/outofband/create-invitation")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return res.Invitation, nil
	} else {
		return nil, errors.New(string(resp.Body()))
	}
}

func AcceptOOBRequest(client *resty.Client, invitation json.RawMessage) (connection domain.Connection, err error) {
	var i struct {
		InvitationId string `json:"@id"`
	}

	err = json.Unmarshal(invitation, &i)
	if err != nil {
		return domain.Connection{}, err
	}

	var c struct {
		Results []struct {
			ConnectionId string `json:"ConnectionID"`
		} `json:"results"`
	}

	var attempts int = 0
	var resp *resty.Response
	for attempts <= 3 {
		resp, err = client.R().
			SetPathParam("state", "requested").
			SetPathParam("invitationId", i.InvitationId).
			SetResult(&c).
			Get("/connections?state={state}&invitation_id={invitationId}")

		if len(c.Results) > 0 {
			break
		}

		time.Sleep(200 * time.Millisecond)
	}

	if err != nil {
		return domain.Connection{}, err
	}

	if resp.StatusCode() == http.StatusOK && len(c.Results) > 0 {
		connectionId := c.Results[0].ConnectionId
		resp, err = client.R().
			SetPathParam("connectionId", connectionId).
			Post("/connections/{connectionId}/accept-request")

		if err != nil {
			return domain.Connection{}, err
		}

		if resp.StatusCode() == http.StatusOK {
			var res struct {
				Result struct {
					ConnectionID string `json:"ConnectionID"`
					MyDID        string `json:"MyDID"`
					TheirDID     string `json:"TheirDID"`
				} `json:"result"`
			}
			resp, err := client.R().
				SetPathParam("connectionId", connectionId).
				SetResult(&res).
				Get("/connections/{connectionId}")

			if err != nil {
				return domain.Connection{}, err
			}

			if resp.StatusCode() == http.StatusOK {
				conn := domain.Connection{
					InviterDID:   res.Result.MyDID,
					InviteeDID:   res.Result.TheirDID,
					ConnectionId: res.Result.ConnectionID,
				}

				return conn, nil
			} else {
				return domain.Connection{}, errors.New(string(resp.Body()))
			}

		} else {
			return domain.Connection{}, errors.New(string(resp.Body()))
		}
	} else {
		return domain.Connection{}, errors.New(string(resp.Body()))
	}
}
