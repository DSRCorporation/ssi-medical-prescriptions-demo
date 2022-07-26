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

	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

type DoctorStorage struct {
	levelDB *LevelDB
	doctors []byte
}

func NewDoctorStorage(dbPath string) (*DoctorStorage, error) {
	levelDB, err := NewLevelDB(dbPath)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile("etc/ssimp/testdata/doctors.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read etc/ssimp/testdata/doctors.json file: %v", err)
	}

	return &DoctorStorage{levelDB: levelDB, doctors: data}, nil
}

func (s *DoctorStorage) CreatePrescriptionOffer(offerId string, prescription domain.Prescription) (err error) {
	if offerId == "" {
		return fmt.Errorf("offerId cannot be empty")
	}

	exist, err := s.levelDB.Has(offerId)
	if err != nil {
		return err
	}

	if exist {
		return fmt.Errorf("offerId already exists: %v", offerId)
	}

	prescriptionOffer := &PrescriptionOffer{
		OfferId:      &offerId,
		Prescription: &prescription,
		CredentialId: nil,
	}

	if err = s.levelDB.WriteAsJson(offerId, prescriptionOffer); err != nil {
		return err
	}

	return nil
}

func (s *DoctorStorage) GetPrescriptionByOfferId(offerId string) (prescription domain.Prescription, err error) {
	var prescriptionOffer PrescriptionOffer

	if err = s.levelDB.ReadFromJson(offerId, &prescriptionOffer); err != nil {
		return prescription, err
	}

	prescription = *prescriptionOffer.Prescription

	return prescription, nil
}

func (s *DoctorStorage) AddCredentialIdByOfferId(offerId string, credentialId string) (err error) {
	if offerId == "" {
		return fmt.Errorf("offerId cannot be empty")
	}

	var prescriptionOffer PrescriptionOffer

	if err = s.levelDB.ReadFromJson(offerId, &prescriptionOffer); err != nil {
		return err
	}

	if prescriptionOffer.CredentialId != nil {
		return fmt.Errorf("offerId already exists: %v", offerId)
	}

	prescriptionOffer.CredentialId = &credentialId

	if err = s.levelDB.WriteAsJson(offerId, prescriptionOffer); err != nil {
		return err
	}

	return nil
}

func (s *DoctorStorage) GetCredentialIdByOfferId(offerId string) (credentialId string, err error) {
	var prescriptionOffer PrescriptionOffer

	if err = s.levelDB.ReadFromJson(offerId, &prescriptionOffer); err != nil {
		return credentialId, err
	}

	if prescriptionOffer.CredentialId == nil {
		return credentialId, fmt.Errorf("no credential id for offer id: %s", offerId)
	}

	credentialId = *prescriptionOffer.CredentialId

	return credentialId, nil
}

func (s *DoctorStorage) AddCredentialIdByDoctorId(doctorId string, credentials string) (err error) {
	if doctorId == "" {
		return fmt.Errorf("doctor id cannot be empty")
	}

	exist, err := s.levelDB.Has(doctorId)
	if err != nil {
		return err
	}

	var credentialIds []string

	if exist {
		if err = s.levelDB.ReadFromJson(doctorId, &credentialIds); err != nil {
			return err
		}
	}

	credentialIds = append(credentialIds, credentials)

	if err = s.levelDB.WriteAsJson(doctorId, credentialIds); err != nil {
		return err
	}

	return nil
}

func (s *DoctorStorage) GetCredentialIdsByDoctorId(doctorId string) (credentialIds []string, err error) {
	if err = s.levelDB.ReadFromJson(doctorId, &credentialIds); err != nil {
		return credentialIds, err
	}

	return credentialIds, nil
}

func (s *DoctorStorage) GetKMSPassphrase(doctorId string) (kmspassphrase string, err error) {
	return "Np6VR4Yg6PPL", nil
}

func (s *DoctorStorage) GetDID(doctorId string) (did string, err error) {
	type doctor struct {
		DoctorId string   `json:"doctorId"`
		Dids     []string `json:"dids"`
	}

	type doctors struct {
		Doctors []doctor `json:"doctors"`
	}

	var res doctors
	if err = json.Unmarshal(s.doctors, &res); err != nil {
		return "", fmt.Errorf("failed to unmarshalling etc/ssimp/testdata/doctors.json file: %v", err)
	}

	for _, doctor := range res.Doctors {
		if doctor.DoctorId == doctorId {
			if len(doctor.Dids) > 0 {
				return doctor.Dids[0], nil
			} else {
				return "", fmt.Errorf("doctor did not found")
			}
		}
	}

	return "", fmt.Errorf("doctor not found")
}

type PrescriptionOffer struct {
	OfferId      *string
	Prescription *domain.Prescription
	CredentialId *string
}
