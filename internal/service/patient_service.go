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
	"fmt"

	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/storage"
)

type PatientService struct {
	storage storage.PatientStorage
}

func NewPatientService(storage storage.PatientStorage) *PatientService {
	return &PatientService{
		storage: storage,
	}
}

func (s *PatientService) PatientExists(username string) bool {
	return s.storage.PatientExists(username)
}

func (s *PatientService) CreatePatient(username string, password string) (*domain.Patient, error) {
	if s.PatientExists(username) {
		return nil, fmt.Errorf("patient %s already exists", username)
	}
	return s.storage.CreatePatient(username, password)
}

func (s *PatientService) GetPatientByCredentials(username string, password string) (*domain.Patient, error) {
	return s.storage.GetPatientByCredentials(username, password)
}

func (s *PatientService) AddPatientDID(patientId string, did string) (err error) {
	return s.storage.AddPatientDID(patientId, did)
}

func (s *PatientService) GetDIDs(patientId string) (dids []string, err error) {
	return s.storage.GetDIDs(patientId)
}

func (s *PatientService) SaveCredentialId(patientId string, credentialId string) (err error) {
	return s.storage.AddCredentialIdByPatientId(patientId, credentialId)
}

func (s *PatientService) GetCredentialIdsByPatientId(patientId string) (credentialIds []string, err error) {
	return s.storage.GetCredentialIdsByPatientId(patientId)
}
