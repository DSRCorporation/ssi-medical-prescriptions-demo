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
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/controller/rest"
)

// Global constants for integration tests.
var DoctorID = tmrand.Str(6)
var PatientID = tmrand.Str(6)
var PharmacyID = tmrand.Str(6)

// Credential offer response info
var credentialOfferResponse string = "some credential offer"

// Challenge info
var challenge string = "some challange"

// Patient info
var patientName string = "John"
var patientBirthday string = "13.10.1985"
var patientSex string = "Man"
var patientID string = "did:example:567"

// Issuance info
var issuanceDate string = "20.02.2022"
var expirationDate string = "20.02.2032"

// Hospital info
var location string = "Tokyo, Japan"
var hospitalName string = "The University of Tokyo Hospital Tokyo"
var doctorName string = "Akuyiki Sato"
var phone string = "+81 3 3473 8321"
var doctorSignatureStamp string = "MILLE"
var prefectureNumber float32 = 875
var scoreVoteNumber float32 = 98
var medicalInstitutionNumber float32 = 13

// Drugs info
var drugNumber float32 = 3
var drugName string = "some drug names"
var drugType string = "tablet"
var refillAvailability bool = true

// Ointment drug info
var amountOfMedicine float32 = 5
var useArea string = "some use area"
var applicationFrequency string = "sometimes"
var usage string = "sometimes"

// Tablet drug info
var numberOfDrug float32 = 5
var numberOfDoses = "3 times"
var daysOfMedication float32 = 7
var timingToTakeMedicine = "in the morning after breakfast"

// Context info
var context1 string = "https://www.w3.org/2018/credentials/v1"
var context2 string = "https://www.w3.org/2018/credentials/examples/v1"
var context3 string = "https://w3id.org/security/suites/jws-2020/v1"

// Type info
var type1 string = "VerifiableCredential"

// Issuer info
var id string = "did:example:123"

// Credential proof info
var typeSignature string = "JsonWebSignature2020"
var created string = "2021-10-02T17:58:00Z"
var proofPurpose string = "assertionMethod"
var verificationMethod string = "did:example:123#key-0"
var jws string = "ejJrNjQiOmZshbHNlfdJjcmlLKjpbImI2NCJdLCJhbGciOiJFZERTQSJ9..TR0VQqAghUT6AIVdHc8W8Q2aj12LOQjV_VZ1e134NU9Q20eBsNySPjNdmTWp2HkdquCnbRhBHxIbNeFEIOOhAg"

// Dids
var dids1 string = "did:example:1"
var dids2 string = "did:example:2"
var dids3 string = "did:example:3"

// Generate Verifiable Id
var verifiableId = "did:4347d8e1-9d7e-47cc-9dab-97de8afc4d95"

// Create presentation request response info
var presentationRequestId string = "some presentation request ID"

// Get verifiable presentation response info
var verificationStatus bool = true
var verificationComment string = "Verification Successful"

// Ointment drug info
var OintmentDrugInfoI interface{} = rest.OintmentDrugInfo{
	AmountOfMedicine:     &amountOfMedicine,
	UseArea:              &useArea,
	ApplicationFrequency: &applicationFrequency,
	Usage:                &usage,
}

// Tablet drug info
var TabletDrugInfoI interface{} = rest.TabletDrugInfo{
	NumberOfDrug:         &numberOfDrug,
	NumberOfDoses:        &numberOfDoses,
	DaysOfMedication:     &daysOfMedication,
	TimingToTakeMedicine: &timingToTakeMedicine,
}

// Credential offer response info
var CredentialOfferResponseInfo = rest.CredentialOfferResponse{
	CredentialOfferId: &credentialOfferResponse,
}

