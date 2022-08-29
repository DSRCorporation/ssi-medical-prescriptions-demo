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

	"github.com/syndtr/goleveldb/leveldb"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type PatientStorage struct {
	levelDB *LevelDB
}

type patient struct {
	PatientId      string   `json:"patientId"`
	Username       string   `json:"username"`
	HashedPassword string   `json:"hashedPassword"`
	Dids           []string `json:"dids"`
	CredentialIds  []string `json:"credentialIds"`
}

func NewPatientStorage(path string) (*PatientStorage, error) {
	levelDB, err := NewLevelDB(path)
	if err != nil {
		return nil, err
	}

	return &PatientStorage{levelDB: levelDB}, nil
}

func (s *PatientStorage) AddCredentialIdByPatientId(patientId string, credentialId string) (err error) {
	if patientId == "" {
		return fmt.Errorf("patientId cannot be empty")
	}

	exist, err := s.levelDB.Has(patientId)
	if err != nil {
		return err
	}

	var patientInfo patient

	if exist {
		if err = s.levelDB.ReadFromJson(patientId, &patientInfo); err != nil {
			return err
		}

		patientInfo.CredentialIds = append(patientInfo.CredentialIds, credentialId)
	} else {
		return fmt.Errorf("no patient found for patientId: %s", patientId)
	}

	if err = s.levelDB.WriteAsJson(patientId, patientInfo); err != nil {
		return err
	}

	return nil
}

func (s *PatientStorage) GetCredentialIdsByPatientId(patientId string) (credentialIds []string, err error) {
	var patientInfo patient

	if err = s.levelDB.ReadFromJson(patientId, &patientInfo); err != nil {
		return credentialIds, fmt.Errorf("no patient found for patientId: %s", patientId)
	}

	return patientInfo.CredentialIds, nil
}

func (s *PatientStorage) GetDIDs(patientId string) (dids []string, err error) {
	var patientInfo patient

	if err = s.levelDB.ReadFromJson(patientId, &patientInfo); err != nil {
		return dids, fmt.Errorf("no dids found for patientId: %s", patientId)
	}

	dids = patientInfo.Dids

	return dids, nil
}

func (s *PatientStorage) PatientExists(username string) bool {
	exist, err := s.levelDB.Has(username)
	if err != nil {
		return false
	}

	return exist
}

func (s *PatientStorage) CreatePatient(username string, password string) (*domain.Patient, error) {
	minCountOfCharacters := 4
	maxCountOfCharacters := 100

	if len(username) < minCountOfCharacters || len(username) > maxCountOfCharacters {
		return nil, fmt.Errorf("username should be between %d and %d characters",
			minCountOfCharacters, maxCountOfCharacters)
	}
	if len(password) < minCountOfCharacters || len(password) > maxCountOfCharacters {
		return nil, fmt.Errorf("password should be between %d and %d characters",
			minCountOfCharacters, maxCountOfCharacters)
	}

	if s.PatientExists(username) {
		return nil, fmt.Errorf("username already exists")
	}

	patientId, err := createUniquePatientIdForPatientDatabase(s.levelDB.path)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	patientInfo := patient{
		PatientId:      patientId,
		Username:       username,
		HashedPassword: hashedPassword,
	}

	if err := s.levelDB.WriteAsJson(patientInfo.PatientId, patientInfo); err != nil {
		return nil, err
	}

	if err := s.levelDB.Write(patientInfo.Username, patientInfo.PatientId); err != nil {
		return nil, err
	}

	return &domain.Patient{PatientId: patientInfo.PatientId, Username: patientInfo.Username}, nil
}

func (s *PatientStorage) GetPatientByCredentials(username string, password string) (*domain.Patient, error) {
	if username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}

	patientId, err := s.levelDB.Read(username)
	if err != nil {
		return nil, fmt.Errorf("no patient found for username: %v", username)
	}

	var patientInfo patient

	if err := s.levelDB.ReadFromJson(patientId, &patientInfo); err != nil {
		return nil, fmt.Errorf("no patient found for patientId: %s", patientId)
	}

	if checkPasswordHash(password, patientInfo.HashedPassword) {
		return &domain.Patient{PatientId: patientInfo.PatientId, Username: patientInfo.Username}, nil
	}

	return nil, fmt.Errorf("Incorrect password for username: %s", username)
}

func (s *PatientStorage) AddPatientDID(patientId string, did string) (err error) {
	if patientId == "" {
		return fmt.Errorf("patientId cannot be empty")
	}
	if did == "" {
		return fmt.Errorf("did cannot be empty")
	}

	var patientInfo patient

	if err := s.levelDB.ReadFromJson(patientId, &patientInfo); err != nil {
		return fmt.Errorf("no patient found for patientId: %s", patientId)
	}

	for _, existsDid := range patientInfo.Dids {
		if did == existsDid {
			return fmt.Errorf("We already have did: %s for this patientId: %s", existsDid, patientId)
		}
	}

	patientInfo.Dids = append(patientInfo.Dids, did)

	if err := s.levelDB.WriteAsJson(patientInfo.PatientId, patientInfo); err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func createUniquePatientIdForPatientDatabase(dbPath string) (string, error) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return "", err
	}
	defer db.Close()

	patientId := tmrand.Str(6)
	exist, err := db.Has([]byte(patientId), nil)
	if err != nil {
		return "", err
	}

	maxIteration := 10000
	for i := 0; exist && i < maxIteration; i++ {
		patientId = tmrand.Str(6)

		exist, err = db.Has([]byte(patientId), nil)
		if err != nil {
			return "", err
		}
	}

	return patientId, nil
}
