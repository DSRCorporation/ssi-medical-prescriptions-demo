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

package leveldb

import (
	"encoding/json"
	"fmt"
	"os"
)

type PatientStorage struct {
	levelDB  *LevelDB
	patients *patients
}

var patientsDIDsPath = "/etc/ssimp/testdata/patients.json"

type patient struct {
	PatientId string   `json:"patientId"`
	Dids      []string `json:"dids"`
}

type patients struct {
	Patients []patient `json:"patients"`
}

func NewPatientStorage(path string) (*PatientStorage, error) {
	levelDB, err := NewLevelDB(path)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(patientsDIDsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s file: %v", patientsDIDsPath, err)
	}

	var res patients
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshalling %s file: %v", patientsDIDsPath, err)
	}

	return &PatientStorage{levelDB: levelDB, patients: &res}, nil
}

func (s *PatientStorage) AddCredentialIdByPatientId(patientId string, credentialId string) (err error) {
	if patientId == "" {
		return fmt.Errorf("patientId cannot be empty")
	}

	exist, err := s.levelDB.Has(patientId)
	if err != nil {
		return err
	}

	var credentialIds []string

	if exist {
		if err = s.levelDB.ReadFromJson(patientId, &credentialIds); err != nil {
			return err
		}
	}

	credentialIds = append(credentialIds, credentialId)

	if err = s.levelDB.WriteAsJson(patientId, credentialIds); err != nil {
		return err
	}

	return nil
}

func (s *PatientStorage) GetCredentialIdsByPatientId(patientId string) (credentialIds []string, err error) {
	if err = s.levelDB.ReadFromJson(patientId, &credentialIds); err != nil {
		return credentialIds, err
	}

	return credentialIds, nil
}

func (s *PatientStorage) GetDIDs(patientId string) (dids []string, err error) {
	for _, patient := range s.patients.Patients {
		if patient.PatientId == patientId {
			return patient.Dids, nil
		}
	}

	return []string{}, fmt.Errorf("patient not found")
}
