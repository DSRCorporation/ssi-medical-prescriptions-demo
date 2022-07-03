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

package vc

import (
	"encoding/json"
	"time"

	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

type ProofOptions struct {
	// Controller is a DID to be for signing. This option is required for issue/prove wallet features.
	Controller string
	// VerificationMethod is the URI of the verificationMethod used for the proof.
	// Optional, by default Controller public key matching 'assertion' for issue or 'authentication' for prove functions.
	VerificationMethod string
	// Created date of the proof.
	// Optional, current system time will be used.
	Created *time.Time
	// Domain is operational domain of a digital proof.
	// Optional, by default domain will not be part of proof.
	Domain string
	// Challenge is a random or pseudo-random value option authentication.
	// Optional, by default challenge will not be part of proof.
	Challenge string
	// ProofType is signature type used for signing.
	// Optional, by default proof will be generated in Ed25519Signature2018 format.
	ProofType string
	// ProofRepresentation is type of proof data expected, (Refer verifiable.SignatureProofValue)
	// Optional, by default proof will be represented as 'verifiable.SignatureProofValue'.
	ProofRepresentation *verifiable.SignatureRepresentation
}

type Wallet interface {
	SignCredential(userId string, passphrase string, proofOptions ProofOptions, credential domain.Credential) (domain.Credential, error)
	SignPresentation(userId string, passphrase string, proofOptions ProofOptions, presentaion domain.Presentation) (domain.Presentation, error)
	VerifyCredential(userId string, passphrase string, rawCredential json.RawMessage) error
	VerifyPresentation(userId string, passphrase string, rawPresentation json.RawMessage) error
}
