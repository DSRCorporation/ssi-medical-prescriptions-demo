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
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/constants"
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
	response := CredentialOfferResponse{
		CredentialOfferId: &constants.CredentialOfferResponse,
	}

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets credential offer by id
// (GET /v1/doctors/{doctorId}/prescriptions/credential-offers/{credentialOfferId})
func (*SSIMPMockImpl) GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferId(ctx echo.Context, doctorId string, credentialOfferId string) error {
	response := getCredentialOfferResponse

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets credential issued for given credential offer
// (GET /v1/doctors/{doctorId}/prescriptions/credential-offers/{credentialOfferId}/credential)
func (*SSIMPMockImpl) GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferIdCredential(ctx echo.Context, doctorId string, credentialOfferId string) error {
	response := credentialResponse

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets all dids belonging to given patient
// (GET /v1/patients/{patientId}/dids)
func (*SSIMPMockImpl) GetV1PatientsPatientIdDids(ctx echo.Context, patientId string) error {
	response := getAllDidsResponse

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets all prescription credentials issued for given patient
// (GET /v1/patients/{patientId}/prescriptions/credentials)
func (*SSIMPMockImpl) GetV1PatientsPatientIdPrescriptionsCredentials(ctx echo.Context, patientId string) error {
	response := getAllPrescriptionCredentialResponse

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Creates credential in response to credential offer from doctor
// (POST /v1/patients/{patientId}/prescriptions/credentials/)
func (*SSIMPMockImpl) PostV1PatientsPatientIdPrescriptionsCredentials(ctx echo.Context, patientId string) error {
	response := credentialResponse

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets prescription credential by id issued for given patient
// (GET /v1/patients/{patientId}/prescriptions/credentials/{credentialId})
func (*SSIMPMockImpl) GetV1PatientsPatientIdPrescriptionsCredentialsCredentialId(ctx echo.Context, patientId string, credentialId string) error {
	response := getPrescriptionCredentialResponse

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Creates verifiable presentation in response to prescription presentation request from pharmacy
// (POST /v1/patients/{patientId}/prescriptions/presentations/)
func (*SSIMPMockImpl) PostV1PatientsPatientIdPrescriptionsPresentations(ctx echo.Context, patientId string) error {
	response := presentation

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Creates presentation request for prescription (generates link for a QR code)
// (POST /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests)
func (*SSIMPMockImpl) PostV1PharmaciesPharmacyIdPrescriptionsPresentationRequests(ctx echo.Context, pharmacyId string) error {
	response := createPresentationRequestResponse

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets presentation request for prescription by request id
// (GET /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId})
func (*SSIMPMockImpl) GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestId(ctx echo.Context, pharmacyId string, presentationRequestId string) error {
	response := getPresentationRequestResponse

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Gets verifiable presentation for given presentation request
// (GET /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId}/presentation)
func (*SSIMPMockImpl) GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestIdPresentation(ctx echo.Context, pharmacyId string, presentationRequestId string) error {
	response := getVerifiablePresentationResponse

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Verify Credential
// (POST /v1/vc/verify-credential)
func (*SSIMPMockImpl) PostV1VcVerifyCredential(ctx echo.Context) error {
	response := credentialResponse

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Verify Credential
// (POST /v1/vc/verify-presentation)
func (*SSIMPMockImpl) PostV1VcVerifyPresentation(ctx echo.Context) error {
	response := presentation

	err := ctx.JSON(http.StatusOK, response)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

// Ointment drug info
var ointmentDrugInfo interface{} = OintmentDrugInfo{
	AmountOfMedicine:     &constants.AmountOfMedicine,
	UseArea:              &constants.UseArea,
	ApplicationFrequency: &constants.ApplicationFrequency,
	Usage:                &constants.Usage,
}

// Tablet drug info
var tabletDrugInfo interface{} = TabletDrugInfo{
	NumberOfDrug:         &constants.NumberOfDrug,
	NumberOfDoses:        &constants.NumberOfDoses,
	DaysOfMedication:     &constants.DaysOfMedication,
	TimingToTakeMedicine: &constants.TimingToTakeMedicine,
}

// Get credential offer response info
var getCredentialOfferResponse = GetCredentialOfferResponse{
	Prescription: &Prescription{
		PatientInfo: &struct {
			Birthday *string "json:\"birthday,omitempty\""
			Name     *string "json:\"name,omitempty\""
			Sex      *string "json:\"sex,omitempty\""
		}{
			Birthday: &constants.PatientBirthday,
			Name:     &constants.PatientName,
			Sex:      &constants.PatientSex,
		},
		IssuanceInfo: &struct {
			ExpirationDate *string "json:\"expirationDate,omitempty\""
			IssuanceDate   *string "json:\"issuanceDate,omitempty\""
		}{
			IssuanceDate:   &constants.IssuanceDate,
			ExpirationDate: &constants.ExpirationDate,
		},
		HospitalInfo: &struct {
			DoctorName               *string  "json:\"doctorName,omitempty\""
			DoctorSignatureStamp     *string  "json:\"doctorSignatureStamp,omitempty\""
			HospitalName             *string  "json:\"hospitalName,omitempty\""
			Location                 *string  "json:\"location,omitempty\""
			MedicalInstitutionNumber *float32 "json:\"medicalInstitutionNumber,omitempty\""
			Phone                    *string  "json:\"phone,omitempty\""
			PrefectureNumber         *float32 "json:\"prefectureNumber,omitempty\""
			ScoreVoteNumber          *float32 "json:\"scoreVoteNumber,omitempty\""
		}{
			DoctorName:               &constants.DoctorName,
			DoctorSignatureStamp:     &constants.DoctorSignatureStamp,
			HospitalName:             &constants.HospitalName,
			Location:                 &constants.Location,
			MedicalInstitutionNumber: &constants.MedicalInstitutionNumber,
			Phone:                    &constants.Phone,
			PrefectureNumber:         &constants.PrefectureNumber,
			ScoreVoteNumber:          &constants.ScoreVoteNumber,
		},
		Drugs: &[]Drug{
			{
				DrugName:           &constants.DrugName,
				DrugNumber:         &constants.DrugNumber,
				DrugType:           &constants.DrugType,
				Info:               &ointmentDrugInfo,
				RefillAvailability: &constants.RefillAvailability,
			},
			{
				DrugName:           &constants.DrugName,
				DrugNumber:         &constants.DrugNumber,
				DrugType:           &constants.DrugType,
				Info:               &tabletDrugInfo,
				RefillAvailability: &constants.RefillAvailability,
			},
		},
	},
	Challenge: &constants.Challenge,
}

// Credential response info
var credentialResponse = Credential{
	Context: &[]string{
		constants.Context1,
		constants.Context2,
		constants.Context3,
	},
	Type: &[]string{
		constants.Type1,
		constants.Type2,
	},
	Issuer: &struct {
		Id   *string "json:\"id,omitempty\""
		Name *string "json:\"name,omitempty\""
	}{
		Id:   &constants.ID,
		Name: &constants.DoctorName,
	},
	IssuanceDate:   &constants.IssuanceDate,
	ExpirationDate: &constants.ExpirationDate,
	CredentialSubject: &struct {
		Id           *string       "json:\"id,omitempty\""
		Name         *string       "json:\"name,omitempty\""
		Prescription *Prescription "json:\"prescription,omitempty\""
	}{
		Id:           &constants.PatientID,
		Name:         &constants.PatientName,
		Prescription: getCredentialOfferResponse.Prescription,
	},
	Proof: &CredentialProof{
		Created:            &constants.Created,
		Jws:                &constants.JWS,
		ProofPurpose:       &constants.ProofPurpose,
		Type:               &constants.Type,
		VerificationMethod: &constants.VerificationMethod,
	},
}

// Get all DIDs response info
var getAllDidsResponse = GetAllDidsResponse{
	Dids: &[]string{
		constants.DIDs1,
		constants.DIDs2,
		constants.DIDs3,
	},
}

// Get all prescription credential response info
var getAllPrescriptionCredentialResponse = GetAllPrescriptionCredentialResponse{
	Credentials: &[]map[string]interface{}{
		{
			"first": credentialResponse,
		},
	},
}

// Get prescription credential response info
var getPrescriptionCredentialResponse = GetPrescriptionCredentialResponse{
	Credentials: &map[string]interface{}{"first": credentialResponse},
}

// Presentation info
var presentation = Presentation{
	Context: &[]string{
		constants.Context1,
		constants.Context2,
		constants.Context3,
	},
	Type: &constants.Type1,
	VerifiableCredential: &[]Credential{
		credentialResponse,
	},
	Proof: &[]PresentationProof{
		{
			Type:               &constants.Type,
			Created:            &constants.Created,
			ProofPurpose:       &constants.ProofPurpose,
			VerificationMethod: &constants.VerificationMethod,
			Challenge:          &constants.Challenge,
			Jws:                &constants.JWS,
		},
	},
}

// Create presentation request response info
var createPresentationRequestResponse = CreatePresentationRequestResponse{
	PresentationRequestId: &constants.PresentationRequestId,
}

// Get presentation request response info
var getPresentationRequestResponse = GetPresentationRequestResponse{
	Challenge:             &constants.Challenge,
	PresentationRequestId: &constants.PresentationRequestId,
}

// Get verifiable presentation response info
var getVerifiablePresentationResponse = GetVerifiablePresentationResponse{
	Presentation:        &presentation,
	VerificationComment: &constants.VerificationComment,
	VerificationStatus:  &constants.VerificationStatus,
}
