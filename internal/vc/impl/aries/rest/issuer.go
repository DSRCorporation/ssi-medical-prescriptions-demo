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

	client "github.com/hyperledger/aries-framework-go/pkg/client/issuecredential"
	issuecredentialcmd "github.com/hyperledger/aries-framework-go/pkg/controller/command/issuecredential"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/decorator"
)

type Issuer struct {
	client *resty.Client
}

func NewIssuer(endpoint string) (*Issuer, error) {
	client := resty.New()
	client.SetBaseURL(endpoint)
	return &Issuer{client: client}, nil
}

func (i *Issuer) SendCredentialOffer(connection domain.Connection, credential domain.Credential) (piid string, err error) {

	rawCredential, err := makeRawCredential(credential)
	if err != nil {
		return "", err
	}

	offerCredential := client.OfferCredentialV2{
		OffersAttach: []decorator.Attachment{{
			Data: decorator.AttachmentData{
				JSON: rawCredential,
			},
		}},
	}

	var res map[string]interface{}
	resp, err := i.client.R().
		SetBody(issuecredentialcmd.SendOfferArgsV2{
			MyDID:           connection.InviterDID,
			TheirDID:        connection.InviteeDID,
			OfferCredential: &offerCredential,
		}).
		SetResult(&res).
		Post("/issuecredential/send-offer")

	if err != nil {
		return "", err
	}

	if resp.StatusCode() == http.StatusOK {
		return res["piid"].(string), nil
	} else {
		return "", errors.New(string(resp.Body()))
	}
}

func (i *Issuer) GetCredentialFromRequest(piid string) (credential *domain.Credential, err error) {
	return getCredentialFromActions(i.client, piid)
}

func (i *Issuer) AcceptCredentialRequest(piid string, credential domain.Credential) (err error) {
	if credential.RawCredentialWithProof == nil {
		return errors.New("raw credential cannot be nil")
	}

	issueCredential := client.IssueCredentialV2{
		CredentialsAttach: []decorator.Attachment{{
			Data: decorator.AttachmentData{
				JSON: credential.RawCredentialWithProof,
			},
		}},
	}

	resp, err := i.client.R().
		SetBody(issuecredentialcmd.AcceptRequestArgsV2{
			IssueCredential: &issueCredential,
		}).
		SetPathParam("piid", piid).
		Post("/issuecredential/{piid}/accept-request")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return errors.New(string(resp.Body()))
	}
}

func (i *Issuer) CreateOOBInvitation() (invitation json.RawMessage, err error) {
	return CreateOOBInvitation(i.client)
}

func (i *Issuer) AcceptOOBRequest(invitation json.RawMessage) (connection domain.Connection, err error) {
	return AcceptOOBRequest(i.client, invitation)
}
