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
	"testing"

	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
)

func TestCreatePatientWithValidUsernameAndPassword(t *testing.T) {
	var username = tmrand.Str(6)
	var password = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	domain, err := patientStorage.CreatePatient(username, password)
	require.NoError(t, err)
	require.Equal(t, domain.Username, username)
}

func TestCreatePatientWithInvalidUsername(t *testing.T) {
	var username = tmrand.Str(3)
	var password = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	// minimum lenght of username should be 4 characters
	_, err = patientStorage.CreatePatient(username, password)
	require.Error(t, err)
	require.ErrorContains(t, err, "username should be between 4 and 100 characters")

	username = tmrand.Str(101)

	// maximum lenght of username should be 100 characters
	_, err = patientStorage.CreatePatient(username, password)
	require.Error(t, err)
	require.ErrorContains(t, err, "username should be between 4 and 100 characters")
}

func TestCreatePatientWithInvalidPassword(t *testing.T) {
	var username = tmrand.Str(6)
	var password = tmrand.Str(3)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	// minimum lenght of password should be 4 characters
	_, err = patientStorage.CreatePatient(username, password)
	require.Error(t, err)
	require.ErrorContains(t, err, "password should be between 4 and 100 characters")

	password = tmrand.Str(101)

	// maximum lenght of password should be 100 characters
	_, err = patientStorage.CreatePatient(username, password)
	require.Error(t, err)
	require.ErrorContains(t, err, "password should be between 4 and 100 characters")
}

func TestTwoTimesCreateSamePatient(t *testing.T) {
	var username = tmrand.Str(6)
	var password = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	_, err = patientStorage.CreatePatient(username, password)
	require.NoError(t, err)

	_, err = patientStorage.CreatePatient(username, password)
	require.Error(t, err)
	require.ErrorContains(t, err, "username already exists")
}

func TestLoginWithExistUsernameAndPassword(t *testing.T) {
	var username = tmrand.Str(6)
	var password = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	registrationDomain, err := patientStorage.CreatePatient(username, password)
	require.NoError(t, err)

	loginDomain, err := patientStorage.GetPatientByCredentials(username, password)
	require.NoError(t, err)
	require.Equal(t, registrationDomain.PatientId, loginDomain.PatientId)
	require.Equal(t, registrationDomain.Username, loginDomain.Username)
}

func TestLoginWithNotExistUsernameAndPassword(t *testing.T) {
	var username = tmrand.Str(6)
	var password = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	_, err = patientStorage.GetPatientByCredentials(username, password)
	require.Error(t, err)
	require.ErrorContains(t, err, fmt.Sprintf("no patient found for username: %v", username))
}

func TestLoginWithInvalidUsername(t *testing.T) {
	var username string
	var password = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	_, err = patientStorage.GetPatientByCredentials(username, password)
	require.Error(t, err)
	require.ErrorContains(t, err, "username cannot be empty")
}

func TestLoginWithInvalidPassword(t *testing.T) {
	var username = tmrand.Str(6)
	var password = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	_, err = patientStorage.CreatePatient(username, password)
	require.NoError(t, err)

	var invalidPassword = tmrand.Str(6)

	_, err = patientStorage.GetPatientByCredentials(username, invalidPassword)
	require.Error(t, err)
	require.ErrorContains(t, err, fmt.Sprintf("Incorrect password for username: %s", username))
}

func TestAddPatientDIDWithValidPatientIDAndDID(t *testing.T) {
	var username = tmrand.Str(6)
	var password = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	domain, err := patientStorage.CreatePatient(username, password)
	require.NoError(t, err)

	var did = fmt.Sprintf("did:key:%s", tmrand.Str(32))
	require.NoError(t, patientStorage.AddPatientDID(domain.PatientId, did))
}

func TestAddPatientDIDWithInvalidPatientID(t *testing.T) {
	var patientId = tmrand.Str(5)
	var did = fmt.Sprintf("did:key:%s", tmrand.Str(32))

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	// doesn't exist patient id
	err = patientStorage.AddPatientDID(patientId, did)
	require.Error(t, err)
	require.ErrorContains(t, err, fmt.Sprintf("no patient found for patientId: %s", patientId))
}

func TestAddPatientDIDWithInvalidDID(t *testing.T) {
	var username = tmrand.Str(6)
	var password = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	domain, err := patientStorage.CreatePatient(username, password)
	require.NoError(t, err)

	var did string

	err = patientStorage.AddPatientDID(domain.PatientId, did)
	require.Error(t, err)
	require.ErrorContains(t, err, "did cannot be empty")
}

func TestDoubleTimeAddPatientDID(t *testing.T) {
	var username = tmrand.Str(6)
	var password = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	domain, err := patientStorage.CreatePatient(username, password)
	require.NoError(t, err)

	var did = fmt.Sprintf("did:key:%s", tmrand.Str(32))
	require.NoError(t, patientStorage.AddPatientDID(domain.PatientId, did))

	err = patientStorage.AddPatientDID(domain.PatientId, did)
	require.Error(t, err)
	require.ErrorContains(t, err, fmt.Sprintf("We already have did: %s for this patientId: %s", did, domain.PatientId))
}

func TestGetExistCredentialIDsByPatientID(t *testing.T) {
	var username = tmrand.Str(6)
	var password = tmrand.Str(6)

	var credential1 = tmrand.Str(10)
	var credential2 = tmrand.Str(10)
	var credential3 = tmrand.Str(10)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	domain, err := patientStorage.CreatePatient(username, password)
	require.NoError(t, err)

	err = patientStorage.AddCredentialIdByPatientId(domain.PatientId, credential1)
	require.NoError(t, err)

	credentials, err := patientStorage.GetCredentialIdsByPatientId(domain.PatientId)
	require.NoError(t, err)
	require.Equal(t, 1, len(credentials))
	require.Equal(t, credential1, credentials[0])

	err = patientStorage.AddCredentialIdByPatientId(domain.PatientId, credential2)
	require.NoError(t, err)

	credentials, err = patientStorage.GetCredentialIdsByPatientId(domain.PatientId)
	require.NoError(t, err)
	require.Equal(t, 2, len(credentials))
	require.Equal(t, credential1, credentials[0])
	require.Equal(t, credential2, credentials[1])

	err = patientStorage.AddCredentialIdByPatientId(domain.PatientId, credential3)
	require.NoError(t, err)

	credentials, err = patientStorage.GetCredentialIdsByPatientId(domain.PatientId)
	require.NoError(t, err)
	require.Equal(t, 3, len(credentials))
	require.Equal(t, credential1, credentials[0])
	require.Equal(t, credential2, credentials[1])
	require.Equal(t, credential3, credentials[2])
}

func TestGetNotExistCredentialIDsByNotExistPatientID(t *testing.T) {
	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	var patientId = tmrand.Str(6)

	_, err = patientStorage.GetCredentialIdsByPatientId(patientId)
	require.Error(t, err)
	require.ErrorContains(t, err, fmt.Sprintf("no patient found for patientId: %s", patientId))
}

func TestCreatePatientWithValidUsernameAndPassword(t *testing.T) {
	var username = tmrand.Str(6)
	var password = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	domain, err := patientStorage.CreatePatient(username, password)
	require.NoError(t, err)
	require.Equal(t, domain.Username, username)
}
