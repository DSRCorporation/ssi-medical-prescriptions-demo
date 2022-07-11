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

func TestGetExistPresentationRequestForPrescriptionByRequestID(t *testing.T) {
	client := resty.New()

	body, err := json.Marshal(testconstants.GetCredentialOfferResponseInfo)
	require.NoError(t, err)

	var prescription controllerRest.GetCredentialOfferResponse
	require.NoError(t, json.Unmarshal(body, &prescription))

	var receivedCredentialID controllerRest.CredentialOfferResponse

	// Creates credential offer for prescription.
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(prescription).
		SetResult(&receivedCredentialID).
		Post(fmt.Sprintf("http://localhost:8989/v1/doctors/%s/prescriptions/credential-offers/", testconstants.DoctorID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var credential controllerRest.Credential

	// Creates credential in response to credential offer from doctor.
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(struct {
			credentialOfferId string
			did               string
			kmsPassphrase     string
		}{
			credentialOfferId: tmrand.Str(6),
			did:               fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
			kmsPassphrase:     tmrand.Str(16),
		}).
		SetResult(&credential).
		Post(fmt.Sprintf("http://localhost:8989/v1/patients/%s/prescriptions/credentials/", testconstants.PatientID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var createPresentationRequestResponse controllerRest.CreatePresentationRequestResponse

	// Creates presentation request for prescriptions.
	resp, err = client.R().
		SetHeader("ContentType", "application/json").
		SetResult(&createPresentationRequestResponse).
		Post(fmt.Sprintf("http://localhost:8989/v1/pharmacies/%s/prescriptions/presentation-requests", testconstants.PharmacyID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var receivedPresentationRequestResponse controllerRest.GetPresentationRequestResponse

	// Get presentation request for prescription by request ID.
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&receivedPresentationRequestResponse).
		Get(fmt.Sprintf("http://localhost:8989/v1/pharmacies/%s/prescriptions/presentation-requests/%s",
			testconstants.PharmacyID, *createPresentationRequestResponse.PresentationRequestId))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())
	require.Equal(t, *createPresentationRequestResponse.PresentationRequestId, *receivedPresentationRequestResponse.PresentationRequestId)
}

func TestGetExistVerifiablePresentationForGivenPresentationRequest(t *testing.T) {
	client := resty.New()

	body, err := json.Marshal(testconstants.GetCredentialOfferResponseInfo)
	require.NoError(t, err)

	var prescription controllerRest.GetCredentialOfferResponse
	require.NoError(t, json.Unmarshal(body, &prescription))

	var receivedCredentialID controllerRest.CredentialOfferResponse
	// Creates credential offer for prescription.
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(prescription).
		SetResult(&receivedCredentialID).
		Post(fmt.Sprintf("http://localhost:8989/v1/doctors/%s/prescriptions/credential-offers/", testconstants.DoctorID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var credential controllerRest.Credential

	// Creates credential in response to credential offer from doctor.
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(struct {
			credentialOfferId string
			did               string
			kmsPassphrase     string
		}{
			credentialOfferId: tmrand.Str(6),
			did:               fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
			kmsPassphrase:     tmrand.Str(16),
		}).
		SetResult(&credential).
		Post(fmt.Sprintf("http://localhost:8989/v1/patients/%s/prescriptions/credentials/", testconstants.PatientID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var createPresentationRequestResponse controllerRest.CreatePresentationRequestResponse

	// Creates presentation request for prescriptions.
	resp, err = client.R().
		SetHeader("ContentType", "application/json").
		SetResult(&createPresentationRequestResponse).
		Post(fmt.Sprintf("http://localhost:8989/v1/pharmacies/%s/prescriptions/presentation-requests", testconstants.PharmacyID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var presentation controllerRest.Presentation

	// Creates verifiable presentation in response to prescription presentation request from pharmacy.
	resp, err = client.R().
		SetHeader("ContentType", "application/json").
		SetBody(struct {
			presentationRequestId string
			credentialId          string
			challenge             string
			kmsPassphrase         string
		}{
			presentationRequestId: *createPresentationRequestResponse.PresentationRequestId,
			credentialId:          tmrand.Str(6),
			challenge:             tmrand.Str(10),
			kmsPassphrase:         tmrand.Str(16),
		}).
		SetResult(&presentation).
		Post(fmt.Sprintf("http://localhost:8989/v1/patients/%s/prescriptions/presentations", testconstants.PatientID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var receivedVerifiablePresentationResponse controllerRest.GetVerifiablePresentationResponse

	// Gets verifiable presentation for given presentation request.
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&receivedVerifiablePresentationResponse).
		Get(fmt.Sprintf("http://localhost:8989/v1/pharmacies/%s/prescriptions/presentation-requests/%s/presentation",
			testconstants.PharmacyID, *createPresentationRequestResponse.PresentationRequestId))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())
	require.True(t, reflect.DeepEqual(presentation, *receivedVerifiablePresentationResponse.Presentation))
	require.Equal(t, "Verification Successful", *receivedVerifiablePresentationResponse.VerificationComment)
	require.True(t, *receivedVerifiablePresentationResponse.VerificationStatus)
}
