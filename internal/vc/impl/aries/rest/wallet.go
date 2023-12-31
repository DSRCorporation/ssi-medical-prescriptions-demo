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
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/btcsuite/btcutil/base58"
	"github.com/go-resty/resty/v2"
	"github.com/hyperledger/aries-framework-go/pkg/controller/command/vcwallet"
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

	rawCredential, err := toRawCredential(credential)
	if err != nil {
		return domain.Credential{}, err
	}

	var res struct {
		Credential json.RawMessage `json:"credential"`
	}
	resp, err := w.client.R().
		SetBody(&vcwallet.IssueRequest{
			WalletAuth: vcwallet.WalletAuth{UserID: userId, Auth: token},
			Credential: rawCredential,
			ProofOptions: &wallet.ProofOptions{
				Controller: did,
			}}).
		SetResult(&res).
		Post("/vcwallet/issue")

	if err != nil {
		return domain.Credential{}, err
	}

	if resp.StatusCode() == http.StatusOK {
		credential.RawCredential = res.Credential
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

	rawPresentation, err := toRawPresentation(presentation)
	if err != nil {
		return domain.Presentation{}, err
	}

	var res struct {
		Presentation json.RawMessage `json:"presentation"`
	}

	resp, err := w.client.R().
		SetBody(&vcwallet.ProveRequest{
			WalletAuth:   vcwallet.WalletAuth{UserID: userId, Auth: token},
			Presentation: rawPresentation,
			ProofOptions: &wallet.ProofOptions{
				Controller: did,
			}}).
		SetResult(&res).
		Post("/vcwallet/prove")

	if err != nil {
		return domain.Presentation{}, err
	}

	if resp.StatusCode() == http.StatusOK {
		presentation.RawPresentation = res.Presentation
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

	body := vcwallet.VerifyRequest{
		WalletAuth:    vcwallet.WalletAuth{UserID: userId, Auth: token},
		RawCredential: rawCredential,
	}

	return w.Verify(body)
}

func (w *Wallet) VerifyPresentation(userId string, passphrase string, rawPresentation json.RawMessage) error {
	token, err := w.open(userId, passphrase)
	if err != nil {
		return err
	}

	defer w.close(userId, passphrase)

	body := vcwallet.VerifyRequest{
		WalletAuth:   vcwallet.WalletAuth{UserID: userId, Auth: token},
		Presentation: rawPresentation,
	}

	return w.Verify(body)
}

func (w *Wallet) Verify(body any) error {
	var res vcwallet.VerifyResponse
	resp, err := w.client.R().
		SetResult(&res).
		SetBody(body).
		Post("/vcwallet/verify")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		if res.Verified == true {
			return nil
		} else {
			return fmt.Errorf(res.Error)
		}
	} else {
		return fmt.Errorf(resp.String())
	}
}

func (w *Wallet) WalletExists(userId string) bool {
	resp, err := w.client.R().
		Get(fmt.Sprintf("/vcwallet/profile/%s", userId))

	if err != nil {
		return false
	}

	if resp.StatusCode() == http.StatusOK {
		return true
	}

	return false
}

func (w *Wallet) CreateWallet(userId string, passphrase string) (err error) {
	resp, err := w.client.R().
		SetBody(vcwallet.CreateOrUpdateProfileRequest{
			UserID:             userId,
			LocalKMSPassphrase: passphrase,
		}).
		Post("/vcwallet/create-profile")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}

func (w *Wallet) AddKey(userId string, passphrase string, keyId string, privKey ed25519.PrivateKey) (err error) {
	token, err := w.open(userId, passphrase)
	if err != nil {
		return err
	}

	defer w.close(userId, passphrase)

	ct := struct {
		Id               string `json:"id"`
		Type             string `json:"type"`
		PrivateKeyBase58 string `json:"privateKeyBase58"`
	}{
		Id:               keyId,
		Type:             "ed25519verificationkey2018",
		PrivateKeyBase58: base58.Encode(privKey),
	}

	ctRaw, err := json.Marshal(ct)
	if err != nil {
		return err
	}

	resp, err := w.client.R().
		SetBody(&vcwallet.AddContentRequest{
			WalletAuth: vcwallet.WalletAuth{
				Auth:   token,
				UserID: userId,
			},
			Content:     json.RawMessage(ctRaw),
			ContentType: wallet.ContentType("key"),
		}).
		Post("/vcwallet/add")

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}

func (w *Wallet) open(userId string, passphrase string) (token string, err error) {
	var res map[string]interface{}
	resp, err := w.client.R().
		SetBody(&vcwallet.UnlockWalletRequest{
			UserID:             userId,
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
		SetBody(&vcwallet.LockWalletRequest{
			UserID: userId,
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
