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
// (POST /doctors/{doctorId}/prescriptions/credential-offers/)
func (*SSIMPMockImpl) PostDoctorsDoctorIdPrescriptionsCredentialOffers(ctx echo.Context, doctorId string) error {
	return &ErrorImpl{msg: "not implemented"}
}

// Gets credential offer by id
// (GET /doctors/{doctorId}/prescriptions/credential-offers/{credentialOfferId})
func (*SSIMPMockImpl) GetDoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferId(ctx echo.Context, doctorId string, credentialOfferId string) error {
	// return ctx.String(http.StatusOK, "Hello, World!")
	return &ErrorImpl{msg: "not implemented"}
}

// Gets credential issued for given credential offer
// (GET /doctors/{doctorId}/prescriptions/credential-offers/{credentialOfferId}/credential)
func (*SSIMPMockImpl) GetDoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferIdCredential(ctx echo.Context, doctorId string, credentialOfferId string) error {
	return &ErrorImpl{msg: "not implemented"}
}

// Gets all dids belonging to given patient
// (GET /patients/{patientId}/dids)
func (*SSIMPMockImpl) GetPatientsPatientIdDids(ctx echo.Context, patientId string) error {
	return &ErrorImpl{msg: "not implemented"}
}

// Gets all prescription credentials issued for given patient
// (GET /patients/{patientId}/prescriptions/credentials)
func (*SSIMPMockImpl) GetPatientsPatientIdPrescriptionsCredentials(ctx echo.Context, patientId string) error {
	return &ErrorImpl{msg: "not implemented"}
}

// Creates credential in response to credential offer from doctor
// (POST /patients/{patientId}/prescriptions/credentials/)
func (*SSIMPMockImpl) PostPatientsPatientIdPrescriptionsCredentials(ctx echo.Context, patientId string) error {
	return &ErrorImpl{msg: "not implemented"}
}

// Gets prescription credential by id issued for given patient
// (GET /patients/{patientId}/prescriptions/credentials/{credentialId})
func (*SSIMPMockImpl) GetPatientsPatientIdPrescriptionsCredentialsCredentialId(ctx echo.Context, patientId string, credentialId string) error {
	return &ErrorImpl{msg: "not implemented"}
}

// Generates canonical jws payload of credential presentaion for signing
// (POST /patients/{patientId}/prescriptions/credentials/{credentialId}/presentation-jws-payload)
func (*SSIMPMockImpl) PostPatientsPatientIdPrescriptionsCredentialsCredentialIdPresentationJwsPayload(ctx echo.Context, patientId string, credentialId string) error {
	return &ErrorImpl{msg: "not implemented"}
}

// Creates verifiable presentation in response to prescription presentation request from pharmacy
// (POST /patients/{patientId}/prescriptions/presentations/)
func (*SSIMPMockImpl) PostPatientsPatientIdPrescriptionsPresentations(ctx echo.Context, patientId string) error {
	return &ErrorImpl{msg: "not implemented"}
}

// Creates presentation request for prescription (generates link for a QR code)
// (POST /pharmacies/{pharmacyId}/prescriptions/presentation-requests)
func (*SSIMPMockImpl) PostPharmaciesPharmacyIdPrescriptionsPresentationRequests(ctx echo.Context, pharmacyId string) error {
	return &ErrorImpl{msg: "not implemented"}
}

// Gets presentation request for prescription by request id
// (GET /pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId})
func (*SSIMPMockImpl) GetPharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestId(ctx echo.Context, pharmacyId string, presentationRequestId string) error {
	return &ErrorImpl{msg: "not implemented"}
}

// Gets verifiable presentation for given presentation request
// (GET /pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId}/presentation)
func (*SSIMPMockImpl) GetPharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestIdPresentation(ctx echo.Context, pharmacyId string, presentationRequestId string) error {
	return &ErrorImpl{msg: "not implemented"}
}
