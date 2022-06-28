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

package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/controller/rest"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/service"
)

const (
	LONG_POLLING_TIMEOUT_SECONDS = 60 * 5 // 5 minutes
)

type RestHandler struct {
	doctorService   *service.DoctorService
	patientService  *service.PatientService
	pharmacyService *service.PharmacyService
	vcService       *service.VCService
}

// Creates credential offer for prescription (generates a link for QR code)
// (POST /v1/doctors/{doctorId}/prescriptions/credential-offers/)
func (h *RestHandler) PostV1DoctorsDoctorIdPrescriptionsCredentialOffers(ctx echo.Context, doctorId string) error {
	var in rest.Prescription
	ctx.Bind(in)

	prescription := ConvertToPrescription(in)

	credentialOfferId, err := h.doctorService.CreatePrescriptionCredentialOffer(doctorId, prescription)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := rest.CredentialOfferResponse{
		CredentialOfferId: &credentialOfferId,
	}

	err = ctx.JSON(http.StatusOK, response)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Gets credential offer by id
// (GET /v1/doctors/{doctorId}/prescriptions/credential-offers/{credentialOfferId})
func (h *RestHandler) GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferId(ctx echo.Context, doctorId string, credentialOfferId string) error {
	prescription, err := h.doctorService.GetPrescriptionByCredentialOfferId(doctorId, credentialOfferId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	fromPrescription := ConvertFromPrescription(prescription)
	response := rest.GetCredentialOfferResponse{
		Prescription: &fromPrescription,
	}

	err = ctx.JSON(http.StatusOK, response)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Gets credential issued for given credential offer
// (GET /v1/doctors/{doctorId}/prescriptions/credential-offers/{credentialOfferId}/credential)
func (h *RestHandler) GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferIdCredential(ctx echo.Context, doctorId string, credentialOfferId string) error {
	// long polling
	for i := 0; i < LONG_POLLING_TIMEOUT_SECONDS; i++ {
		credential, err := h.doctorService.GetPrescriptionCredentialByCredentialOfferId(doctorId, credentialOfferId)

		if err == nil {
			return ctx.JSONBlob(http.StatusOK, credential.RawCredential)
		}

		time.Sleep(time.Second)
	}
	return echo.NewHTTPError(http.StatusInternalServerError, "Credential not found")
}

// Gets all dids belonging to given patient
// (GET /v1/patients/{patientId}/dids)
func (h *RestHandler) GetV1PatientsPatientIdDids(ctx echo.Context, patientId string) error {
	dids, err := h.patientService.GetDIDs(patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := rest.GetAllDidsResponse{
		Dids: &dids,
	}

	err = ctx.JSON(http.StatusOK, response)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Gets all prescription credentials issued for given patient
// (GET /v1/patients/{patientId}/prescriptions/credentials)
func (h *RestHandler) GetV1PatientsPatientIdPrescriptionsCredentials(ctx echo.Context, patientId string) error {
	return nil
}

// Creates credential in response to credential offer from doctor
// (POST /v1/patients/{patientId}/prescriptions/credentials/)
func (*RestHandler) PostV1PatientsPatientIdPrescriptionsCredentials(ctx echo.Context, patientId string) error {
	return nil
}

// Gets prescription credential by id issued for given patient
// (GET /v1/patients/{patientId}/prescriptions/credentials/{credentialId})
func (*RestHandler) GetV1PatientsPatientIdPrescriptionsCredentialsCredentialId(ctx echo.Context, patientId string, credentialId string) error {
	return nil
}

// Creates verifiable presentation in response to prescription presentation request from pharmacy
// (POST /v1/patients/{patientId}/prescriptions/presentations/)
func (*RestHandler) PostV1PatientsPatientIdPrescriptionsPresentations(ctx echo.Context, patientId string) error {
	return nil
}

// Creates presentation request for prescription (generates link for a QR code)
// (POST /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests)
func (*RestHandler) PostV1PharmaciesPharmacyIdPrescriptionsPresentationRequests(ctx echo.Context, pharmacyId string) error {
	return nil
}

// Gets presentation request for prescription by request id
// (GET /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId})
func (*RestHandler) GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestId(ctx echo.Context, pharmacyId string, presentationRequestId string) error {
	return nil
}

// Gets verifiable presentation for given presentation request
// (GET /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId}/presentation)
func (*RestHandler) GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestIdPresentation(ctx echo.Context, pharmacyId string, presentationRequestId string) error {
	return nil
}

// Verify Credential
// (POST /v1/vc/verify-credential)
func (*RestHandler) PostV1VcVerifyCredential(ctx echo.Context) error {
	return nil
}

// Verify Credential
// (POST /v1/vc/verify-presentation)
func (*RestHandler) PostV1VcVerifyPresentation(ctx echo.Context) error {
	return nil
}
