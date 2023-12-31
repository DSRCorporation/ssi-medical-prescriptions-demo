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
	"github.com/hyperledger/aries-framework-go/pkg/client/presentproof"
	presentproofcmd "github.com/hyperledger/aries-framework-go/pkg/controller/command/presentproof"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

type Verifier struct {
	client *resty.Client
}

func NewVerifier(endpoint string) (*Verifier, error) {
	client := resty.New()
	client.SetBaseURL(endpoint)
	return &Verifier{client: client}, nil
}

func (v *Verifier) SendPresentationRequest(connection domain.Connection) (piid string, err error) {
	var res map[string]interface{}
	resp, err := v.client.R().
		SetBody(presentproofcmd.SendRequestPresentationV2Args{
			MyDID:               connection.InviterDID,
			TheirDID:            connection.InviteeDID,
			RequestPresentation: &presentproof.RequestPresentationV2{},
		}).
		SetResult(&res).
		Post("/presentproof/send-request-presentation")

	if err != nil {
		return "", err
	}

	if resp.StatusCode() == http.StatusOK {
		return res["piid"].(string), nil
	} else {
		return "", errors.New(string(resp.Body()))
	}
}

func (v *Verifier) GetIssuedPresentation(piid string) (presentation *domain.Presentation, err error) {
	return getPresentationFromActions(v.client, piid)
}

func (v *Verifier) AcceptPresentation(piid string, name string) error {
	resp, err := v.client.R().
		SetPathParam("piid", piid).
		SetBody(presentproofcmd.AcceptPresentationArgs{
			Names: []string{name},
		}).
		Post("/presentproof/{piid}/accept-presentation")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return errors.New(string(resp.Body()))
	}
}

func (v *Verifier) CreateOOBInvitation() (invitation json.RawMessage, err error) {
	return CreateOOBInvitation(v.client)
}

func (v *Verifier) AcceptOOBRequest(invitation json.RawMessage) (connection domain.Connection, err error) {
	return AcceptOOBRequest(v.client, invitation)
}
