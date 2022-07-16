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
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/hyperledger/aries-framework-go/pkg/client/issuecredential"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service"
	"github.com/hyperledger/aries-framework-go/pkg/doc/util"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/piprate/json-gold/ld"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

const (
	maxAttempts = 3
)

var prescriptionContext = []string{"https://www.w3.org/2018/credentials/v1", "https://ssimp.s3.amazonaws.com/schemas/prescription"}

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

	req := client.R().
		SetPathParam("state", "requested").
		SetPathParam("invitationId", i.InvitationId).
		SetResult(&c)

	cond := func() bool {
		return len(c.Results) > 0
	}

	err = executeRequestUntil(req, resty.MethodGet, http.StatusOK, "/connections?state={state}&invitation_id={invitationId}", cond)
	if err != nil {
		return domain.Connection{}, err
	}

	connectionId := c.Results[0].ConnectionId

	resp, err := client.R().
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
				State        string `json:"State"`
			} `json:"result"`
		}

		req := client.R().
			SetPathParam("connectionId", connectionId).
			SetResult(&res)

		cond := func() bool {
			return res.Result.State == "completed"
		}

		err := executeRequestUntil(req, resty.MethodGet, http.StatusOK, "/connections/{connectionId}", cond)
		if err != nil {
			return domain.Connection{}, fmt.Errorf("connection could not be established: %v", err)
		}

		conn := domain.Connection{
			InviterDID:   res.Result.MyDID,
			InviteeDID:   res.Result.TheirDID,
			ConnectionId: res.Result.ConnectionID,
		}

		return conn, nil
	} else {
		return domain.Connection{}, errors.New(string(resp.Body()))
	}
}

func getCredentialFromActions(client *resty.Client, piid string) (*domain.Credential, error) {
	rawCredential, err := getRawJsonFromAttachmentData(client, "/issuecredential/actions", piid)
	if err != nil {
		return nil, err
	}

	return toCredential(rawCredential)
}

func getPresentationFromActions(client *resty.Client, piid string) (*domain.Presentation, error) {
	rawPresentation, err := getRawJsonFromAttachmentData(client, "/presentproof/actions", piid)
	if err != nil {
		return nil, err
	}

	return toPresentation(rawPresentation)
}

func getAttachmentFromActions(client *resty.Client, endpoint string, piid string) (attachment interface{}, err error) {
	var res struct {
		Actions []issuecredential.Action `json:"actions"`
	}

	req := client.R().
		SetResult(&res)

	cond := func() bool {
		for _, action := range res.Actions {
			if action.PIID == piid {
				return true
			}
		}
		return false
	}

	err = executeRequestUntil(req, resty.MethodGet, http.StatusOK, endpoint, cond)
	if err != nil {
		return nil, err
	}

	for _, action := range res.Actions {
		if action.PIID == piid {
			return getAttachmentFromActionMsg(action.Msg)
		}
	}

	return nil, errors.New("action with given piid not found")
}

func getAttachmentFromActionMsg(msg service.DIDCommMsgMap) (interface{}, error) {
	for _, key := range []string{"offers~attach", "requests~attach", "credentials~attach", "presentations~attach"} {
		if val, ok := msg[key]; ok {
			attachments := val.([]interface{})
			if len(attachments) > 0 {
				return attachments[0], nil
			}
		}
	}
	return nil, errors.New("no attachments found")
}

func getRawJsonFromAttachmentData(client *resty.Client, endpoint string, piid string) (json.RawMessage, error) {
	attachment, err := getAttachmentFromActions(client, endpoint, piid)
	if err != nil {
		return nil, err
	}

	raw, err := json.Marshal(attachment.(map[string]interface{})["data"].(map[string]interface{})["json"])
	if err == nil {
		return json.RawMessage(raw), nil
	}
	return nil, err
}

func toRawCredential(credential domain.Credential) (json.RawMessage, error) {
	var cred verifiable.Credential

	cred.ID = credential.CredentialId
	cred.Issuer = verifiable.Issuer{
		ID: credential.IssuerDID,
	}
	cred.Issued = util.NewTime(time.Now().UTC())
	cred.Context = prescriptionContext
	cred.Types = []string{"VerifiableCredential"}

	var prescription map[string]interface{}
	err := json.Unmarshal(credential.Prescription.RawPrescription, &prescription)
	if err != nil {
		return nil, err
	}

	cred.Subject = verifiable.Subject{
		ID: credential.HolderDID,
		CustomFields: verifiable.CustomFields{
			"prescription": prescription,
		},
	}

	bytes, err := cred.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return json.RawMessage(bytes), nil
}

func toRawPresentation(presentation domain.Presentation) (rawPresentation json.RawMessage, err error) {
	cred, err := verifiable.ParseCredential([]byte(presentation.Credential.RawCredential), verifiable.WithBaseContextExtendedValidation(prescriptionContext, []string{}), verifiable.WithDisabledProofCheck())
	if err != nil {
		return nil, err
	}

	pres, err := verifiable.NewPresentation(verifiable.WithCredentials(cred))
	if err != nil {
		return nil, err
	}

	pres.ID = presentation.PresentationId
	pres.Holder = presentation.HolderDID

	bytes, err := pres.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return json.RawMessage(bytes), nil
}

func toCredential(rawCredential json.RawMessage) (*domain.Credential, error) {
	cred, err := verifiable.ParseCredential([]byte(rawCredential), verifiable.WithBaseContextExtendedValidation(prescriptionContext, []string{}), verifiable.WithDisabledProofCheck())
	if err != nil {
		return nil, err
	}

	subjectID, err := verifiable.SubjectID(cred.Subject)
	if err != nil {
		return nil, err
	}

	var rawPrescription json.RawMessage
	subjects := cred.Subject.([]verifiable.Subject)
	if len(subjects) > 0 {
		if prescription, ok := subjects[0].CustomFields["prescription"]; ok {
			rawPrescription, err = json.Marshal(prescription)
			if err != nil {
				return nil, err
			}
		}
	}

	credential := &domain.Credential{
		CredentialId: cred.ID,
		IssuerDID:    cred.Issuer.ID,
		HolderDID:    subjectID,
		Prescription: domain.Prescription{
			RawPrescription: rawPrescription,
		},
		RawCredential: rawCredential,
	}

	return credential, err
}

func toPresentation(rawPresentation json.RawMessage) (*domain.Presentation, error) {
	pres, err := verifiable.ParsePresentation([]byte(rawPresentation), verifiable.WithPresDisabledProofCheck(), verifiable.WithPresJSONLDDocumentLoader(ld.NewDefaultDocumentLoader(nil)))
	if err != nil {
		return nil, err
	}

	credentials := pres.Credentials()

	if len(credentials) == 0 {
		return nil, fmt.Errorf("no credentials found")
	}

	rawCredential, err := json.Marshal(credentials[0])
	if err != nil {
		return nil, err
	}

	credential, err := toCredential(rawCredential)
	if err != nil {
		return nil, err
	}

	presentation := domain.NewPresentation(pres.ID, pres.Holder, *credential, rawPresentation)

	return presentation, nil
}

func executeRequestUntil(request *resty.Request, method string, statusCode int, endpoint string, condition func() bool) error {
	attempts := 0
	for attempts < maxAttempts {
		resp, err := request.Execute(method, endpoint)
		if err != nil {
			return err
		}

		if resp.StatusCode() != statusCode {
			return fmt.Errorf(string(resp.Body()))
		}

		if condition() == true {
			return nil
		}

		time.Sleep(time.Millisecond * 100)
	}
	return fmt.Errorf("http requet condition is not fulfilled within given attempts")
}
