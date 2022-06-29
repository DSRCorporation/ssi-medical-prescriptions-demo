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

import (
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

type DoctorStorage interface {
	CreatePrescriptionOffer(offerId string, prescription domain.Prescription) (err error)
	GetPrescriptionByOfferId(offerId string) (prescription domain.Prescription, err error)
	GetCredentialIdByOfferId(offerId string) (credentialId string, err error)
	GetCredentialIdsByDoctorId(doctorId string) (credentialIds []string, err error)
	AddCredentialIdByDoctorId(doctorId string, credentialId string) (err error)
	AddCredentialIdByOfferId(offerId string, credentialId string) (err error)
	GetKMSPassphrase(doctorId string) (kmspassphrase string, err error)
	GetDID(doctorId string) (did string, err error)
}
