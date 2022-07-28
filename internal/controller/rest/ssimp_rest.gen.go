// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package rest

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// CreatePresentationRequestResponse defines model for createPresentationRequestResponse.
type CreatePresentationRequestResponse struct {
	PresentationRequestId *string `json:"presentationRequestId,omitempty"`
}

// Credential defines model for credential.
type Credential struct {
	Context           *[]string `json:"@context,omitempty"`
	CredentialSubject *struct {
		Id           *string       `json:"id,omitempty"`
		Name         *string       `json:"name,omitempty"`
		Prescription *Prescription `json:"prescription,omitempty"`
	} `json:"credentialSubject,omitempty"`
	ExpirationDate *string          `json:"expirationDate,omitempty"`
	Id             *string          `json:"id,omitempty"`
	IssuanceDate   *string          `json:"issuanceDate,omitempty"`
	Issuer         *string          `json:"issuer,omitempty"`
	Proof          *CredentialProof `json:"proof,omitempty"`
	Type           *string          `json:"type,omitempty"`
}

// CredentialOfferResponse defines model for credentialOfferResponse.
type CredentialOfferResponse struct {
	CredentialOfferId *string `json:"credentialOfferId,omitempty"`
}

// CredentialProof defines model for credentialProof.
type CredentialProof struct {
	Created            *string `json:"created,omitempty"`
	Jws                *string `json:"jws,omitempty"`
	ProofPurpose       *string `json:"proofPurpose,omitempty"`
	Type               *string `json:"type,omitempty"`
	VerificationMethod *string `json:"verificationMethod,omitempty"`
}

// CredentialResponse defines model for credentialResponse.
type CredentialResponse struct {
	Credential *Credential `json:"credential,omitempty"`
}

// Drug defines model for drug.
type Drug struct {
	DrugName           *string      `json:"drugName,omitempty"`
	DrugNumber         *float32     `json:"drugNumber,omitempty"`
	DrugType           *string      `json:"drugType,omitempty"`
	Info               *interface{} `json:"info,omitempty"`
	RefillAvailability *bool        `json:"refillAvailability,omitempty"`
}

// Error defines model for error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// GetAllDidsResponse defines model for getAllDidsResponse.
type GetAllDidsResponse struct {
	Dids *[]string `json:"dids,omitempty"`
}

// GetAllPrescriptionCredentialResponse defines model for getAllPrescriptionCredentialResponse.
type GetAllPrescriptionCredentialResponse struct {
	Credentials *[]Credential `json:"credentials,omitempty"`
}

// GetCredentialOfferResponse defines model for getCredentialOfferResponse.
type GetCredentialOfferResponse struct {
	Prescription *Prescription `json:"prescription,omitempty"`
}

// GetPresentationRequestResponse defines model for getPresentationRequestResponse.
type GetPresentationRequestResponse struct {
	Challenge             *string `json:"challenge,omitempty"`
	PresentationRequestId *string `json:"presentationRequestId,omitempty"`
}

// GetVerifiablePresentationResponse defines model for getVerifiablePresentationResponse.
type GetVerifiablePresentationResponse struct {
	Presentation        *Presentation `json:"presentation,omitempty"`
	VerificationComment *string       `json:"verificationComment,omitempty"`
	VerificationStatus  *bool         `json:"verificationStatus,omitempty"`
}

// OintmentDrugInfo defines model for ointmentDrugInfo.
type OintmentDrugInfo struct {
	AmountOfMedicine     *float32 `json:"amountOfMedicine,omitempty"`
	ApplicationFrequency *string  `json:"applicationFrequency,omitempty"`
	Usage                *string  `json:"usage,omitempty"`
	UseArea              *string  `json:"useArea,omitempty"`
}

