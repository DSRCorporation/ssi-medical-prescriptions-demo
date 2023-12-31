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

package vc

import (
	"crypto/ed25519"
	"encoding/json"

	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

type Wallet interface {
	SignCredential(userId string, passphrase string, did string, credential domain.Credential) (domain.Credential, error)
	SignPresentation(userId string, passphrase string, did string, presentaion domain.Presentation) (domain.Presentation, error)
	VerifyCredential(userId string, passphrase string, rawCredential json.RawMessage) error
	VerifyPresentation(userId string, passphrase string, rawPresentation json.RawMessage) error

	WalletExists(userId string) bool
	CreateWallet(userId string, passphrase string) (err error)

	AddKey(userId string, passphrase string, keyId string, privKey ed25519.PrivateKey) (err error)
}
