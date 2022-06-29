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

type DoctorService struct {
	storage storage.DoctorStorage
}

func (s *DoctorService) CreatePrescriptionOffer(offerId string, prescription domain.Prescription) (err error) {
	// @TODO: check prescription data
	return s.storage.CreatePrescriptionOffer(offerId, prescription)
}

func (s *DoctorService) GetPrescriptionByOfferId(offerId string) (prescription domain.Prescription, err error) {
	return s.storage.GetPrescriptionByOfferId(offerId)
}

func (s *DoctorService) SaveCredentialId(doctorId string, credentialOfferId string, credentialId string) (err error) {
	err = s.storage.AddCredentialIdByDoctorId(doctorId, credentialId)
	if err != nil {
		return err
	}

	err = s.storage.AddCredentialIdByOfferId(credentialOfferId, credentialId)
	if err != nil {
		return err
	}
	return
}

func (s *DoctorService) GetCredentialIdByOfferId(offerId string) (credentialId string, err error) {
	return s.storage.GetCredentialIdByOfferId(offerId)
}

func (s *DoctorService) GetCredentialIdsByDoctorId(offerId string) (credentialId []string, err error) {
	return s.storage.GetCredentialIdsByDoctorId(offerId)
}

func (s *DoctorService) GetKMSPassphrase(doctorId string) (kmsPassphrase string, err error) {
	return s.storage.GetKMSPassphrase(doctorId)
}

func (s *DoctorService) GetDID(doctorId string) (did string, err error) {
	return s.storage.GetDID(doctorId)
}
