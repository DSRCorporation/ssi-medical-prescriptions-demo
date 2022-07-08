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
	"testing"

	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
)

func TestDoubleTimeCreatePresentationRequest(t *testing.T) {
	var pharmacyId = tmrand.Str(6)
	var requestId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	pharmacyStorage, err := NewPharmacyStorage(dbPath)
	require.NoError(t, err)

	err = pharmacyStorage.CreatePresentationRequest(pharmacyId, requestId)
	require.NoError(t, err)

	err = pharmacyStorage.CreatePresentationRequest(pharmacyId, requestId)
	require.Error(t, err)
}

func TestGetExistPharmacyIDByRequestID(t *testing.T) {
	var pharmacyId = tmrand.Str(6)
	var requestId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	pharmacyStorage, err := NewPharmacyStorage(dbPath)
	require.NoError(t, err)

	err = pharmacyStorage.CreatePresentationRequest(pharmacyId, requestId)
	require.NoError(t, err)

	receivedPharmacyId, err := pharmacyStorage.GetPharmacyIdByRequestId(requestId)
	require.NoError(t, err)
	require.Equal(t, pharmacyId, receivedPharmacyId)
}

func TestGetNotExistPharmacyIDByRequestID(t *testing.T) {
	var requestId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	pharmacyStorage, err := NewPharmacyStorage(dbPath)
	require.NoError(t, err)

	_, err = pharmacyStorage.GetPharmacyIdByRequestId(requestId)
	require.Error(t, err)
}

func TestAddPresentationIDByRequestIDWithoutPresentationRequest(t *testing.T) {
	var requestId = tmrand.Str(6)
	var presentationId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	pharmacyStorage, err := NewPharmacyStorage(dbPath)
	require.NoError(t, err)

	err = pharmacyStorage.AddPresentationIdByRequestId(requestId, presentationId)
	require.Error(t, err)
}

func TestDoubleTimeAddPresentationIDByRequestID(t *testing.T) {
	var pharmacyId = tmrand.Str(6)
	var requestId = tmrand.Str(6)
	var presentationId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	pharmacyStorage, err := NewPharmacyStorage(dbPath)
	require.NoError(t, err)

	err = pharmacyStorage.CreatePresentationRequest(pharmacyId, requestId)
	require.NoError(t, err)

	err = pharmacyStorage.AddPresentationIdByRequestId(requestId, presentationId)
	require.NoError(t, err)

	err = pharmacyStorage.AddPresentationIdByRequestId(requestId, presentationId)
	require.Error(t, err)
}

func TestGetExistPresentationIDByRequestID(t *testing.T) {
	var pharmacyId = tmrand.Str(6)
	var requestId = tmrand.Str(6)
	var presentationId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	pharmacyStorage, err := NewPharmacyStorage(dbPath)
	require.NoError(t, err)

	err = pharmacyStorage.CreatePresentationRequest(pharmacyId, requestId)
	require.NoError(t, err)

	err = pharmacyStorage.AddPresentationIdByRequestId(requestId, presentationId)
	require.NoError(t, err)

	receivedPresentationId, err := pharmacyStorage.GetPresentationIdByRequestId(requestId)
	require.NoError(t, err)
	require.Equal(t, presentationId, receivedPresentationId)
}

func TestAlreadyCreatePresentationRequestAndGetNotExistPresentationIDByRequestID(t *testing.T) {
	var pharmacyId = tmrand.Str(6)
	var requestId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	pharmacyStorage, err := NewPharmacyStorage(dbPath)
	require.NoError(t, err)

	// Create Presentation Request
	err = pharmacyStorage.CreatePresentationRequest(pharmacyId, requestId)
	require.NoError(t, err)

	// We don't add Presentation ID by Request ID and
	// Get not exist Presentation ID by Request ID, it should be error.
	_, err = pharmacyStorage.GetPresentationIdByRequestId(requestId)
	require.Error(t, err)
}

func TestNotCreatePresentationRequestAndGetNotExistPresentationIDByRequestID(t *testing.T) {
	var requestId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	pharmacyStorage, err := NewPharmacyStorage(dbPath)
	require.NoError(t, err)

	_, err = pharmacyStorage.GetPresentationIdByRequestId(requestId)
	require.Error(t, err)
}

func TestGetExistPresentationIDByPharmacyID(t *testing.T) {
	var pharmacyId = tmrand.Str(6)
	var presentationId1 = tmrand.Str(10)
	var presentationId2 = tmrand.Str(10)
	var presentationId3 = tmrand.Str(10)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	pharmacyStorage, err := NewPharmacyStorage(dbPath)
	require.NoError(t, err)

	err = pharmacyStorage.AddPresentationIdByPharmacyId(pharmacyId, presentationId1)
	require.NoError(t, err)

	receivedPresentationIds, err := pharmacyStorage.GetPresentationIdsByPharmacyId(pharmacyId)
	require.NoError(t, err)
	require.Equal(t, 1, len(receivedPresentationIds))
	require.Equal(t, presentationId1, receivedPresentationIds[0])

	err = pharmacyStorage.AddPresentationIdByPharmacyId(pharmacyId, presentationId2)
	require.NoError(t, err)

	receivedPresentationIds, err = pharmacyStorage.GetPresentationIdsByPharmacyId(pharmacyId)
	require.NoError(t, err)
	require.Equal(t, 2, len(receivedPresentationIds))
	require.Equal(t, presentationId1, receivedPresentationIds[0])
	require.Equal(t, presentationId2, receivedPresentationIds[1])

	err = pharmacyStorage.AddPresentationIdByPharmacyId(pharmacyId, presentationId3)
	require.NoError(t, err)

	receivedPresentationIds, err = pharmacyStorage.GetPresentationIdsByPharmacyId(pharmacyId)
	require.NoError(t, err)
	require.Equal(t, 3, len(receivedPresentationIds))
	require.Equal(t, presentationId1, receivedPresentationIds[0])
	require.Equal(t, presentationId2, receivedPresentationIds[1])
	require.Equal(t, presentationId3, receivedPresentationIds[2])
}

func TestGetNotExistPresentationIDByPharmacyID(t *testing.T) {
	var pharmacyId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	pharmacyStorage, err := NewPharmacyStorage(dbPath)
	require.NoError(t, err)

	_, err = pharmacyStorage.GetPresentationIdsByPharmacyId(pharmacyId)
	require.Error(t, err)
}
