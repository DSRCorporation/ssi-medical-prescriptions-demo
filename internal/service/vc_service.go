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

package service

import (
	"encoding/json"

	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/storage"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/vc"
)

type VCService struct {
	issuerAgent    vc.Issuer
	holderAgent    vc.Holder
	verifierAgent  vc.Verifier
	issuerWallet   vc.Wallet
	holderWallet   vc.Wallet
	verifierWallet vc.Wallet
	storage        storage.VCStorage
}

func NewVCService(storage storage.VCStorage, issuerAgent vc.Issuer, holderAgent vc.Holder, verifierAgent vc.Verifier, issuerWallet vc.Wallet, holderWallet vc.Wallet, verifierWallet vc.Wallet) *VCService {
	return &VCService{
		storage:       storage,
		issuerAgent:   issuerAgent,
		holderAgent:   holderAgent,
		verifierAgent: verifierAgent,
		issuerWallet:  issuerWallet,
		holderWallet:  holderWallet,
	}
}

func (s *VCService) ExchangeCredential(issuerId string, issuerKMSPassphrase string, holderId string, holderKMSPassphrase string, unsignedCredential domain.Credential) (credential domain.Credential, err error) {
	conn, err := s.createConnection(issuerId, s.issuerAgent, holderId, s.holderAgent)
	if err != nil {
		return domain.Credential{}, err
	}

	piid, err := s.issuerAgent.SendCredentialOffer(conn, unsignedCredential)
	if err != nil {
		return domain.Credential{}, err
	}

	offeredCredential, err := s.holderAgent.GetCredentialFromOffer(piid)
	if err != nil {
		return domain.Credential{}, err
	}

	err = s.holderAgent.AcceptOffer(piid)
	if err != nil {
		return domain.Credential{}, err
	}

	credential, err = s.holderWallet.SignCredential(holderId, holderKMSPassphrase, offeredCredential.HolderDID, *offeredCredential)
	if err != nil {
		return domain.Credential{}, err
	}

	_, err = s.holderAgent.SendCredentialRequest(conn, credential)
	if err != nil {
		return domain.Credential{}, err
	}

	requestedCredential, err := s.issuerAgent.GetCredentialFromRequest(piid)
	if err != nil {
		return domain.Credential{}, err
	}

	credential, err = s.issuerWallet.SignCredential(issuerId, issuerKMSPassphrase, requestedCredential.IssuerDID, *requestedCredential)
	if err != nil {
		return domain.Credential{}, err
	}

	err = s.issuerAgent.AcceptCredentialRequest(piid, credential)
	if err != nil {
		return domain.Credential{}, err
	}

	issuedCredential, err := s.holderAgent.GetIssuedCredential(piid)
	if err != nil {
		return domain.Credential{}, err
	}

	err = s.holderAgent.AcceptCredential(piid, issuedCredential.CredentialId)
	if err != nil {
		return domain.Credential{}, err
	}

	err = s.storage.SaveCredential(*issuedCredential)
	if err != nil {
		return domain.Credential{}, err
	}

	return *issuedCredential, nil
}

func (s *VCService) ExchangePresentation(verifierId string, holderId string, holderKMSPassphrase string, unsignedPresentation domain.Presentation) (presentation domain.Presentation, err error) {
	conn, err := s.createConnection(verifierId, s.verifierAgent, holderId, s.holderAgent)
	if err != nil {
		return domain.Presentation{}, err
	}

	var piid string
	piid, err = s.verifierAgent.SendPresentationRequest(conn)
	if err != nil {
		return domain.Presentation{}, err
	}

	presentation, err = s.holderWallet.SignPresentation(holderId, holderKMSPassphrase, unsignedPresentation.HolderDID, unsignedPresentation)
	if err != nil {
		return domain.Presentation{}, err
	}

	err = s.holderAgent.AcceptPresentationRequest(piid, presentation)
	if err != nil {
		return domain.Presentation{}, err
	}

	issuedPresentation, err := s.verifierAgent.GetIssuedPresentation(piid)

	err = s.verifierAgent.AcceptPresentation(piid, issuedPresentation.PresentationId)
	if err != nil {
		return domain.Presentation{}, err
	}

	err = s.storage.SavePresentation(*issuedPresentation)
	if err != nil {
		return domain.Presentation{}, err
	}

	return *issuedPresentation, nil
}

func (s *VCService) GetCredentialById(credentialId string) (domain.Credential, error) {
	return s.storage.GetCredentialById(credentialId)
}

func (s *VCService) GetPresentationById(presentationId string) (domain.Presentation, error) {
	return s.storage.GetPresentationById(presentationId)
}

func (s *VCService) VerifyCredential(rawCredential json.RawMessage) error {
	userId, passphrase, err := s.storage.GetVerifierWalletCredentials()
	if err != nil {
		return err
	}

	return s.verifierWallet.VerifyCredential(userId, passphrase, rawCredential)
}

func (s *VCService) VerifyPresentation(rawPresentation json.RawMessage) error {
	userId, passphrase, err := s.storage.GetVerifierWalletCredentials()
	if err != nil {
		return err
	}

	return s.verifierWallet.VerifyPresentation(userId, passphrase, rawPresentation)
}

func (s *VCService) createConnection(inviterId string, inviter vc.OOBInviter, inviteeId string, invitee vc.OOBInvitee) (conn domain.Connection, err error) {

	invitation, err := inviter.CreateOOBInvitation()
	if err != nil {
		return domain.Connection{}, err
	}

	err = invitee.AcceptOOBInvitation(invitation)
	if err != nil {
		return domain.Connection{}, err
	}

	conn, err = inviter.AcceptOOBRequest(invitation)
	if err != nil {
		return domain.Connection{}, err
	}

	return conn, nil
}
