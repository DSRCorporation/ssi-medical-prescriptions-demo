// Package mock provides primitives to interact with the openapi HTTP API.
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
	ExpirationDate *string `json:"expirationDate,omitempty"`
	IssuanceDate   *string `json:"issuanceDate,omitempty"`
	Issuer         *struct {
		Id   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"issuer,omitempty"`
	Proof *CredentialProof `json:"proof,omitempty"`
	Type  *[]string        `json:"type,omitempty"`
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
	Credentials *[]map[string]interface{} `json:"credentials,omitempty"`
}

// GetCredentialOfferResponse defines model for getCredentialOfferResponse.
type GetCredentialOfferResponse struct {
	Challenge    *string       `json:"challenge,omitempty"`
	Prescription *Prescription `json:"prescription,omitempty"`
}

// GetPrescriptionCredentialResponse defines model for getPrescriptionCredentialResponse.
type GetPrescriptionCredentialResponse struct {
	Credentials *map[string]interface{} `json:"credentials,omitempty"`
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
	Challenge          *string `json:"challenge,omitempty"`
	Created            *string `json:"created,omitempty"`
	Jws                *string `json:"jws,omitempty"`
	ProofPurpose       *string `json:"proofPurpose,omitempty"`
	Type               *string `json:"type,omitempty"`
	VerificationMethod *string `json:"verificationMethod,omitempty"`
}

// TabletDrugInfo defines model for tabletDrugInfo.
type TabletDrugInfo struct {
	DaysOfMedication     *float32 `json:"daysOfMedication,omitempty"`
	NumberOfDoses        *string  `json:"numberOfDoses,omitempty"`
	NumberOfDrug         *float32 `json:"numberOfDrug,omitempty"`
	TimingToTakeMedicine *string  `json:"timingToTakeMedicine,omitempty"`
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
	Challenge             *string `json:"challenge,omitempty"`
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