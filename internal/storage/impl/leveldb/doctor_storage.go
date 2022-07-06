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
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

type DoctorStorage struct {
	LeveldbPath string
}

func NewDoctorStorage(dbPath string) (*DoctorStorage, error) {
	return &DoctorStorage{LeveldbPath: dbPath}, nil
}

func (s *DoctorStorage) CreatePrescriptionOffer(offerId string, prescription domain.Prescription) (err error) {
	db, err := leveldb.OpenFile(s.LeveldbPath, nil)
	if err != nil {
		return fmt.Errorf("error opening database file: %v", err)
	}
	defer db.Close()

	exist, _ := db.Has([]byte(offerId), nil)
	if exist {
		return fmt.Errorf("offer already exists: %v", offerId)
	}

	data := &Prescription{
		OfferId:      &offerId,
		Prescription: &prescription,
		CredentialId: nil,
	}

	value, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error converting type Prescription to json bytes: %v", err)
	}

	if err = db.Put([]byte(offerId), value, nil); err != nil {
		return fmt.Errorf("error writing to database: %v", err)
	}

	return nil
}

func (s *DoctorStorage) GetPrescriptionByOfferId(offerId string) (prescription domain.Prescription, err error) {
	db, err := leveldb.OpenFile(s.LeveldbPath, nil)
	if err != nil {
		return prescription, fmt.Errorf("error opening database file: %v", err)
	}
	defer db.Close()

	value, err := db.Get([]byte(offerId), nil)
	if err != nil {
		return prescription, fmt.Errorf("error reading from database: %v", err)
	}

	var data Prescription

	if err = json.Unmarshal(value, &data); err != nil {
		return prescription, fmt.Errorf("error unmarshalling data: %v", err)
	}

	prescription = *data.Prescription

	return prescription, nil
}

func (s *DoctorStorage) AddCredentialIdByOfferId(offerId string, credentialId string) (err error) {
	db, err := leveldb.OpenFile(s.LeveldbPath, nil)
	if err != nil {
		return fmt.Errorf("error opening database file: %v", err)
	}
	defer db.Close()

	value, err := db.Get([]byte(offerId), nil)
	if err != nil {
		return fmt.Errorf("error not found prescription offer: %v", err)
	}

	var prescription Prescription

	if err = json.Unmarshal(value, &prescription); err != nil {
		return fmt.Errorf("error unmarshalling data: %v", err)
	}

	if prescription.CredentialId != nil {
		return fmt.Errorf("offerId already exists: %v", offerId)
	}

	prescription.CredentialId = &credentialId

	value, err = json.Marshal(prescription)
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}

	if err = db.Put([]byte(offerId), value, nil); err != nil {
		return fmt.Errorf("error writing to database: %v", err)
	}

	return nil
}

func (s *DoctorStorage) GetCredentialIdByOfferId(offerId string) (credentialId string, err error) {
	db, err := leveldb.OpenFile(s.LeveldbPath, nil)
	if err != nil {
		return credentialId, fmt.Errorf("error opening database file: %v", err)
	}
	defer db.Close()

	value, err := db.Get([]byte(offerId), nil)
	if err != nil {
		return credentialId, fmt.Errorf("error reading from database: %v", err)
	}

	var prescription Prescription

	if err := json.Unmarshal(value, &prescription); err != nil {
		return credentialId, fmt.Errorf("error unmarshalling data: %v", err)
	}

	if prescription.CredentialId == nil {
		return credentialId, fmt.Errorf("no credential id for offer id: %s", offerId)
	}

	return *prescription.CredentialId, nil
}

func (s *DoctorStorage) AddCredentialIdByDoctorId(doctorId string, credentials string) (err error) {
	db, err := leveldb.OpenFile(s.LeveldbPath, nil)
	if err != nil {
		return fmt.Errorf("error opening database file: %v", err)
	}
	defer db.Close()

	var credentialIds []string

	exist, err := db.Has([]byte(doctorId), nil)
	if err != nil {
		return fmt.Errorf("error reading from database: %v", err)
	}

	if exist {
		value, err := db.Get([]byte(doctorId), nil)
		if err != nil {
			return fmt.Errorf("error reading from database: %v", err)
		}

		if err = json.Unmarshal(value, &credentialIds); err != nil {
			return fmt.Errorf("error unmarshalling credentials ids: %v", err)
		}
	}

	credentialIds = append(credentialIds, credentials)

	value, err := json.Marshal(credentialIds)
	if err != nil {
		return fmt.Errorf("error marshalling credential ids: %v", err)
	}

	if err = db.Put([]byte(doctorId), value, nil); err != nil {
		return fmt.Errorf("error writing to database: %v", err)
	}

	return nil
}

func (s *DoctorStorage) GetCredentialIdsByDoctorId(doctorId string) (credentialIds []string, err error) {
	db, err := leveldb.OpenFile(s.LeveldbPath, nil)
	if err != nil {
		return credentialIds, fmt.Errorf("error opening database file: %v", err)
	}
	defer db.Close()

	value, err := db.Get([]byte(doctorId), nil)
	if err != nil {
		return credentialIds, fmt.Errorf("error reading from database: %v", err)
	}

	if err = json.Unmarshal(value, &credentialIds); err != nil {
		return credentialIds, fmt.Errorf("error unmarshalling credential ids: %v", err)
	}

	return credentialIds, nil
}

func (s *DoctorStorage) GetKMSPassphrase(doctorId string) (kmspassphrase string, err error) {
	return kmspassphrase, fmt.Errorf("not implemented!")
}

func (s *DoctorStorage) GetDID(doctorId string) (did string, err error) {
	return did, fmt.Errorf("not implemented!")
}

type Prescription struct {
	OfferId      *string
	Prescription *domain.Prescription
	CredentialId *string
}
