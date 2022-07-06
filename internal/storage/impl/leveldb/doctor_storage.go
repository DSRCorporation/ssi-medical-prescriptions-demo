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
	"fmt"

	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

type DoctorStorage struct {
	levelDB *LevelDB
}

func NewDoctorStorage(dbPath string) (*DoctorStorage, error) {
	levelDB, err := NewLevelDB(dbPath)
	if err != nil {
		return nil, err
	}
	return &DoctorStorage{levelDB: levelDB}, nil
}

func (s *DoctorStorage) CreatePrescriptionOffer(offerId string, prescription domain.Prescription) (err error) {
	db, err := NewLevelDB(s.levelDB.path)
	if err != nil {
		return err
	}

	prescriptionOffer := &PrescriptionOffer{
		OfferId:      &offerId,
		Prescription: &prescription,
		CredentialId: nil,
	}

	exist, err := db.Has(offerId)
	if err != nil {
		return err
	}

	if exist {
		return fmt.Errorf("offerId already exists: %v", offerId)
	}

	if err = db.WriteAsJson(offerId, prescriptionOffer); err != nil {
		return err
	}

	return nil
}

func (s *DoctorStorage) GetPrescriptionByOfferId(offerId string) (prescription domain.Prescription, err error) {
	db, err := NewLevelDB(s.levelDB.path)
	if err != nil {
		return prescription, err
	}

	var prescriptionOffer PrescriptionOffer

	if err = db.ReadFromJson(offerId, &prescriptionOffer); err != nil {
		return prescription, err
	}

	prescription = *prescriptionOffer.Prescription

	return prescription, nil
}

func (s *DoctorStorage) AddCredentialIdByOfferId(offerId string, credentialId string) (err error) {
	db, err := NewLevelDB(s.levelDB.path)
	if err != nil {
		return err
	}

	var prescriptionOffer PrescriptionOffer

	if err = db.ReadFromJson(offerId, &prescriptionOffer); err != nil {
		return err
	}

	if prescriptionOffer.CredentialId != nil {
		return fmt.Errorf("offerId already exists: %v", offerId)
	}

	prescriptionOffer.CredentialId = &credentialId

	if err = db.WriteAsJson(offerId, prescriptionOffer); err != nil {
		return err
	}

	return nil
}

func (s *DoctorStorage) GetCredentialIdByOfferId(offerId string) (credentialId string, err error) {
	db, err := NewLevelDB(s.levelDB.path)
	if err != nil {
		return credentialId, err
	}

	var prescriptionOffer PrescriptionOffer

	if err = db.ReadFromJson(offerId, &prescriptionOffer); err != nil {
		return credentialId, err
	}

	if prescriptionOffer.CredentialId == nil {
		return credentialId, fmt.Errorf("no credential id for offer id: %s", offerId)
	}

	credentialId = *prescriptionOffer.CredentialId

	return credentialId, nil
}

func (s *DoctorStorage) AddCredentialIdByDoctorId(doctorId string, credentials string) (err error) {
	db, err := NewLevelDB(s.levelDB.path)
	if err != nil {
		return err
	}

	var credentialIds []string

	exist, err := db.Has(doctorId)
	if err != nil {
		return err
	}

	if exist {
		if err = db.ReadFromJson(doctorId, &credentialIds); err != nil {
			return err
		}
	}

	credentialIds = append(credentialIds, credentials)

	if err = db.WriteAsJson(doctorId, credentialIds); err != nil {
		return err
	}

	return nil
}

func (s *DoctorStorage) GetCredentialIdsByDoctorId(doctorId string) (credentialIds []string, err error) {
	db, err := NewLevelDB(s.levelDB.path)
	if err != nil {
		return credentialIds, err
	}

	if err = db.ReadFromJson(doctorId, &credentialIds); err != nil {
		return credentialIds, err
	}

	return credentialIds, nil
}

func (s *DoctorStorage) GetKMSPassphrase(doctorId string) (kmspassphrase string, err error) {
	return kmspassphrase, fmt.Errorf("not implemented!")
}

func (s *DoctorStorage) GetDID(doctorId string) (did string, err error) {
	return did, fmt.Errorf("not implemented!")
}

type PrescriptionOffer struct {
	OfferId      *string
	Prescription *domain.Prescription
	CredentialId *string
}
