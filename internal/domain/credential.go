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

package domain

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Credential struct {
	CredentialId string
	IssuerDID    string
	HolderDID    string
	Prescription Prescription

	RawCredentialWithProof json.RawMessage
}

func NewCredential(issuerDID string, holderDID string, prescription Prescription) (credential *Credential, err error) {
	return &Credential{
		CredentialId: generateCredentialId(),
		IssuerDID:    issuerDID,
		HolderDID:    holderDID,
		Prescription: prescription,
	}, nil
}

func generateCredentialId() string {
	credentialId := fmt.Sprintf("did:%s", uuid.New().String())
	return credentialId
}
