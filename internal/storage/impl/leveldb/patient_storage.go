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

	"github.com/syndtr/goleveldb/leveldb"
)

type PatientStorage struct {
	LeveldbPath string
}

func NewPatientStorage(path string) (*PatientStorage, error) {
	return &PatientStorage{LeveldbPath: path}, nil
}

func (s *PatientStorage) AddCredentialIdByPatientId(patientId string, credentialId string) (err error) {
	db, err := leveldb.OpenFile(s.LeveldbPath, nil)
	if err != nil {
		return fmt.Errorf("error opening database file: %v", err)
	}
	defer db.Close()

	var credentialIds []string

	exist, err := db.Has([]byte(patientId), nil)
	if err != nil {
		return fmt.Errorf("error reading from database: %v", err)
	}

	if exist {
		value, err := db.Get([]byte(patientId), nil)
		if err != nil {
			return fmt.Errorf("error reading from database: %v", err)
		}

		if err = json.Unmarshal(value, &credentialIds); err != nil {
			return fmt.Errorf("error unmarshalling credentialIds: %v", err)
		}
	}

	credentialIds = append(credentialIds, credentialId)

	value, err := json.Marshal(credentialIds)
	if err != nil {
		return fmt.Errorf("error marshalling credential ids: %v", err)
	}

	if err = db.Put([]byte(patientId), value, nil); err != nil {
		return fmt.Errorf("error writing to database: %v", err)
	}

	return nil
}

func (s *PatientStorage) GetCredentialIdsByPatientId(patientId string) (credentialIds []string, err error) {
	db, err := leveldb.OpenFile(s.LeveldbPath, nil)
	if err != nil {
		return credentialIds, fmt.Errorf("error opening database file: %v", err)
	}
	defer db.Close()

	value, err := db.Get([]byte(patientId), nil)
	if err != nil {
		return credentialIds, fmt.Errorf("error reading from database: %v", err)
	}

	if err = json.Unmarshal(value, &credentialIds); err != nil {
		return credentialIds, fmt.Errorf("error unmarshalling credential ids: %v", err)
	}

	return credentialIds, nil
}

func (s *PatientStorage) GetDIDs(patientId string) (dids []string, err error) {
	return dids, fmt.Errorf("not implemented!")
}