// Prescription defines model for prescription.
type Prescription struct {
	Drugs        *[]Drug `json:"drugs,omitempty"`
	HospitalInfo *struct {
		DoctorName               *string  `json:"doctorName,omitempty"`
		DoctorSignatureStamp     *string  `json:"doctorSignatureStamp,omitempty"`
		HospitalName             *string  `json:"hospitalName,omitempty"`
		Location                 *string  `json:"location,omitempty"`
		MedicalInstitutionNumber *float32 `json:"medicalInstitutionNumber,omitempty"`
		Phone                    *string  `json:"phone,omitempty"`
		PrefectureNumber         *float32 `json:"prefectureNumber,omitempty"`
		ScoreVoteNumber          *float32 `json:"scoreVoteNumber,omitempty"`
	} `json:"hospitalInfo,omitempty"`
	IssuanceInfo *struct {
		ExpirationDate *string `json:"expirationDate,omitempty"`
		IssuanceDate   *string `json:"issuanceDate,omitempty"`
	} `json:"issuanceInfo,omitempty"`
	PatientInfo *struct {
		Birthday *string `json:"birthday,omitempty"`
		Name     *string `json:"name,omitempty"`
		Sex      *string `json:"sex,omitempty"`
	} `json:"patientInfo,omitempty"`
}

// Presentation defines model for presentation.
type Presentation struct {
	Context              *[]string            `json:"@context,omitempty"`
	Proof                *[]PresentationProof `json:"proof,omitempty"`
	Type                 *string              `json:"type,omitempty"`
	VerifiableCredential *[]Credential        `json:"verifiableCredential,omitempty"`
}

// PresentationProof defines model for presentationProof.
type PresentationProof struct {
	Created            *string `json:"created,omitempty"`
	Jws                *string `json:"jws,omitempty"`
	ProofPurpose       *string `json:"proofPurpose,omitempty"`
	Type               *string `json:"type,omitempty"`
	VerificationMethod *string `json:"verificationMethod,omitempty"`
}

// PresentationResponse defines model for presentationResponse.
type PresentationResponse struct {
	Presentation *Presentation `json:"presentation,omitempty"`
}

// TabletDrugInfo defines model for tabletDrugInfo.
type TabletDrugInfo struct {
	DaysOfMedication     *float32 `json:"daysOfMedication,omitempty"`
	NumberOfDoses        *string  `json:"numberOfDoses,omitempty"`
	NumberOfDrug         *float32 `json:"numberOfDrug,omitempty"`
	TimingToTakeMedicine *string  `json:"timingToTakeMedicine,omitempty"`
}

// VerificationResponse defines model for verificationResponse.
type VerificationResponse struct {
	Verified *bool `json:"verified,omitempty"`
}

// PostV1DoctorsDoctorIdPrescriptionsCredentialOffersJSONBody defines parameters for PostV1DoctorsDoctorIdPrescriptionsCredentialOffers.
type PostV1DoctorsDoctorIdPrescriptionsCredentialOffersJSONBody = Prescription

// PostV1PatientsPatientIdPrescriptionsCredentialsJSONBody defines parameters for PostV1PatientsPatientIdPrescriptionsCredentials.
type PostV1PatientsPatientIdPrescriptionsCredentialsJSONBody struct {
	CredentialOfferId *string `json:"credentialOfferId,omitempty"`
	Did               *string `json:"did,omitempty"`
	KmsPassphrase     *string `json:"kmsPassphrase,omitempty"`
}

// PostV1PatientsPatientIdPrescriptionsPresentationsJSONBody defines parameters for PostV1PatientsPatientIdPrescriptionsPresentations.
type PostV1PatientsPatientIdPrescriptionsPresentationsJSONBody struct {
	CredentialId          *string `json:"credentialId,omitempty"`
	KmsPassphrase         *string `json:"kmsPassphrase,omitempty"`
	PresentationRequestId *string `json:"presentationRequestId,omitempty"`
}

// PostV1VcVerifyCredentialJSONBody defines parameters for PostV1VcVerifyCredential.
type PostV1VcVerifyCredentialJSONBody = Credential

