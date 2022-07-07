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

import "fmt"

type PharmacyStorage struct {
	levelDB *LevelDB
}

func NewPharmacyStorage(dbPath string) (*PharmacyStorage, error) {
	levelDB, err := NewLevelDB(dbPath)
	if err != nil {
		return nil, err
	}

	return &PharmacyStorage{levelDB: levelDB}, nil
}

func (s *PharmacyStorage) CreatePresentationRequest(pharmacyId string, requestId string) (err error) {
	exist, err := s.levelDB.Has(requestId)
	if err != nil {
		return err
	}

	if exist {
		return fmt.Errorf("pharmacyId already exists: %v", requestId)
	}

	presentationRequest := PresentationRequest{
		RequestId:      &requestId,
		PharmacyId:     &pharmacyId,
		PresentationId: nil,
	}

	if err = s.levelDB.WriteAsJson(requestId, presentationRequest); err != nil {
		return err
	}

	return nil
}

func (s *PharmacyStorage) GetPharmacyIdByRequestId(requestId string) (pharmacyId string, err error) {
	var presentationRequest PresentationRequest

	if err = s.levelDB.ReadFromJson(requestId, &presentationRequest); err != nil {
		return pharmacyId, err
	}

	pharmacyId = *presentationRequest.PharmacyId

	return pharmacyId, err
}

func (s *PharmacyStorage) AddPresentationIdByRequestId(requestId string, presentationId string) (err error) {
	var presentationRequest PresentationRequest

	if err = s.levelDB.ReadFromJson(requestId, &presentationRequest); err != nil {
		return err
	}

	if presentationRequest.PresentationId != nil {
		return fmt.Errorf("requestId already exists: %v", requestId)
	}

	presentationRequest.PresentationId = &presentationId

	if err = s.levelDB.WriteAsJson(requestId, presentationRequest); err != nil {
		return err
	}

	return nil
}

func (s *PharmacyStorage) GetPresentationIdByRequestId(requestId string) (presentationId string, err error) {
	var presentationRequest PresentationRequest

	if err = s.levelDB.ReadFromJson(requestId, &presentationRequest); err != nil {
		return presentationId, err
	}

	if presentationRequest.PresentationId == nil {
		return presentationId, fmt.Errorf("no presentation id for request id: %s", requestId)
	}

	presentationId = *presentationRequest.PresentationId

	return presentationId, nil
}

func (s *PharmacyStorage) AddPresentationIdByPharmacyId(pharmacyId string, presentationId string) (err error) {
	exist, err := s.levelDB.Has(pharmacyId)
	if err != nil {
		return err
	}

	var presentationIds []string

	if exist {
		if err = s.levelDB.ReadFromJson(pharmacyId, &presentationIds); err != nil {
			return err
		}
	}

	presentationIds = append(presentationIds, presentationId)

	if err = s.levelDB.WriteAsJson(pharmacyId, presentationIds); err != nil {
		return err
	}

	return nil
}

func (s *PharmacyStorage) GetPresentationIdsByPharmacyId(pharmacyId string) (presentationIds []string, err error) {
	if err = s.levelDB.ReadFromJson(pharmacyId, &presentationIds); err != nil {
		return presentationIds, err
	}

	return presentationIds, nil
}

type PresentationRequest struct {
	RequestId      *string
	PharmacyId     *string
	PresentationId *string
}
