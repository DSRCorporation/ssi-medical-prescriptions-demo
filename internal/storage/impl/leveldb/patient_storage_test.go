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

func TestGetExistCredentialIDsByPatientID(t *testing.T) {
	var patientId = tmrand.Str(6)
	var credential1 = tmrand.Str(10)
	var credential2 = tmrand.Str(10)
	var credential3 = tmrand.Str(10)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	err = patientStorage.AddCredentialIdByPatientId(patientId, credential1)
	require.NoError(t, err)

	credentials, err := patientStorage.GetCredentialIdsByPatientId(patientId)
	require.NoError(t, err)
	require.Equal(t, 1, len(credentials))
	require.Equal(t, credential1, credentials[0])

	err = patientStorage.AddCredentialIdByPatientId(patientId, credential2)
	require.NoError(t, err)

	credentials, err = patientStorage.GetCredentialIdsByPatientId(patientId)
	require.NoError(t, err)
	require.Equal(t, 2, len(credentials))
	require.Equal(t, credential1, credentials[0])
	require.Equal(t, credential2, credentials[1])

	err = patientStorage.AddCredentialIdByPatientId(patientId, credential3)
	require.NoError(t, err)

	credentials, err = patientStorage.GetCredentialIdsByPatientId(patientId)
	require.NoError(t, err)
	require.Equal(t, 3, len(credentials))
	require.Equal(t, credential1, credentials[0])
	require.Equal(t, credential2, credentials[1])
	require.Equal(t, credential3, credentials[2])
}

func TestGetNotExistCredentialIDsByPatientID(t *testing.T) {
	var patientId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	patientStorage, err := NewPatientStorage(dbPath)
	require.NoError(t, err)

	_, err = patientStorage.GetCredentialIdsByPatientId(patientId)
	require.Error(t, err)
}
