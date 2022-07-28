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
	testconstants "github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/controller/mock"
	controllerRest "github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/controller/rest"
)

func TestGetExistPresentationRequestForPrescriptionByRequestID(t *testing.T) {
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

	patientsPatientIdPrescriptionsCredentials := controllerRest.PostV1PatientsPatientIdPrescriptionsCredentialsJSONBody{
		CredentialOfferId: receivedCredentialID.CredentialOfferId,
		Did:               &testconstants.PatientDID,
		KmsPassphrase:     &testconstants.PatientKMSPassphrase,
	}

	var credential Credential

	// Creates credential in response to credential offer from doctor.
	resp, err = client.R().
		SetBody(patientsPatientIdPrescriptionsCredentials).
		SetResult(&credential).
		Post(fmt.Sprintf("%s/patients/%s/prescriptions/credentials/",
			testconstants.Host, testconstants.PatientID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var createPresentationRequestResponse controllerRest.CreatePresentationRequestResponse

	// Creates presentation request for prescriptions.
	resp, err = client.R().
		SetResult(&createPresentationRequestResponse).
		Post(fmt.Sprintf("%s/pharmacies/%s/prescriptions/presentation-requests",
			testconstants.Host, testconstants.PharmacyID))
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode())

	var receivedPresentationRequestResponse controllerRest.GetPresentationRequestResponse

	// Get presentation request for prescription by request ID.
	resp, err = client.R().
		SetResult(&receivedPresentationRequestResponse).
		Get(fmt.Sprintf("%s/pharmacies/%s/prescriptions/presentation-requests/%s",
			testconstants.Host, testconstants.PharmacyID, *createPresentationRequestResponse.PresentationRequestId))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())
	require.Equal(t, *createPresentationRequestResponse.PresentationRequestId, *receivedPresentationRequestResponse.PresentationRequestId)
}

func TestGetExistVerifiablePresentationForGivenPresentationRequest(t *testing.T) {
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

	patientsPatientIdPrescriptionsCredentials := controllerRest.PostV1PatientsPatientIdPrescriptionsCredentialsJSONBody{
		CredentialOfferId: receivedCredentialID.CredentialOfferId,
		Did:               &testconstants.PatientDID,
		KmsPassphrase:     &testconstants.PatientKMSPassphrase,
	}

	var credential Credential

	// Creates credential in response to credential offer from doctor.
	resp, err = client.R().
		SetBody(patientsPatientIdPrescriptionsCredentials).
		SetResult(&credential).
		Post(fmt.Sprintf("%s/patients/%s/prescriptions/credentials/",
			testconstants.Host, testconstants.PatientID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var createPresentationRequestResponse controllerRest.CreatePresentationRequestResponse

	// Creates presentation request for prescriptions.
	resp, err = client.R().
		SetResult(&createPresentationRequestResponse).
		Post(fmt.Sprintf("%s/pharmacies/%s/prescriptions/presentation-requests",
			testconstants.Host, testconstants.PharmacyID))
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode())

	var receivedCredential controllerRest.Credential
	err = json.Unmarshal(*credential.Credential, &receivedCredential)
	require.NoError(t, err)

	patientIdPrescriptionsPresentations := controllerRest.PostV1PatientsPatientIdPrescriptionsPresentationsJSONBody{
		PresentationRequestId: createPresentationRequestResponse.PresentationRequestId,
		CredentialId:          receivedCredential.Id,
		KmsPassphrase:         &testconstants.PatientKMSPassphrase,
	}
	var presentation verifiablePresentationResponse

	// Creates verifiable presentation in response to prescription presentation request from pharmacy.
	resp, err = client.R().
		SetBody(patientIdPrescriptionsPresentations).
		SetResult(&presentation).
		Post(fmt.Sprintf("%s/patients/%s/prescriptions/presentations",
			testconstants.Host, testconstants.PatientID))
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode())

	var receivedVerifiablePresentationResponse verifiablePresentationResponse

	// Gets verifiable presentation for given presentation request.
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&receivedVerifiablePresentationResponse).
		Get(fmt.Sprintf("%s/pharmacies/%s/prescriptions/presentation-requests/%s/presentation",
			testconstants.Host, testconstants.PharmacyID, *createPresentationRequestResponse.PresentationRequestId))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())
	require.True(t, reflect.DeepEqual(presentation.Presentation, receivedVerifiablePresentationResponse.Presentation))
	require.Equal(t, "Verification successful", receivedVerifiablePresentationResponse.VerificationComment)
	require.True(t, receivedVerifiablePresentationResponse.VerificationStatus)
}

type verifiablePresentationResponse struct {
	Presentation        *json.RawMessage `json:"presentation"`
	VerificationComment string           `json:"verificationComment"`
	VerificationStatus  bool             `json:"verificationStatus"`
}
