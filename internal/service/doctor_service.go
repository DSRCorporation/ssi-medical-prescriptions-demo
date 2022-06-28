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

func (s *DoctorService) CreatePrescriptionCredentialOffer(doctorId string, prescription domain.Prescription) (credentialOfferId string, err error) {
	// @TODO: check prescription data
	return s.storage.CreatePrescriptionCredentialOffer(doctorId, prescription)
}

func (s *DoctorService) GetPrescriptionByCredentialOfferId(doctorId string, credentialOfferId string) (prescription domain.Prescription, err error) {
	return s.storage.GetPrescriptionByCredentialOfferId(doctorId, credentialOfferId)
}

func (s *DoctorService) SavePrescriptionCredential(doctorId string, credentialOfferId string, credential domain.Credential) (err error) {
	return s.storage.SavePrescriptionCredential(doctorId, credentialOfferId, credential)
}

func (s *DoctorService) GetPrescriptionCredentialByCredentialOfferId(doctorId string, credentialOfferId string) (credential domain.Credential, err error) {
	return s.storage.GetPrescriptionCredentialByCredentialOfferId(doctorId, credentialOfferId)
}