// PostV1VcVerifyPresentationJSONBody defines parameters for PostV1VcVerifyPresentation.
type PostV1VcVerifyPresentationJSONBody = Presentation

// PostV1DoctorsDoctorIdPrescriptionsCredentialOffersJSONRequestBody defines body for PostV1DoctorsDoctorIdPrescriptionsCredentialOffers for application/json ContentType.
type PostV1DoctorsDoctorIdPrescriptionsCredentialOffersJSONRequestBody = PostV1DoctorsDoctorIdPrescriptionsCredentialOffersJSONBody

// PostV1PatientsPatientIdPrescriptionsCredentialsJSONRequestBody defines body for PostV1PatientsPatientIdPrescriptionsCredentials for application/json ContentType.
type PostV1PatientsPatientIdPrescriptionsCredentialsJSONRequestBody PostV1PatientsPatientIdPrescriptionsCredentialsJSONBody

// PostV1PatientsPatientIdPrescriptionsPresentationsJSONRequestBody defines body for PostV1PatientsPatientIdPrescriptionsPresentations for application/json ContentType.
type PostV1PatientsPatientIdPrescriptionsPresentationsJSONRequestBody PostV1PatientsPatientIdPrescriptionsPresentationsJSONBody

// PostV1VcVerifyCredentialJSONRequestBody defines body for PostV1VcVerifyCredential for application/json ContentType.
type PostV1VcVerifyCredentialJSONRequestBody = PostV1VcVerifyCredentialJSONBody

// PostV1VcVerifyPresentationJSONRequestBody defines body for PostV1VcVerifyPresentation for application/json ContentType.
type PostV1VcVerifyPresentationJSONRequestBody = PostV1VcVerifyPresentationJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Creates credential offer for prescription (generates a link for QR code)
	// (POST /v1/doctors/{doctorId}/prescriptions/credential-offers/)
	PostV1DoctorsDoctorIdPrescriptionsCredentialOffers(ctx echo.Context, doctorId string) error
	// Gets credential offer by id
	// (GET /v1/doctors/{doctorId}/prescriptions/credential-offers/{credentialOfferId})
	GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferId(ctx echo.Context, doctorId string, credentialOfferId string) error
	// Gets credential issued for given credential offer
	// (GET /v1/doctors/{doctorId}/prescriptions/credential-offers/{credentialOfferId}/credential)
	GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferIdCredential(ctx echo.Context, doctorId string, credentialOfferId string) error
	// Gets all dids belonging to given patient
	// (GET /v1/patients/{patientId}/dids)
	GetV1PatientsPatientIdDids(ctx echo.Context, patientId string) error
	// Gets all prescription credentials issued for given patient
	// (GET /v1/patients/{patientId}/prescriptions/credentials)
	GetV1PatientsPatientIdPrescriptionsCredentials(ctx echo.Context, patientId string) error
	// Creates credential in response to credential offer from doctor
	// (POST /v1/patients/{patientId}/prescriptions/credentials/)
	PostV1PatientsPatientIdPrescriptionsCredentials(ctx echo.Context, patientId string) error
	// Gets prescription credential by id issued for given patient
	// (GET /v1/patients/{patientId}/prescriptions/credentials/{credentialId})
	GetV1PatientsPatientIdPrescriptionsCredentialsCredentialId(ctx echo.Context, patientId string, credentialId string) error
	// Creates verifiable presentation in response to prescription presentation request from pharmacy
	// (POST /v1/patients/{patientId}/prescriptions/presentations)
	PostV1PatientsPatientIdPrescriptionsPresentations(ctx echo.Context, patientId string) error
	// Creates presentation request for prescription (generates link for a QR code)
	// (POST /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests)
	PostV1PharmaciesPharmacyIdPrescriptionsPresentationRequests(ctx echo.Context, pharmacyId string) error
	// Gets presentation request for prescription by request id
	// (GET /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId})
	GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestId(ctx echo.Context, pharmacyId string, presentationRequestId string) error
	// Gets verifiable presentation for given presentation request
	// (GET /v1/pharmacies/{pharmacyId}/prescriptions/presentation-requests/{presentationRequestId}/presentation)
	GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestIdPresentation(ctx echo.Context, pharmacyId string, presentationRequestId string) error
	// Verify Credential
	// (POST /v1/vc/verify-credential)
	PostV1VcVerifyCredential(ctx echo.Context) error
	// Verify Credential
	// (POST /v1/vc/verify-presentation)
	PostV1VcVerifyPresentation(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostV1DoctorsDoctorIdPrescriptionsCredentialOffers converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1DoctorsDoctorIdPrescriptionsCredentialOffers(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "doctorId" -------------
	var doctorId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "doctorId", runtime.ParamLocationPath, ctx.Param("doctorId"), &doctorId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter doctorId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostV1DoctorsDoctorIdPrescriptionsCredentialOffers(ctx, doctorId)
	return err
}

// GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferId converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "doctorId" -------------
	var doctorId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "doctorId", runtime.ParamLocationPath, ctx.Param("doctorId"), &doctorId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter doctorId: %s", err))
	}

	// ------------- Path parameter "credentialOfferId" -------------
	var credentialOfferId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "credentialOfferId", runtime.ParamLocationPath, ctx.Param("credentialOfferId"), &credentialOfferId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter credentialOfferId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferId(ctx, doctorId, credentialOfferId)
	return err
}

// GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferIdCredential converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferIdCredential(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "doctorId" -------------
	var doctorId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "doctorId", runtime.ParamLocationPath, ctx.Param("doctorId"), &doctorId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter doctorId: %s", err))
	}

	// ------------- Path parameter "credentialOfferId" -------------
	var credentialOfferId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "credentialOfferId", runtime.ParamLocationPath, ctx.Param("credentialOfferId"), &credentialOfferId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter credentialOfferId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferIdCredential(ctx, doctorId, credentialOfferId)
	return err
}

// GetV1PatientsPatientIdDids converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1PatientsPatientIdDids(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "patientId" -------------
	var patientId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "patientId", runtime.ParamLocationPath, ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1PatientsPatientIdDids(ctx, patientId)
	return err
}

// GetV1PatientsPatientIdPrescriptionsCredentials converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1PatientsPatientIdPrescriptionsCredentials(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "patientId" -------------
	var patientId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "patientId", runtime.ParamLocationPath, ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1PatientsPatientIdPrescriptionsCredentials(ctx, patientId)
	return err
}

// PostV1PatientsPatientIdPrescriptionsCredentials converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1PatientsPatientIdPrescriptionsCredentials(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "patientId" -------------
	var patientId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "patientId", runtime.ParamLocationPath, ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostV1PatientsPatientIdPrescriptionsCredentials(ctx, patientId)
	return err
}

// GetV1PatientsPatientIdPrescriptionsCredentialsCredentialId converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1PatientsPatientIdPrescriptionsCredentialsCredentialId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "patientId" -------------
	var patientId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "patientId", runtime.ParamLocationPath, ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	// ------------- Path parameter "credentialId" -------------
	var credentialId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "credentialId", runtime.ParamLocationPath, ctx.Param("credentialId"), &credentialId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter credentialId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1PatientsPatientIdPrescriptionsCredentialsCredentialId(ctx, patientId, credentialId)
	return err
}

// PostV1PatientsPatientIdPrescriptionsPresentations converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1PatientsPatientIdPrescriptionsPresentations(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "patientId" -------------
	var patientId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "patientId", runtime.ParamLocationPath, ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostV1PatientsPatientIdPrescriptionsPresentations(ctx, patientId)
	return err
}

