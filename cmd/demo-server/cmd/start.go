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

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/controller/handler"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/controller/rest"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/service"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/storage/impl/leveldb"
	aries "github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/vc/impl/aries/rest"
)

const (
	// level db path
	levelDBStoragePathFlagName = "level-db-storage-path"
	levelDBStoragePathEnvKey   = "SSIMP_LEVELDB_STORAGE_PATH"

	// aries agent endpoints
	issuerRestEndpointFlagName = "issuer-rest-endpoint"
	issuerRestEndpointEnvKey   = "SSIMP_ISSUER_REST_ENDPOINT"

	holderRestEndpointFlagName = "holder-rest-endpoint"
	holderRestEndpointEnvKey   = "SSIMP_HOLDER_REST_ENDPOINT"

	verifierRestEndpointFlagName = "verifier-rest-endpoint"
	verifierRestEndpointEnvKey   = "SSIMP_VERIFIER_REST_ENDPOINT"
)

type ServerParameters struct {
	levelDBStoragePath   string
	issuerRestEndpoint   string
	holderRestEndpoint   string
	verifierRestEndpoint string
}

func Cmd() (*cobra.Command, error) {
	startCmd := createStartCMD()

	createFlags(startCmd)

	return startCmd, nil
}

//nolint:funlen
func createFlags(startCmd *cobra.Command) {
	startCmd.Flags().String(levelDBStoragePathFlagName, "", "")
	startCmd.Flags().String(issuerRestEndpointFlagName, "", "")
	startCmd.Flags().String(holderRestEndpointFlagName, "", "")
	startCmd.Flags().String(verifierRestEndpointFlagName, "", "")
}

// NewServerParameters constructs ServerParameters with the given cobra command.
func NewServerParameters(cmd *cobra.Command) (*ServerParameters, error) { //nolint: funlen,gocyclo

	levelDBStoragePath, err := getUserSetVar(cmd, levelDBStoragePathFlagName, levelDBStoragePathEnvKey, false)
	if err != nil {
		return nil, err
	}

	issuerRestEndpoint, err := getUserSetVar(cmd, issuerRestEndpointFlagName, issuerRestEndpointEnvKey, false)
	if err != nil {
		return nil, err
	}

	holderRestEndpoint, err := getUserSetVar(cmd, holderRestEndpointFlagName, holderRestEndpointEnvKey, false)
	if err != nil {
		return nil, err
	}

	verifierRestEndpoint, err := getUserSetVar(cmd, verifierRestEndpointFlagName, verifierRestEndpointEnvKey, false)
	if err != nil {
		return nil, err
	}

	return &ServerParameters{
		levelDBStoragePath:   levelDBStoragePath,
		issuerRestEndpoint:   issuerRestEndpoint,
		holderRestEndpoint:   holderRestEndpoint,
		verifierRestEndpoint: verifierRestEndpoint,
	}, nil
}

func initializeRestHandler(params *ServerParameters) (h *handler.RestHandler, err error) {
	// read user set variables
	// levelDBStoragePath := getUserSetVar()

	// initialize doctor service
	doctorStorage, err := leveldb.NewDoctorStorage(params.levelDBStoragePath + "/doctor")
	if err != nil {
		return nil, fmt.Errorf("error creating doctor storage: %v\n", err)
	}
	doctorService := service.NewDoctorService(doctorStorage)

	// initialize patient service
	patientStorage, err := leveldb.NewPatientStorage(params.levelDBStoragePath + "/patient")
	if err != nil {
		return nil, fmt.Errorf("error creating patient storage: %v\n", err)
	}
	patientService := service.NewPatientService(patientStorage)

	// initialize pharmacy service
	pharmacyStorage, err := leveldb.NewPharmacyStorage(params.levelDBStoragePath + "/pharmacy")
	if err != nil {
		return nil, fmt.Errorf("error creating pharmacy storage: %v\n", err)
	}

	pharmacyService := service.NewPharmacyService(pharmacyStorage)

	// initialize vc service
	vcStorage, err := leveldb.NewVCStorage(params.levelDBStoragePath + "/vc")
	if err != nil {
		return nil, fmt.Errorf("error creating vc storage: %v\n", err)
	}

	holderAgent, err := aries.NewHolder(params.holderRestEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error creating vc holder agent: %v\n", err)
	}

	issuerAgent, err := aries.NewIssuer(params.issuerRestEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error creating vc issuer agent: %v\n", err)
	}

	verifierAgent, err := aries.NewVerifier(params.verifierRestEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error creating vc verifier agent: %v\n", err)
	}

	// Use holder agent for VDR for simplicity
	// @TODO: may need dedicated agent in future
	vdr, err := aries.NewVDR(params.holderRestEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error creating vc verifier agent: %v\n", err)
	}

	issuerWallet, err := aries.NewWallet(params.issuerRestEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error creating issuer wallet: %v\n", err)
	}

	holderWallet, err := aries.NewWallet(params.holderRestEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error creating holder wallet: %v\n", err)
	}

	verifierWallet, err := aries.NewWallet(params.verifierRestEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error creating holder wallet: %v\n", err)
	}

	vcService := service.NewVCService(vcStorage, issuerAgent, holderAgent, verifierAgent, vdr, issuerWallet, holderWallet, verifierWallet)

	// initialize rest hander
	h = handler.New(doctorService, patientService, pharmacyService, vcService)

	return
}

func getUserSetVar(cmd *cobra.Command, flagName, envKey string, isOptional bool) (string, error) {
	if cmd != nil && cmd.Flags().Changed(flagName) {
		value, err := cmd.Flags().GetString(flagName)
		if err != nil {
			return "", fmt.Errorf(flagName+" flag not found: %s", err)
		}

		return value, nil
	}

	value, isSet := os.LookupEnv(envKey)

	if isOptional || isSet {
		return value, nil
	}

	return "", errors.New("Neither " + flagName + " (command line flag) nor " + envKey +
		" (environment variable) have been set.")
}

func createStartCMD() *cobra.Command {
	return &cobra.Command{
		Use: "start",
		RunE: func(cmd *cobra.Command, args []string) error {
			parameters, err := NewServerParameters(cmd)
			if err != nil {
				return err
			}
			return startServer(parameters)
		},
	}
}

func startServer(parameters *ServerParameters) error {
	// initialize rest handler
	restHandler, err := initializeRestHandler(parameters)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initializing rest handler: %v\n", err)
		os.Exit(1)
	}

	// start echo server
	e := echo.New()
	rest.RegisterHandlers(e, restHandler)

	return e.Start(":8888")
}
