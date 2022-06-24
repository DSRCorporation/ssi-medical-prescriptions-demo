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

package mock

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SSIMPMockImpl struct{}

type ErrorImpl struct {
	msg string
}

func (e *ErrorImpl) New(msg string) {
	e.msg = msg
}

func (e *ErrorImpl) Error() string {
	return e.msg
}

// Creates credential offer for prescription (generates a link for QR code)
// (POST /v1/doctors/{doctorId}/prescriptions/credential-offers/)
func (*SSIMPMockImpl) PostV1DoctorsDoctorIdPrescriptionsCredentialOffers(ctx echo.Context, doctorId string) error {
	response := CredentialOfferResponseInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets credential offer by id
// (GET /v1/doctors/{doctorId}/prescriptions/credential-offers/{credentialOfferId})
func (*SSIMPMockImpl) GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferId(ctx echo.Context, doctorId string, credentialOfferId string) error {
	response := GetCredentialOfferResponseInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets credential issued for given credential offer
// (GET /v1/doctors/{doctorId}/prescriptions/credential-offers/{credentialOfferId}/credential)
func (*SSIMPMockImpl) GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferIdCredential(ctx echo.Context, doctorId string, credentialOfferId string) error {
	response := CredentialResponseInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets all dids belonging to given patient
// (GET /v1/patients/{patientId}/dids)
func (*SSIMPMockImpl) GetV1PatientsPatientIdDids(ctx echo.Context, patientId string) error {
	response := GetAllDidsResponseInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets all prescription credentials issued for given patient
// (GET /v1/patients/{patientId}/prescriptions/credentials)
func (*SSIMPMockImpl) GetV1PatientsPatientIdPrescriptionsCredentials(ctx echo.Context, patientId string) error {
	response := GetAllPrescriptionCredentialResponseInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Creates credential in response to credential offer from doctor
// (POST /v1/patients/{patientId}/prescriptions/credentials/)
func (*SSIMPMockImpl) PostV1PatientsPatientIdPrescriptionsCredentials(ctx echo.Context, patientId string) error {
	response := CredentialResponseInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets prescription credential by id issued for given patient
// (GET /v1/patients/{patientId}/prescriptions/credentials/{credentialId})
func (*SSIMPMockImpl) GetV1PatientsPatientIdPrescriptionsCredentialsCredentialId(ctx echo.Context, patientId string, credentialId string) error {
	response := GetPrescriptionCredentialResponseInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Creates verifiable presentation in response to prescription presentation request from pharmacy
// (POST /v1/patients/{patientId}/prescriptions/presentations/)
func (*SSIMPMockImpl) PostV1PatientsPatientIdPrescriptionsPresentations(ctx echo.Context, patientId string) error {
	response := PresentationInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Creates presentation request for prescription (generates link for a QR code)
// (POST /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests)
func (*SSIMPMockImpl) PostV1PharmaciesPharmacyIdPrescriptionsPresentationRequests(ctx echo.Context, pharmacyId string) error {
	response := CreatePresentationRequestResponseInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets presentation request for prescription by request id
// (GET /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId})
func (*SSIMPMockImpl) GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestId(ctx echo.Context, pharmacyId string, presentationRequestId string) error {
	response := GetPresentationRequestResponseInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets verifiable presentation for given presentation request
// (GET /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId}/presentation)
func (*SSIMPMockImpl) GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestIdPresentation(ctx echo.Context, pharmacyId string, presentationRequestId string) error {
	response := GetVerifiablePresentationResponseInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Verify Credential
// (POST /v1/vc/verify-credential)
func (*SSIMPMockImpl) PostV1VcVerifyCredential(ctx echo.Context) error {
	response := CredentialResponseInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Verify Credential
// (POST /v1/vc/verify-presentation)
func (*SSIMPMockImpl) PostV1VcVerifyPresentation(ctx echo.Context) error {
	response := PresentationInfo

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}
