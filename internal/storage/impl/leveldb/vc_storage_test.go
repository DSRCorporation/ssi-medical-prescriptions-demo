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
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

func TestDoubleTimeCreateConnection(t *testing.T) {
	var inviterId = tmrand.Str(3)
	var inviteeId = tmrand.Str(3)
	var connection = domain.Connection{
		InviterDID:   fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
		InviteeDID:   fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
		ConnectionId: tmrand.Str(10),
	}

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	vcStorage, err := NewVCStorage(dbPath)
	require.NoError(t, err)

	err = vcStorage.SaveConnection(inviterId, inviteeId, connection)
	require.NoError(t, err)

	err = vcStorage.SaveConnection(inviterId, inviteeId, connection)
	require.Error(t, err)
}

func TestGetExistConnection(t *testing.T) {
	var inviterId = tmrand.Str(3)
	var inviteeId = tmrand.Str(3)
	var connection = domain.Connection{
		InviterDID:   fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
		InviteeDID:   fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
		ConnectionId: tmrand.Str(10),
	}

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	vcStorage, err := NewVCStorage(dbPath)
	require.NoError(t, err)

	err = vcStorage.SaveConnection(inviterId, inviteeId, connection)
	require.NoError(t, err)

	receivedConnection, err := vcStorage.GetConnection(inviterId, inviteeId)
	require.NoError(t, err)
	require.Equal(t, connection.InviterDID, receivedConnection.InviterDID)
	require.Equal(t, connection.InviteeDID, receivedConnection.InviteeDID)
	require.Equal(t, connection.ConnectionId, receivedConnection.ConnectionId)
}

func TestGetNotExistConnection(t *testing.T) {
	var inviterId = tmrand.Str(3)
	var inviteeId = tmrand.Str(3)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	vcStorage, err := NewVCStorage(dbPath)
	require.NoError(t, err)

	_, err = vcStorage.GetConnection(inviterId, inviteeId)
	require.Error(t, err)
}

func TestDoubleTimeSaveCredential(t *testing.T) {
	var credential = domain.Credential{
		CredentialId: tmrand.Str(6),
		IssuerDID:    fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
		HolderDID:    fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
		Prescription: domain.Prescription{
			DoctorId:        tmrand.Str(6),
			RawPrescription: []byte(`{"somePrescription":"someValue"}`),
		},
		RawCredentialWithProof: []byte(`{"someCredential":"someValue"}`),
	}

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	vcStorage, err := NewVCStorage(dbPath)
	require.NoError(t, err)

	err = vcStorage.SaveCredential(credential)
	require.NoError(t, err)

	err = vcStorage.SaveCredential(credential)
	require.Error(t, err)
}

func TestGetExistCredentialByID(t *testing.T) {
	var credential = domain.Credential{
		CredentialId: tmrand.Str(6),
		IssuerDID:    fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
		HolderDID:    fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
		Prescription: domain.Prescription{
			DoctorId:        tmrand.Str(6),
			RawPrescription: []byte(`{"somePrescription":"someValue"}`),
		},
		RawCredentialWithProof: []byte(`{"someCredential":"someValue"}`),
	}

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	vcStorage, err := NewVCStorage(dbPath)
	require.NoError(t, err)

	err = vcStorage.SaveCredential(credential)
	require.NoError(t, err)

	receivedCredential, err := vcStorage.GetCredentialById(credential.CredentialId)
	require.NoError(t, err)
	require.Equal(t, credential.CredentialId, receivedCredential.CredentialId)
	require.Equal(t, credential.IssuerDID, receivedCredential.IssuerDID)
	require.Equal(t, credential.HolderDID, receivedCredential.HolderDID)
	require.Equal(t, credential.Prescription.DoctorId, receivedCredential.Prescription.DoctorId)
	require.Equal(t, credential.Prescription.RawPrescription, receivedCredential.Prescription.RawPrescription)
	require.Equal(t, credential.RawCredentialWithProof, receivedCredential.RawCredentialWithProof)
}

func TestGetNotExistCredentialByID(t *testing.T) {
	var credentialId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	vcStorage, err := NewVCStorage(dbPath)
	require.NoError(t, err)

	_, err = vcStorage.GetCredentialById(credentialId)
	require.Error(t, err)
}

func TestDoubleTimeSavePresentation(t *testing.T) {
	var presentation = domain.Presentation{
		PresentationId: tmrand.Str(6),
		HolderDID:      fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
		Type:           tmrand.Str(10),
		Credential: domain.Credential{
			CredentialId: tmrand.Str(6),
			IssuerDID:    fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
			HolderDID:    fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
			Prescription: domain.Prescription{
				DoctorId:        tmrand.Str(6),
				RawPrescription: []byte(`{"somePrescription":"someValue"}`),
			},
			RawCredentialWithProof: []byte(`{"someCredential":"someValue"}`),
		},
		RawPresentationWithProof: []byte(`{"somePresentation":"someValue"}`),
	}

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	vcStorage, err := NewVCStorage(dbPath)
	require.NoError(t, err)

	err = vcStorage.SavePresentation(presentation)
	require.NoError(t, err)

	err = vcStorage.SavePresentation(presentation)
	require.Error(t, err)
}

func TestGetExistPresentationByID(t *testing.T) {
	var presentation = domain.Presentation{
		PresentationId: tmrand.Str(6),
		HolderDID:      fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
		Type:           tmrand.Str(10),
		Credential: domain.Credential{
			CredentialId: tmrand.Str(6),
			IssuerDID:    fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
			HolderDID:    fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
			Prescription: domain.Prescription{
				DoctorId:        tmrand.Str(6),
				RawPrescription: []byte(`{"somePrescription":"someValue"}`),
			},
			RawCredentialWithProof: []byte(`{"someCredential":"someValue"}`),
		},
		RawPresentationWithProof: []byte(`{"somePresentation":"someValue"}`),
	}

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	vcStorage, err := NewVCStorage(dbPath)
	require.NoError(t, err)

	err = vcStorage.SavePresentation(presentation)
	require.NoError(t, err)

	receivedPresentation, err := vcStorage.GetPresentationById(presentation.PresentationId)
	require.NoError(t, err)
	require.Equal(t, presentation.PresentationId, receivedPresentation.PresentationId)
	require.Equal(t, presentation.HolderDID, receivedPresentation.HolderDID)
	require.Equal(t, presentation.Type, receivedPresentation.Type)
	require.Equal(t, presentation.Credential.CredentialId, receivedPresentation.Credential.CredentialId)
	require.Equal(t, presentation.Credential.IssuerDID, receivedPresentation.Credential.IssuerDID)
	require.Equal(t, presentation.Credential.HolderDID, receivedPresentation.Credential.HolderDID)
	require.Equal(t, presentation.Credential.Prescription.DoctorId, receivedPresentation.Credential.Prescription.DoctorId)
	require.Equal(t, presentation.Credential.Prescription.RawPrescription, receivedPresentation.Credential.Prescription.RawPrescription)
	require.Equal(t, presentation.Credential.RawCredentialWithProof, receivedPresentation.Credential.RawCredentialWithProof)
	require.Equal(t, presentation.RawPresentationWithProof, receivedPresentation.RawPresentationWithProof)
}

func TestGetNotExistPresentationByID(t *testing.T) {
	var presentationId = tmrand.Str(6)

	var dbPath = generateDBPath()
	defer cleanUp(dbPath)

	vcStorage, err := NewVCStorage(dbPath)
	require.NoError(t, err)

	_, err = vcStorage.GetPresentationById(presentationId)
	require.Error(t, err)
}
