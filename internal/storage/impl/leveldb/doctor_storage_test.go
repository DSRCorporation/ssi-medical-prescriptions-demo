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
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

func TestDoubleTimeCreatePrescriptionOfferWithSameOfferID(t *testing.T) {
	var offerId = tmrand.Str(6)
	var prescriptions = domain.Prescription{
		DoctorId:        tmrand.Str(6),
		RawPrescription: []byte(`{"some":"some"}`),
	}

	var dbPath = GenerateDBPath()
	defer cleanUp(dbPath)

	doctorStorage, err := NewDoctorStorage(dbPath)
	require.NoError(t, err)

	err = doctorStorage.CreatePrescriptionOffer(offerId, prescriptions)
	require.NoError(t, err)

	err = doctorStorage.CreatePrescriptionOffer(offerId, prescriptions)
	require.Error(t, err)
}

func TestGetExistPrescription(t *testing.T) {
	var offerId = tmrand.Str(6)
	var prescriptions = domain.Prescription{
		DoctorId:        tmrand.Str(6),
		RawPrescription: []byte(`{"some":"some"}`),
	}

	var dbPath = GenerateDBPath()
	defer cleanUp(dbPath)

	doctorStorage, err := NewDoctorStorage(dbPath)
	require.NoError(t, err)

	err = doctorStorage.CreatePrescriptionOffer(offerId, prescriptions)
	require.NoError(t, err)

	prescription, err := doctorStorage.GetPrescriptionByOfferId(offerId)
	require.NoError(t, err)
	require.Equal(t, prescriptions.DoctorId, prescription.DoctorId)
	require.Equal(t, prescriptions.RawPrescription, prescription.RawPrescription)
}

func TestGetNotExistPrescription(t *testing.T) {
	var offerId = tmrand.Str(6)

	var dbPath = GenerateDBPath()
	defer cleanUp(dbPath)

	doctorStorage, err := NewDoctorStorage(dbPath)
	require.NoError(t, err)

	_, err = doctorStorage.GetPrescriptionByOfferId(offerId)
	require.Error(t, leveldb.ErrNotFound, err)
}

func TestCreateCredentialIDByOfferIDWithoutPrescriptionOfferID(t *testing.T) {
	var offerId = tmrand.Str(6)
	var credentialId = tmrand.Str(10)

	var dbPath = GenerateDBPath()
	defer cleanUp(dbPath)

	doctorStorage, err := NewDoctorStorage(dbPath)
	require.NoError(t, err)

	err = doctorStorage.AddCredentialIdByOfferId(offerId, credentialId)
	require.Error(t, err)
}

func TestDoubleTimeCreateCredentialIDByOfferID(t *testing.T) {
	var offerId = tmrand.Str(6)
	var prescriptions = domain.Prescription{
		DoctorId:        tmrand.Str(6),
		RawPrescription: []byte(`{"some":"some"}`),
	}
	var credentialId = tmrand.Str(10)

	var dbPath = GenerateDBPath()
	defer cleanUp(dbPath)

	doctorStorage, err := NewDoctorStorage(dbPath)
	require.NoError(t, err)

	err = doctorStorage.CreatePrescriptionOffer(offerId, prescriptions)
	require.NoError(t, err)

	err = doctorStorage.AddCredentialIdByOfferId(offerId, credentialId)
	require.NoError(t, err)

	err = doctorStorage.AddCredentialIdByOfferId(offerId, credentialId)
	require.Error(t, err)
}

