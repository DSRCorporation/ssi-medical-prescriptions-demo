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

package aries

import (
	"encoding/json"

	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

type Issuer struct{}

func (i *Issuer) SendCredentialOffer(connection domain.Connection, credential domain.Credential) (piid string, err error) {
	return "", nil
}

func (i *Issuer) AcceptCredentialRequest(piid string, credential domain.Credential) error {
	return nil
}

func (i *Issuer) CreateOOBInvitation() (invitation json.RawMessage, err error) {
	return nil, nil
}

func (i *Issuer) AcceptOOBRequest(connectionId string) (connection domain.Connection, err error) {
	return domain.Connection{}, nil
}
