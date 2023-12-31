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

package storage

import "github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"

type VCStorage interface {
	GetConnection(inviterId string, inviteeId string) (connection domain.Connection, err error)
	SaveConnection(inviterId string, inviteeId string, connection domain.Connection) (err error)
	GetCredentialById(credentialId string) (credential domain.Credential, err error)
	GetPresentationById(presentationId string) (presentation domain.Presentation, err error)
	SaveCredential(credential domain.Credential) error
	SavePresentation(presentation domain.Presentation) (err error)

	// Gets wallet credentials used internally for credential/presentation verification using wallet interface
	GetVerifierWalletCredentials() (userId string, passphrase string, err error)
}
