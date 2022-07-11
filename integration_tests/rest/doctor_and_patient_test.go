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

	var receivedCredentialOfferResponse controllerRest.GetCredentialOfferResponse

	// Get credential by offer ID.
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&receivedCredentialOfferResponse).
		Get(fmt.Sprintf("http://localhost:8989/v1/doctors/%s/prescriptions/credential-offers/%s",
			testconstants.DoctorID, *receivedCredentialID.CredentialOfferId))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())
	require.True(t, reflect.DeepEqual(prescription, receivedCredentialOfferResponse))
}

func TestGetExistCredentialIssuedForGivenCredentialOffer(t *testing.T) {
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
			credentialOfferId: *receivedCredentialID.CredentialOfferId,
			did:               fmt.Sprintf("cheqd:testnet:%s", tmrand.Str(32)),
			kmsPassphrase:     tmrand.Str(16),
		}).
		SetResult(&credential).
		Post(fmt.Sprintf("http://localhost:8989/v1/patients/%s/prescriptions/credentials/", testconstants.PatientID))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var receivedCredential controllerRest.Credential

	// Gets credential issued for given credential offer.
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&receivedCredential).
		Get(fmt.Sprintf("http://localhost:8989/v1/doctors/%s/prescriptions/credential-offers/%s/credential",
			testconstants.DoctorID, *receivedCredentialID.CredentialOfferId))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())
	require.True(t, reflect.DeepEqual(credential, receivedCredential))
}
