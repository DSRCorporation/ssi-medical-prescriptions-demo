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

type VCStorage struct {
	levelDB *LevelDB
}

func NewVCStorage(dbPath string) (*VCStorage, error) {
	levelDB, err := NewLevelDB(dbPath)
	if err != nil {
		return nil, err
	}
	return &VCStorage{levelDB: levelDB}, nil
}

func (s *VCStorage) SaveConnection(inviterId string, inviteeId string, connection domain.Connection) (err error) {
	key := makeKey(inviterId, inviteeId)

	exist, err := s.levelDB.Has(key)
	if err != nil {
		return err
	}

	if exist {
		return fmt.Errorf("inviterId: %s and inviteeId: %s already exist", inviterId, inviteeId)
	}

	if err = s.levelDB.WriteAsJson(key, connection); err != nil {
		return err
	}

	return nil
}

func (s *VCStorage) GetConnection(inviterId string, inviteeId string) (connection domain.Connection, err error) {
	key := makeKey(inviterId, inviteeId)

	if err = s.levelDB.ReadFromJson(key, &connection); err != nil {
		return connection, err
	}

	return connection, err
}

func (s *VCStorage) SaveCredential(credential domain.Credential) error {
	if credential.CredentialId == "" {
		return fmt.Errorf("credentialId cannot be empty")
	}

	key := credential.CredentialId

	exist, err := s.levelDB.Has(key)
	if err != nil {
		return err
	}

	if exist {
		return fmt.Errorf("credentialId already exists: %v", key)
	}

	if err = s.levelDB.WriteAsJson(key, credential); err != nil {
		return err
	}

	return nil
}

func (s *VCStorage) GetCredentialById(credentialId string) (credential domain.Credential, err error) {
	if err = s.levelDB.ReadFromJson(credentialId, &credential); err != nil {
		return credential, err
	}

	return credential, nil
}

func (s *VCStorage) SavePresentation(presentation domain.Presentation) (err error) {
	if presentation.PresentationId == "" {
		return fmt.Errorf("presentationId cannot be empty")
	}

	key := presentation.PresentationId

	exist, err := s.levelDB.Has(key)
	if err != nil {
		return err
	}

	if exist {
		return fmt.Errorf("presentationId already exists: %v", key)
	}

	if err = s.levelDB.WriteAsJson(key, presentation); err != nil {
		return err
	}

	return nil
}

func (s *VCStorage) GetPresentationById(presentationId string) (presentation domain.Presentation, err error) {
	if err = s.levelDB.ReadFromJson(presentationId, &presentation); err != nil {
		return presentation, err
	}

	return presentation, nil
}

func (s *VCStorage) GetWalletCredentialsForVerification() (userId string, passphrase string, err error) {
	return "v0001", "Np6VR4Yg6PPL", nil
}

func makeKey(inviterId string, inviteeId string) (key string) {
	key = inviterId + inviteeId

	return key
}
