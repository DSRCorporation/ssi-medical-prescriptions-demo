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
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/storage"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/vc"
)

type VCService struct {
	issuerAgent   vc.Issuer
	holderAgent   vc.Holder
	verifierAgent vc.Verifier
	wallet        vc.Wallet
	storage       storage.VCStorage
}

func (s *VCService) ExchangeCredential(issuerId string, issuerKMSPassphrase string, holderId string, holderKMSPassphrase string, unsignedCredential domain.Credential) (credential domain.Credential, err error) {
	conn, err := s.getOrCreateConnection(issuerId, s.issuerAgent, holderId, s.holderAgent)
	if err != nil {
		return domain.Credential{}, err
	}

	var piid string
	var proofCredential domain.Credential // TODO signed unsignedCredential with key of holder
	piid, err = s.holderAgent.SendCredentialRequest(conn, proofCredential)
	if err != nil {
		return domain.Credential{}, err
	}

	credential, err = s.wallet.SignCredential(issuerId, issuerKMSPassphrase, vc.ProofOptions{}, unsignedCredential)
	if err != nil {
		return domain.Credential{}, err
	}

	err = s.issuerAgent.AcceptCredentialRequest(piid, credential)
	if err != nil {
		return domain.Credential{}, err
	}

	err = s.holderAgent.AcceptCredential(piid, credential.CredentialId)
	if err != nil {
		return domain.Credential{}, err
	}

	err = s.storage.SaveCredential(credential)
	if err != nil {
		return domain.Credential{}, err
	}

	return credential, nil
}

func (s *VCService) ExchangePresentation(verifierId string, holderId string, holderKMSPassphrase string, unsignedPresentation domain.Presentation) (presentation domain.Presentation, err error) {
	conn, err := s.getOrCreateConnection(verifierId, s.verifierAgent, holderId, s.holderAgent)
	if err != nil {
		return domain.Presentation{}, err
	}

	var piid string
	piid, err = s.verifierAgent.SendPresentationRequest(conn)
	if err != nil {
		return domain.Presentation{}, err
	}

	presentation, err = s.wallet.SignPresentation(holderId, holderKMSPassphrase, vc.ProofOptions{}, unsignedPresentation)
	if err != nil {
		return domain.Presentation{}, err
	}

	err = s.holderAgent.AcceptPresentationRequest(piid, presentation)
	if err != nil {
		return domain.Presentation{}, err
	}

	err = s.verifierAgent.AcceptPresentation(piid, presentation.PresentationId)
	if err != nil {
		return domain.Presentation{}, err
	}

	err = s.storage.SavePresentation(presentation)
	if err != nil {
		return domain.Presentation{}, err
	}

	return presentation, nil
}

func (s *VCService) GetCredentialById(credentialId string) (domain.Credential, error) {
	return s.storage.GetCredentialById(credentialId)
}

func (s *VCService) GetPresentationById(presentationId string) (domain.Presentation, error) {
	return s.storage.GetPresentationById(presentationId)
}

func (s *VCService) getOrCreateConnection(inviterId string, inviter vc.OOBInviter, inviteeId string, invitee vc.OOBInvitee) (conn domain.Connection, err error) {
	conn, err = s.storage.GetConnection(inviterId, inviteeId)
	if err == nil {
		return conn, nil
	}

	var invitation []byte
	invitation, err = inviter.CreateOOBInvitation()
	if err != nil {
		return domain.Connection{}, err
	}

	var connectionId string
	connectionId, err = invitee.AcceptOOBInvitation(invitation)
	if err != nil {
		return domain.Connection{}, err
	}

	conn, err = inviter.AcceptOOBRequest(connectionId)
	if err != nil {
		return domain.Connection{}, err
	}

	err = s.storage.SaveConnection(inviterId, inviteeId, conn)
	if err != nil {
		return domain.Connection{}, err
	}

	return conn, nil
}
