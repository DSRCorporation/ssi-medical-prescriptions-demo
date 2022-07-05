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

	"github.com/go-resty/resty/v2"
	"github.com/hyperledger/aries-framework-go/pkg/controller/command/vcwallet"
	"github.com/hyperledger/aries-framework-go/pkg/wallet"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/domain"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/vc"
)

type Wallet struct {
	client   *resty.Client
	endpoint string
}

func NewWallet(endpoint string) (*Wallet, error) {
	return &Wallet{client: resty.New(), endpoint: endpoint}, nil
}

func (w *Wallet) SignCredential(userId string, passphrase string, proofOptions vc.ProofOptions, credential domain.Credential) (domain.Credential, error) {
	token, err := w.open(userId, passphrase)
	if err != nil {
		return domain.Credential{}, err
	}

	defer w.close(userId, passphrase)

	var res struct {
		RawCredential json.RawMessage
	}
	resp, err := w.client.R().
		SetBody(&vcwallet.IssueRequest{
			WalletAuth: vcwallet.WalletAuth{UserID: userId, Auth: token},
			Credential: credential.RawCredential,
			ProofOptions: &wallet.ProofOptions{
				Controller:          proofOptions.Controller,
				VerificationMethod:  proofOptions.VerificationMethod,
				Created:             proofOptions.Created,
				Domain:              proofOptions.Domain,
				Challenge:           proofOptions.Challenge,
				ProofType:           proofOptions.ProofType,
				ProofRepresentation: proofOptions.ProofRepresentation,
			}}).
		SetResult(&res).
		Post(w.endpoint + "/vcwallet/issue")

	if err != nil {
		return domain.Credential{}, err
	}

	if resp.StatusCode() == http.StatusOK {
		credential.RawCredential = res.RawCredential
		return credential, nil
	} else {
		return domain.Credential{}, errors.New(string(resp.Body()))
	}
}

func (w *Wallet) SignPresentation(userId string, passphrase string, proofOptions vc.ProofOptions, presentaion domain.Presentation) (domain.Presentation, error) {
	token, err := w.open(userId, passphrase)
	if err != nil {
		return domain.Presentation{}, err
	}

	defer w.close(userId, passphrase)

	var res struct {
		RawPresentation json.RawMessage
	}
	resp, err := w.client.R().
		SetBody(&vcwallet.ProveRequest{
			WalletAuth:   vcwallet.WalletAuth{UserID: userId, Auth: token},
			Presentation: presentaion.RawPresentation,
			ProofOptions: &wallet.ProofOptions{
				Controller:          proofOptions.Controller,
				VerificationMethod:  proofOptions.VerificationMethod,
				Created:             proofOptions.Created,
				Domain:              proofOptions.Domain,
				Challenge:           proofOptions.Challenge,
				ProofType:           proofOptions.ProofType,
				ProofRepresentation: proofOptions.ProofRepresentation,
			}}).
		SetResult(&res).
		Post(w.endpoint + "/vcwallet/prove")

	if err != nil {
		return domain.Presentation{}, err
	}

	if resp.StatusCode() == http.StatusOK {
		presentaion.RawPresentation = res.RawPresentation
		return presentaion, nil
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
		Post(w.endpoint + "/vcwallet/verify")

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
		Post(w.endpoint + "/vcwallet/verify")

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
		Post(w.endpoint + "/vcwallet/open")

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
		Post(w.endpoint + "/vcwallet/close")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	} else {
		return errors.New(string(resp.Body()))
	}
}
