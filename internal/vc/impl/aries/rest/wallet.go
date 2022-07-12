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

package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/hyperledger/aries-framework-go/pkg/controller/command/vcwallet"
	"github.com/hyperledger/aries-framework-go/pkg/doc/util"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/hyperledger/aries-framework-go/pkg/wallet"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
)

type Wallet struct {
	client *resty.Client
}

func NewWallet(endpoint string) (*Wallet, error) {
	client := resty.New()
	client.SetBaseURL(endpoint)
	return &Wallet{client: client}, nil
}

func (w *Wallet) SignCredential(userId string, passphrase string, did string, credential domain.Credential) (domain.Credential, error) {
	token, err := w.open(userId, passphrase)
	if err != nil {
		return domain.Credential{}, err
	}

	defer w.close(userId, passphrase)

	rawCredential, err := makeRawCredential(credential)
	if err != nil {
		return domain.Credential{}, err
	}

	var res struct {
		Credential json.RawMessage `json:"credential"`
	}
	resp, err := w.client.R().
		SetBody(&vcwallet.IssueRequest{
			WalletAuth: vcwallet.WalletAuth{UserID: userId, Auth: token},
			Credential: *rawCredential,
			ProofOptions: &wallet.ProofOptions{
				Controller: did,
			}}).
		SetResult(&res).
		Post("/vcwallet/issue")

	if err != nil {
		return domain.Credential{}, err
	}

	if resp.StatusCode() == http.StatusOK {
		credential.RawCredentialWithProof = res.Credential
		return credential, nil
	} else {
		return domain.Credential{}, errors.New(string(resp.Body()))
	}
}

func (w *Wallet) SignPresentation(userId string, passphrase string, did string, presentation domain.Presentation) (domain.Presentation, error) {
	token, err := w.open(userId, passphrase)
	if err != nil {
		return domain.Presentation{}, err
	}

	defer w.close(userId, passphrase)

	var res struct {
		Presentation json.RawMessage `json:"presentation"`
	}

	resp, err := w.client.R().
		SetBody(&vcwallet.ProveRequest{
			WalletAuth:     vcwallet.WalletAuth{UserID: userId, Auth: token},
			RawCredentials: []json.RawMessage{presentation.Credential.RawCredentialWithProof},
			ProofOptions: &wallet.ProofOptions{
				Controller: did,
			}}).
		SetResult(&res).
		Post("/vcwallet/prove")

	if err != nil {
		return domain.Presentation{}, err
	}

	if resp.StatusCode() == http.StatusOK {
		presentation.RawPresentationWithProof = res.Presentation
		return presentation, nil
	} else {
		return domain.Presentation{}, errors.New(string(resp.Body()))
	}
}

func (w *Wallet) VerifyCredential(userId string, passphrase string, rawCredential json.RawMessage) (err error) {
	token, err := w.open(userId, passphrase)
	if err != nil {
		return err
	}

	defer w.close(userId, passphrase)

	resp, err := w.client.R().
		SetBody(struct {
			Auth       string          `json:"auth"`
			Credential json.RawMessage `json:"credential"`
			UsedId     string          `json:"usedId"`
		}{
			Auth:       token,
			Credential: rawCredential,
			UsedId:     userId,
		}).
		Post("/vcwallet/verify")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return errors.New(string(resp.Body()))
	}
}

func (w *Wallet) VerifyPresentation(userId string, passphrase string, rawPresentation json.RawMessage) error {
	token, err := w.open(userId, passphrase)
	if err != nil {
		return err
	}

	defer w.close(userId, passphrase)

	resp, err := w.client.R().
		SetBody(struct {
			Auth         string          `json:"auth"`
			Presentation json.RawMessage `json:"presentation"`
			UsedId       string          `json:"usedId"`
		}{
			Auth:         token,
			Presentation: rawPresentation,
			UsedId:       userId,
		}).
		Post("/vcwallet/verify")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return errors.New(string(resp.Body()))
	}
}

func (w *Wallet) open(userId string, passphrase string) (token string, err error) {
	var res map[string]interface{}
	resp, err := w.client.R().
		SetBody(struct {
			UserId             string `json:"userId"`
			LocalKMSPassphrase string `json:"localKMSPassphrase"`
		}{
			UserId:             userId,
			LocalKMSPassphrase: passphrase,
		}).
		SetResult(&res).
		Post("/vcwallet/open")

	if err != nil {
		return "", err
	}

	if resp.StatusCode() == http.StatusOK {
		return res["token"].(string), nil
	} else {
		return "", errors.New(string(resp.Body()))
	}
}

func (w *Wallet) close(userId string, passphrase string) (err error) {
	var res map[string]interface{}
	resp, err := w.client.R().
		SetBody(struct {
			UserId             string `json:"userId"`
			LocalKMSPassphrase string `json:"localKMSPassphrase"`
		}{
			UserId:             userId,
			LocalKMSPassphrase: passphrase,
		}).
		SetResult(&res).
		Post("/vcwallet/close")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return errors.New(string(resp.Body()))
	}
}

func makeRawCredential(credential domain.Credential) (rawCredential *json.RawMessage, err error) {
	var cred verifiable.Credential

	cred.ID = credential.CredentialId
	cred.Issuer = verifiable.Issuer{
		ID: credential.IssuerDID,
	}
	cred.Issued = util.NewTime(time.Now())
	cred.Context = []string{"https://www.w3.org/2018/credentials/v1", "https://ssimp.s3.amazonaws.com/schemas/prescription"}
	cred.Types = []string{"VerifiableCredential", credential.Type}

	var prescription map[string]interface{}
	err = json.Unmarshal(credential.Prescription.RawPrescription, &prescription)
	if err != nil {
		return nil, err
	}

	cred.Subject = verifiable.Subject{
		ID: credential.HolderDID,
		CustomFields: verifiable.CustomFields{
			"prescription": prescription,
		},
	}

	*rawCredential, err = cred.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return rawCredential, nil
}
