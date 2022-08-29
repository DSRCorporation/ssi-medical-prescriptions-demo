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

package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/controller/mock"
	controllerRest "github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/controller/rest"
)

func TestGetExistCredentialByOfferID(t *testing.T) {
	client := resty.New()

	body, err := json.Marshal(testconstants.GetCredentialOfferResponseInfo.Prescription)
	require.NoError(t, err)

	var prescription controllerRest.Prescription
	require.NoError(t, json.Unmarshal(body, &prescription))

	var receivedCredentialID controllerRest.CredentialOfferResponse

	// Creates credential offer for prescription.
	resp, err := client.R().
		SetBody(prescription).
		SetResult(&receivedCredentialID).
		Post(fmt.Sprintf("%s/doctors/%s/prescriptions/credential-offers/",
			testconstants.Host, testconstants.DoctorID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var receivedCredentialOfferResponse controllerRest.GetCredentialOfferResponse

	// Get credential by offer ID.
	resp, err = client.R().
		SetResult(&receivedCredentialOfferResponse).
		Get(fmt.Sprintf("%s/doctors/%s/prescriptions/credential-offers/%s",
			testconstants.Host, testconstants.DoctorID, *receivedCredentialID.CredentialOfferId))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	require.True(t, reflect.DeepEqual(prescription, *receivedCredentialOfferResponse.Prescription))
}

func TestGetExistCredentialIssuedForGivenCredentialOffer(t *testing.T) {
	client := resty.New()

	body, err := json.Marshal(testconstants.GetCredentialOfferResponseInfo.Prescription)
	require.NoError(t, err)

	var prescription controllerRest.Prescription
	require.NoError(t, json.Unmarshal(body, &prescription))

	var receivedCredentialID controllerRest.CredentialOfferResponse

	// Creates credential offer for prescription.
	resp, err := client.R().
		SetBody(prescription).
		SetResult(&receivedCredentialID).
		Post(fmt.Sprintf("%s/doctors/%s/prescriptions/credential-offers/",
			testconstants.Host, testconstants.DoctorID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var patientUserName = tmrand.Str(12)
	var patientPassword = tmrand.Str(8)

	patientAuthInfo := controllerRest.PatientAuthCredential{
		Username: &patientUserName,
		Password: &patientPassword,
	}
	registerationPatientResponse := struct {
		PatientId string `json:"patientId"`
		DID       string `json:"did"`
	}{}

	// Registration patient
	resp, err = client.R().
		SetBody(patientAuthInfo).
		SetResult(&registerationPatientResponse).
		Post(fmt.Sprintf("%s/patients/register", testconstants.Host))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	patientsPatientIdPrescriptionsCredentials := controllerRest.PostV1PatientsPatientIdPrescriptionsCredentialsJSONBody{
		CredentialOfferId: receivedCredentialID.CredentialOfferId,
		Did:               &registerationPatientResponse.DID,
		KmsPassphrase:     patientAuthInfo.Password,
	}
	var credential Credential

	// Creates credential in response to credential offer from doctor.
	resp, err = client.R().
		SetBody(patientsPatientIdPrescriptionsCredentials).
		SetResult(&credential).
		Post(fmt.Sprintf("%s/patients/%s/prescriptions/credentials/",
			testconstants.Host, registerationPatientResponse.PatientId))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var receivedCredential Credential

	// Gets credential issued for given credential offer.
	resp, err = client.R().
		SetResult(&receivedCredential).
		Get(fmt.Sprintf("%s/doctors/%s/prescriptions/credential-offers/%s/credential",
			testconstants.Host, testconstants.DoctorID, *receivedCredentialID.CredentialOfferId))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())
	require.True(t, reflect.DeepEqual(credential, receivedCredential))
}

type Credential struct {
	Credential *json.RawMessage `json:"credential"`
}