// Get credential offer response info
var GetCredentialOfferResponseInfo = rest.GetCredentialOfferResponse{
	Prescription: &rest.Prescription{
		PatientInfo: &struct {
			Birthday *string "json:\"birthday,omitempty\""
			Name     *string "json:\"name,omitempty\""
			Sex      *string "json:\"sex,omitempty\""
		}{
			Birthday: &patientBirthday,
			Name:     &patientName,
			Sex:      &patientSex,
		},
		IssuanceInfo: &struct {
			ExpirationDate *string "json:\"expirationDate,omitempty\""
			IssuanceDate   *string "json:\"issuanceDate,omitempty\""
		}{
			IssuanceDate:   &issuanceDate,
			ExpirationDate: &expirationDate,
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
			DoctorName:               &doctorName,
			DoctorSignatureStamp:     &doctorSignatureStamp,
			HospitalName:             &hospitalName,
			Location:                 &location,
			MedicalInstitutionNumber: &medicalInstitutionNumber,
			Phone:                    &phone,
			PrefectureNumber:         &prefectureNumber,
			ScoreVoteNumber:          &scoreVoteNumber,
		},
		Drugs: &[]rest.Drug{
			{
				DrugName:           &drugName,
				DrugNumber:         &drugNumber,
				DrugType:           &drugType,
				Info:               &OintmentDrugInfoI,
				RefillAvailability: &refillAvailability,
			},
			{
				DrugName:           &drugName,
				DrugNumber:         &drugNumber,
				DrugType:           &drugType,
				Info:               &TabletDrugInfoI,
				RefillAvailability: &refillAvailability,
			},
		},
	},
}

// Credential response info
var CredentialResponseInfo = rest.Credential{
	Context: &[]string{
		context1,
		context2,
		context3,
	},
	Type:           &type1,
	Id:             &verifiableId,
	Issuer:         &id,
	IssuanceDate:   &issuanceDate,
	ExpirationDate: &expirationDate,
	CredentialSubject: &struct {
		Id           *string            "json:\"id,omitempty\""
		Name         *string            "json:\"name,omitempty\""
		Prescription *rest.Prescription "json:\"prescription,omitempty\""
	}{
		Id:           &patientID,
		Name:         &patientName,
		Prescription: GetCredentialOfferResponseInfo.Prescription,
	},
	Proof: &rest.CredentialProof{
		Created:            &created,
		Jws:                &jws,
		ProofPurpose:       &proofPurpose,
		Type:               &typeSignature,
		VerificationMethod: &verificationMethod,
	},
}

// Get all DIDs response info
var GetAllDidsResponseInfo = rest.GetAllDidsResponse{
	Dids: &[]string{
		dids1,
		dids2,
		dids3,
	},
}

// Get all prescription credential response info
var GetAllPrescriptionCredentialResponseInfo = rest.GetAllPrescriptionCredentialResponse{
	Credentials: &[]rest.Credential{CredentialResponseInfo},
}

// Get prescription credential response info
var GetPrescriptionCredentialResponseInfo = rest.CredentialResponse{
	Credential: &CredentialResponseInfo,
}

// Presentation info
var PresentationInfo = rest.Presentation{
	Context: &[]string{
		context1,
		context2,
		context3,
	},
	Type: &type1,
	VerifiableCredential: &[]rest.Credential{
		CredentialResponseInfo,
	},
	Proof: &[]rest.PresentationProof{
		{
			Type:               &typeSignature,
			Created:            &created,
			ProofPurpose:       &proofPurpose,
			VerificationMethod: &verificationMethod,
			Jws:                &jws,
		},
	},
}

// Create presentation request response info
var CreatePresentationRequestResponseInfo = rest.CreatePresentationRequestResponse{
	PresentationRequestId: &presentationRequestId,
}

// Get presentation request response info
var GetPresentationRequestResponseInfo = rest.GetPresentationRequestResponse{
	Challenge:             &challenge,
	PresentationRequestId: &presentationRequestId,
}

// Get verifiable presentation response info
var GetVerifiablePresentationResponseInfo = rest.GetVerifiablePresentationResponse{
	Presentation:        &PresentationInfo,
	VerificationComment: &verificationComment,
	VerificationStatus:  &verificationStatus,
}
