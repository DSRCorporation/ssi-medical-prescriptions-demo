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
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/controller/rest"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
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

func New(doctorService *service.DoctorService, patientService *service.PatientService, pharmacyService *service.PharmacyService, vcService *service.VCService) *RestHandler {
	return &RestHandler{
		doctorService:   doctorService,
		patientService:  patientService,
		pharmacyService: pharmacyService,
		vcService:       vcService,
	}
}

// Creates credential offer for prescription (generates a link for QR code)
// (POST /v1/doctors/{doctorId}/prescriptions/credential-offers/)
func (h *RestHandler) PostV1DoctorsDoctorIdPrescriptionsCredentialOffers(ctx echo.Context, doctorId string) error {
	var body rest.Prescription
	ctx.Bind(&body)

	prescription, err := ConvertToPrescription(body, doctorId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	offerId := uuid.New().String()
	err = h.doctorService.CreatePrescriptionOffer(offerId, *prescription)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := rest.CredentialOfferResponse{
		CredentialOfferId: &offerId,
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
	prescription, err := h.doctorService.GetPrescriptionByOfferId(credentialOfferId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	fromPrescription, err := ConvertFromPrescription(prescription)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := rest.GetCredentialOfferResponse{
		Prescription: fromPrescription,
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
		credentialId, err := h.doctorService.GetCredentialIdByOfferId(credentialOfferId)
		if err == nil {
			credential, err := h.vcService.GetCredentialById(credentialId)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			response := struct {
				Credential json.RawMessage `json:"credential"`
			}{
				Credential: credential.RawCredential,
			}
			return ctx.JSON(http.StatusOK, response)
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
	credentialIds, err := h.patientService.GetCredentialIdsByPatientId(patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var rawCredentials []json.RawMessage
	for _, credentialId := range credentialIds {
		credential, err := h.vcService.GetCredentialById(credentialId)
		if err == nil {
			rawCredentials = append(rawCredentials, credential.RawCredential)
		}
	}

	response := struct {
		Credentials *[]json.RawMessage `json:"credentials"`
	}{Credentials: &rawCredentials}

	err = ctx.JSON(http.StatusOK, response)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Creates credential in response to credential offer from doctor
// (POST /v1/patients/{patientId}/prescriptions/credentials/)
func (h *RestHandler) PostV1PatientsPatientIdPrescriptionsCredentials(ctx echo.Context, patientId string) error {
	var body rest.PostV1PatientsPatientIdPrescriptionsCredentialsJSONBody
	ctx.Bind(&body)

	credentialOfferId := *body.CredentialOfferId
	patientDID := body.Did
	patientKmsPassphrase := body.KmsPassphrase

	prescription, err := h.doctorService.GetPrescriptionByOfferId(credentialOfferId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	doctorId := prescription.DoctorId
	doctorKmsPassphrase, err := h.doctorService.GetKMSPassphrase(doctorId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	doctorDID, err := h.doctorService.GetDID(doctorId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	credentialId := domain.GenerateVerifiableId()
	unsignedCredential := domain.NewCredential(credentialId, doctorDID, *patientDID, prescription, nil)

	signedCredential, err := h.vcService.ExchangeCredential(doctorId, doctorKmsPassphrase, patientId, *patientKmsPassphrase, *unsignedCredential)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = h.doctorService.SaveCredentialId(doctorId, credentialOfferId, signedCredential.CredentialId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = h.patientService.SaveCredentialId(patientId, signedCredential.CredentialId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := struct {
		Credential *json.RawMessage `json:"credential"`
	}{Credential: &signedCredential.RawCredential}

	err = ctx.JSON(http.StatusOK, response)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Gets prescription credential by id issued for given patient
// (GET /v1/patients/{patientId}/prescriptions/credentials/{credentialId})
func (h *RestHandler) GetV1PatientsPatientIdPrescriptionsCredentialsCredentialId(ctx echo.Context, patientId string, credentialId string) error {
	credentialIds, err := h.patientService.GetCredentialIdsByPatientId(patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for _, credId := range credentialIds {
		if credId == credentialId {
			credential, err := h.vcService.GetCredentialById(credentialId)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			response := struct {
				Credential *json.RawMessage `json:"credential"`
			}{Credential: &credential.RawCredential}

			err = ctx.JSON(http.StatusOK, response)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			break
		}
	}

	return nil
}

// Creates verifiable presentation in response to prescription presentation request from pharmacy
// (POST /v1/patients/{patientId}/prescriptions/presentations/)
func (h *RestHandler) PostV1PatientsPatientIdPrescriptionsPresentations(ctx echo.Context, patientId string) error {
	var body rest.PostV1PatientsPatientIdPrescriptionsPresentationsJSONBody
	ctx.Bind(&body)

	requestId := *body.PresentationRequestId
	credentialId := *body.CredentialId

	credential, err := h.vcService.GetCredentialById(credentialId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pharmacyId, err := h.pharmacyService.GetPharmacyIdByRequestId(requestId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	patientDID := credential.HolderDID
	patientKmsPassphrase := body.KmsPassphrase

	presentationId := domain.GenerateVerifiableId()
	unsignedPresentation := domain.NewPresentation(presentationId, patientDID, credential, nil)

	signedPresentation, err := h.vcService.ExchangePresentation(pharmacyId, patientId, *patientKmsPassphrase, *unsignedPresentation)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = h.pharmacyService.SavePresentationId(pharmacyId, requestId, signedPresentation.PresentationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := struct {
		Presentation json.RawMessage `json:"presentation"`
	}{
		Presentation: signedPresentation.RawPresentation,
	}

	err = ctx.JSON(http.StatusCreated, response)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Creates presentation request for prescription (generates link for a QR code)
// (POST /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests)
func (h *RestHandler) PostV1PharmaciesPharmacyIdPrescriptionsPresentationRequests(ctx echo.Context, pharmacyId string) error {
	requestId := uuid.New().String()
	err := h.pharmacyService.CreatePresentationRequest(pharmacyId, requestId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := rest.CreatePresentationRequestResponse{
		PresentationRequestId: &requestId,
	}

	err = ctx.JSON(http.StatusCreated, response)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

// Gets presentation request for prescription by request id
// (GET /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId})
func (h *RestHandler) GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestId(ctx echo.Context, pharmacyId string, presentationRequestId string) error {
	pharmacyIdByRequestId, err := h.pharmacyService.GetPharmacyIdByRequestId(presentationRequestId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if pharmacyIdByRequestId != pharmacyId {
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid presentation request id")
	}

	response := rest.GetPresentationRequestResponse{
		PresentationRequestId: &presentationRequestId,
	}

	err = ctx.JSON(http.StatusOK, response)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Gets verifiable presentation for given presentation request
// (GET /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId}/presentation)
func (h *RestHandler) GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestIdPresentation(ctx echo.Context, pharmacyId string, presentationRequestId string) error {
	// long polling
	for i := 0; i < LONG_POLLING_TIMEOUT_SECONDS; i++ {
		presentationId, err := h.pharmacyService.GetPresentaionIdByRequestId(presentationRequestId)
		if err == nil {
			presentation, err := h.vcService.GetPresentationById(presentationId)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			response := struct {
				Presentation        *json.RawMessage `json:"presentation"`
				VerificationComment string           `json:"verificationComment"`
				VerificationStatus  bool             `json:"verificationStatus"`
			}{
				Presentation: &presentation.RawPresentation,
				// @TODO: verify presentation is valid
				VerificationComment: "Verification successful",
				VerificationStatus:  true,
			}
			return ctx.JSON(http.StatusOK, response)
		}
		time.Sleep(time.Second)
	}
	return echo.NewHTTPError(http.StatusInternalServerError, "Presentation not found")
}

// Verify Credential
// (POST /v1/vc/verify-credential)
func (h *RestHandler) PostV1VcVerifyCredential(ctx echo.Context) error {
	rawCredential, err := ioutil.ReadAll(ctx.Request().Body) // c <= echo.Context
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer ctx.Request().Body.Close()

	err = h.vcService.VerifyCredential(rawCredential)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	verified := true
	err = ctx.JSON(http.StatusOK, rest.VerificationResponse{Verified: &verified})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

// Verify Credential
// (POST /v1/vc/verify-presentation)
func (h *RestHandler) PostV1VcVerifyPresentation(ctx echo.Context) error {
	rawPresentation, err := ioutil.ReadAll(ctx.Request().Body) // c <= echo.Context
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer ctx.Request().Body.Close()

	err = h.vcService.VerifyPresentation(rawPresentation)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	verified := true
	err = ctx.JSON(http.StatusOK, rest.VerificationResponse{Verified: &verified})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