func TestGetExistCredentialIDByOfferID(t *testing.T) {
	var offerId = tmrand.Str(6)
	var prescriptions = domain.Prescription{
		DoctorId:        tmrand.Str(6),
		RawPrescription: []byte(`{"some":"some"}`),
	}
	var credentialId = tmrand.Str(10)

	var dbPath = GenerateDBPath()
	defer cleanUp(dbPath)

	doctorStorage, err := NewDoctorStorage(dbPath)
	require.NoError(t, err)

	err = doctorStorage.CreatePrescriptionOffer(offerId, prescriptions)
	require.NoError(t, err)

	prescription, err := doctorStorage.GetPrescriptionByOfferId(offerId)
	require.NoError(t, err)
	require.Equal(t, prescriptions.DoctorId, prescription.DoctorId)
	require.Equal(t, prescriptions.RawPrescription, prescription.RawPrescription)

	err = doctorStorage.AddCredentialIdByOfferId(offerId, credentialId)
	require.NoError(t, err)

	credential, err := doctorStorage.GetCredentialIdByOfferId(offerId)
	require.NoError(t, err)
	require.Equal(t, credentialId, credential)
}

func TestAlreadyCreatedPrescriptionOfferAndGetNotExistCredentialsIDByOfferID(t *testing.T) {
	var offerId = tmrand.Str(6)
	var prescriptions = domain.Prescription{
		DoctorId:        tmrand.Str(6),
		RawPrescription: []byte(`{"some":"some"}`),
	}

	var dbPath = GenerateDBPath()
	defer cleanUp(dbPath)

	doctorStorage, err := NewDoctorStorage(dbPath)
	require.NoError(t, err)

	// Create Prescription Offer
	err = doctorStorage.CreatePrescriptionOffer(offerId, prescriptions)
	require.NoError(t, err)

	// We don't create Credential ID by Offer ID and
	// Get not exist Credential ID by Offer ID, it should be error.
	_, err = doctorStorage.GetCredentialIdByOfferId(offerId)
	require.Error(t, err)
}

func TestNotCreatePrescriptionOfferAndGetNotExistCredentialsIDByOfferId(t *testing.T) {
	var offerId = tmrand.Str(6)

	var dbPath = GenerateDBPath()
	defer cleanUp(dbPath)

	doctorStorage, err := NewDoctorStorage(dbPath)
	require.NoError(t, err)

	_, err = doctorStorage.GetCredentialIdByOfferId(offerId)
	require.Error(t, err)
}

func TestGetExistCredentialIDByDoctorID(t *testing.T) {
	var doctorId = tmrand.Str(6)
	var credential1 = tmrand.Str(10)
	var credential2 = tmrand.Str(10)
	var credential3 = tmrand.Str(10)

	var dbPath = GenerateDBPath()
	defer cleanUp(dbPath)

	doctorStorage, err := NewDoctorStorage(dbPath)
	require.NoError(t, err)

	err = doctorStorage.AddCredentialIdByDoctorId(doctorId, credential1)
	require.NoError(t, err)

	credentials, err := doctorStorage.GetCredentialIdsByDoctorId(doctorId)
	require.NoError(t, err)
	require.Equal(t, 1, len(credentials))
	require.Equal(t, credential1, credentials[0])

	err = doctorStorage.AddCredentialIdByDoctorId(doctorId, credential2)
	require.NoError(t, err)

	credentials, err = doctorStorage.GetCredentialIdsByDoctorId(doctorId)
	require.NoError(t, err)
	require.Equal(t, 2, len(credentials))
	require.Equal(t, credential1, credentials[0])
	require.Equal(t, credential2, credentials[1])

	err = doctorStorage.AddCredentialIdByDoctorId(doctorId, credential3)
	require.NoError(t, err)

	credentials, err = doctorStorage.GetCredentialIdsByDoctorId(doctorId)
	require.NoError(t, err)
	require.Equal(t, 3, len(credentials))
	require.Equal(t, credential1, credentials[0])
	require.Equal(t, credential2, credentials[1])
	require.Equal(t, credential3, credentials[2])
}

func TestGetNotExistCredentialsIDByDoctorID(t *testing.T) {
	var doctorId = tmrand.Str(6)

	var dbPath = GenerateDBPath()
	defer cleanUp(dbPath)

	doctorStorage, err := NewDoctorStorage(dbPath)
	require.NoError(t, err)

	_, err = doctorStorage.GetCredentialIdsByDoctorId(doctorId)
	require.Error(t, leveldb.ErrNotFound, err)
}

func cleanUp(dbPath string) {
	os.RemoveAll(dbPath)
}
