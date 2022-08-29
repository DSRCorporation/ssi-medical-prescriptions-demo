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
	"crypto/ed25519"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	vdrCmd "github.com/hyperledger/aries-framework-go/pkg/controller/command/vdr"
	"github.com/hyperledger/aries-framework-go/pkg/doc/did"
	didfp "github.com/hyperledger/aries-framework-go/pkg/vdr/fingerprint"
	"github.com/hyperledger/aries-framework-go/pkg/vdr/key"
)

type VDR struct {
	client *resty.Client
	vdr    *key.VDR
}

func NewVDR(endpoint string) (*VDR, error) {
	client := resty.New()
	vdr := key.New()
	client.SetBaseURL(endpoint)

	return &VDR{client: client, vdr: vdr}, nil
}

func (v *VDR) CreateKeyDID(pubKey ed25519.PublicKey) (id string, verificationMethodId string, err error) {
	didKey, _ := didfp.CreateDIDKey(pubKey)

	docResolution, err := v.vdr.Read(didKey)
	if err != nil {
		return "", "", err
	}

	rawDidDoc, err := docResolution.DIDDocument.JSONBytes()
	if err != nil {
		return "", "", err
	}

	request := vdrCmd.CreateDIDRequest{
		Method: "key",
		DID:    rawDidDoc,
	}

	var res vdrCmd.Document

	resp, err := v.client.R().
		SetBody(request).
		SetResult(&res).
		Post("/vdr/did/create")

	if err != nil {
		return "", "", err
	}

	if resp.StatusCode() == http.StatusOK {
		doc, err := did.ParseDocument(res.DID)
		if err != nil {
			return "", "", err
		}

		if len(doc.VerificationMethod) > 0 {
			return doc.ID, doc.VerificationMethod[0].ID, nil
		}

		return "", "", fmt.Errorf("no verification methods found")
	}

	return "", "", fmt.Errorf(resp.String())
}
