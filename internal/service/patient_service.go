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

package service

import (
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/storage"
)

type PatientService struct {
	storage storage.PatientStorage
}

func (s *PatientService) GetDIDs(patientId string) (dids []string, err error) {
	return s.storage.GetDIDs(patientId)
}

func (s *PatientService) GetPrescriptionCredentials(patientId string) (credentials []domain.Credential, err error) {
	return s.storage.GetPrescriptionCredentials(patientId)
}

func (s *PatientService) AddPrescriptionCredential(patientId string, credential domain.Credential) (err error) {
	return s.storage.AddPrescriptionCredential(patientId, credential)
}

func (s *PatientService) GetPrescriptionCredentialById(patientId string, credentialId string) (credential domain.Credential, err error) {
	return s.storage.GetPrescriptionCredentialById(patientId, credentialId)
}
