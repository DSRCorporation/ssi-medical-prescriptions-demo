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

package constants

// Credential offer response info
var CredentialOfferResponse string = "some credential offer"

// Challenge info
var Challenge string = "some challange"

// Patient info
var PatientName string = "John"
var PatientBirthday string = "13.10.1985"
var PatientSex string = "Man"
var PatientID string = "did:example:567"

// Issuance info
var IssuanceDate string = "20.02.2022"
var ExpirationDate string = "20.02.2032"

// Hospital info
var Location string = "Tokyo, Japan"
var HospitalName string = "The University of Tokyo Hospital Tokyo"
var DoctorName string = "Akuyiki Sato"
var Phone string = "+81 3 3473 8321"
var DoctorSignatureStamp string = "MILLE"
var PrefectureNumber float32 = 875
var ScoreVoteNumber float32 = 98
var MedicalInstitutionNumber float32 = 13

// Drugs info
var DrugNumber float32 = 3
var DrugName string = "some drug names"
var DrugType string = "tablet"
var RefillAvailability bool = true

// Ointment drug info
var AmountOfMedicine float32 = 5
var UseArea string = "some use area"
var ApplicationFrequency string = "sometimes"
var Usage string = "sometimes"

// Tablet drug info
var NumberOfDrug float32 = 5
var NumberOfDoses = "3 times"
var DaysOfMedication float32 = 7
var TimingToTakeMedicine = "in the morning after breakfast"

// Context info
var Context1 string = "https://www.w3.org/2018/credentials/v1"
var Context2 string = "https://www.w3.org/2018/credentials/examples/v1"
var Context3 string = "https://w3id.org/security/suites/jws-2020/v1"

// Type info
var Type1 string = "VerifiableCredential"
var Type2 string = "PrescriptionCredential"

// Issuer info
var ID string = "did:example:123"

// Credential proof info
var Type string = "JsonWebSignature2020"
var Created string = "2021-10-02T17:58:00Z"
var ProofPurpose string = "assertionMethod"
var VerificationMethod string = "did:example:123#key-0"
var JWS string = "ejJrNjQiOmZshbHNlfdJjcmlLKjpbImI2NCJdLCJhbGciOiJFZERTQSJ9..TR0VQqAghUT6AIVdHc8W8Q2aj12LOQjV_VZ1e134NU9Q20eBsNySPjNdmTWp2HkdquCnbRhBHxIbNeFEIOOhAg"

// Dids
var DIDs1 string = "did:example:1"
var DIDs2 string = "did:example:2"
var DIDs3 string = "did:example:3"

// Create presentation request response info
var PresentationRequestId string = "some presentation request ID"

// Get verifiable presentation response info
var VerificationStatus bool = true
var VerificationComment string = "some verification comment"
