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
	client "github.com/hyperledger/aries-framework-go/pkg/client/issuecredential"
	"github.com/hyperledger/aries-framework-go/pkg/client/presentproof"
	issuecredentialcmd "github.com/hyperledger/aries-framework-go/pkg/controller/command/issuecredential"
	presentproofcmd "github.com/hyperledger/aries-framework-go/pkg/controller/command/presentproof"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/decorator"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

type Holder struct {
	client *resty.Client
}

func NewHolder(endpoint string) (*Holder, error) {
	client := resty.New()
	client.SetBaseURL(endpoint)
	return &Holder{client: client}, nil
}

func (h *Holder) GetCredentialFromOffer(piid string) (credential *domain.Credential, err error) {
	return getCredentialFromActions(h.client, piid)
}

func (h *Holder) GetIssuedCredential(piid string) (credential *domain.Credential, err error) {
	return getCredentialFromActions(h.client, piid)
}

func (h *Holder) SendCredentialRequest(connection domain.Connection, credential domain.Credential) (piid string, err error) {

	if credential.RawCredential == nil {
		return "", errors.New("raw credential cannot be nil")
	}

	requestCredential := client.RequestCredentialV2{
		RequestsAttach: []decorator.Attachment{{
			Data: decorator.AttachmentData{
				JSON: credential.RawCredential,
			},
		}},
	}

	var res map[string]interface{}
	resp, err := h.client.R().
		SetBody(issuecredentialcmd.SendRequestArgsV2{
			MyDID:             connection.InviteeDID,
			TheirDID:          connection.InviterDID,
			RequestCredential: &requestCredential,
		}).
		SetResult(&res).
		Post("/issuecredential/send-request")

	if err != nil {
		return "", err
	}

	if resp.StatusCode() == http.StatusOK {
		return res["piid"].(string), nil
	} else {
		return "", errors.New(string(resp.Body()))
	}
}

func (h *Holder) AcceptOffer(piid string) error {
	resp, err := h.client.R().
		SetPathParam("piid", piid).
		Post("/issuecredential/{piid}/accept-offer")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return errors.New(string(resp.Body()))
	}
}

func (h *Holder) AcceptCredential(piid string, name string) error {
	resp, err := h.client.R().
		SetPathParam("piid", piid).
		SetBody(issuecredentialcmd.AcceptCredentialArgs{
			Names: []string{name},
		}).
		Post("/issuecredential/{piid}/accept-credential")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return errors.New(string(resp.Body()))
	}
}

func (h *Holder) AcceptPresentationRequest(piid string, presentation domain.Presentation) (err error) {
	presentationV2 := presentproof.PresentationV2{
		PresentationsAttach: []decorator.Attachment{{
			Data: decorator.AttachmentData{
				JSON: presentation.RawPresentationWithProof,
			},
		}},
	}

	resp, err := h.client.R().
		SetPathParam("piid", piid).
		SetBody(presentproofcmd.AcceptRequestPresentationV2Args{
			Presentation: &presentationV2,
		}).
		Post("/presentproof/{piid}/accept-request-presentation")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return errors.New(string(resp.Body()))
	}
}

func (h *Holder) AcceptOOBInvitation(invitation json.RawMessage) (err error) {
	resp, err := h.client.R().
		SetBody(struct {
			Invitation json.RawMessage `json:"invitation"`
			MyLabel    string          `json:"my_label"`
		}{
			Invitation: invitation,
			MyLabel:    "Holder",
		}).
		Post("/outofband/accept-invitation")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return errors.New(string(resp.Body()))
	}
}