// PostV1PharmaciesPharmacyIdPrescriptionsPresentationRequests converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1PharmaciesPharmacyIdPrescriptionsPresentationRequests(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pharmacyId" -------------
	var pharmacyId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "pharmacyId", runtime.ParamLocationPath, ctx.Param("pharmacyId"), &pharmacyId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pharmacyId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostV1PharmaciesPharmacyIdPrescriptionsPresentationRequests(ctx, pharmacyId)
	return err
}

// GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestId converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pharmacyId" -------------
	var pharmacyId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "pharmacyId", runtime.ParamLocationPath, ctx.Param("pharmacyId"), &pharmacyId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pharmacyId: %s", err))
	}

	// ------------- Path parameter "presentationRequestId" -------------
	var presentationRequestId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "presentationRequestId", runtime.ParamLocationPath, ctx.Param("presentationRequestId"), &presentationRequestId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter presentationRequestId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestId(ctx, pharmacyId, presentationRequestId)
	return err
}

// GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestIdPresentation converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestIdPresentation(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pharmacyId" -------------
	var pharmacyId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "pharmacyId", runtime.ParamLocationPath, ctx.Param("pharmacyId"), &pharmacyId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pharmacyId: %s", err))
	}

	// ------------- Path parameter "presentationRequestId" -------------
	var presentationRequestId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "presentationRequestId", runtime.ParamLocationPath, ctx.Param("presentationRequestId"), &presentationRequestId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter presentationRequestId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestIdPresentation(ctx, pharmacyId, presentationRequestId)
	return err
}

// PostV1VcVerifyCredential converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1VcVerifyCredential(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostV1VcVerifyCredential(ctx)
	return err
}

// PostV1VcVerifyPresentation converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1VcVerifyPresentation(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostV1VcVerifyPresentation(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/v1/doctors/:doctorId/prescriptions/credential-offers/", wrapper.PostV1DoctorsDoctorIdPrescriptionsCredentialOffers)
	router.GET(baseURL+"/v1/doctors/:doctorId/prescriptions/credential-offers/:credentialOfferId", wrapper.GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferId)
	router.GET(baseURL+"/v1/doctors/:doctorId/prescriptions/credential-offers/:credentialOfferId/credential", wrapper.GetV1DoctorsDoctorIdPrescriptionsCredentialOffersCredentialOfferIdCredential)
	router.GET(baseURL+"/v1/patients/:patientId/dids", wrapper.GetV1PatientsPatientIdDids)
	router.GET(baseURL+"/v1/patients/:patientId/prescriptions/credentials", wrapper.GetV1PatientsPatientIdPrescriptionsCredentials)
	router.POST(baseURL+"/v1/patients/:patientId/prescriptions/credentials/", wrapper.PostV1PatientsPatientIdPrescriptionsCredentials)
	router.GET(baseURL+"/v1/patients/:patientId/prescriptions/credentials/:credentialId", wrapper.GetV1PatientsPatientIdPrescriptionsCredentialsCredentialId)
	router.POST(baseURL+"/v1/patients/:patientId/prescriptions/presentations", wrapper.PostV1PatientsPatientIdPrescriptionsPresentations)
	router.POST(baseURL+"/v1/pharmacies/:pharmacyId/prescriptions/presentation-requests", wrapper.PostV1PharmaciesPharmacyIdPrescriptionsPresentationRequests)
	router.GET(baseURL+"/v1/pharmacies/:pharmacyId/prescriptions/presentation-requests/:presentationRequestId", wrapper.GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestId)
	router.GET(baseURL+"/v1/pharmacies/:pharmacyId/prescriptions/presentation-requests/:presentationRequestId/presentation", wrapper.GetV1PharmaciesPharmacyIdPrescriptionsPresentationRequestsPresentationRequestIdPresentation)
	router.POST(baseURL+"/v1/vc/verify-credential", wrapper.PostV1VcVerifyCredential)
	router.POST(baseURL+"/v1/vc/verify-presentation", wrapper.PostV1VcVerifyPresentation)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAD//7a6k9XqNQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
