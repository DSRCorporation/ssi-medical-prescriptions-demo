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

package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/controller/handler"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/controller/rest"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/service"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/storage/impl/leveldb"
	aries "github.com/DSRCorporation/ssi-medical-prescriptions-demo/internal/vc/impl/aries/rest"
)

func main() {

	// initialize rest handler
	restHandler, err := initializeRestHandler()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initializing rest handler: %v\n", err)
		os.Exit(1)
	}

	// start echo server
	e := echo.New()
	rest.RegisterHandlers(e, restHandler)

	e.Logger.Fatal(e.Start(":8888"))
}

func initializeRestHandler() (h *handler.RestHandler, err error) {
	// initialize doctor service
	doctorStorage, err := leveldb.NewDoctorStorage("path/to/doctorStorage")
	if err != nil {
		return nil, fmt.Errorf("error creating doctor storage: %v\n", err)
	}
	doctorService := service.NewDoctorService(doctorStorage)

	// initialize patient service
	patientStorage, err := leveldb.NewPatientStorage("path/to/patientStorage")
	if err != nil {
		return nil, fmt.Errorf("error creating patient storage: %v\n", err)
	}
	patientService := service.NewPatientService(patientStorage)

	// initialize pharmacy service
	pharmacyStorage, err := leveldb.NewPharmacyStorage("path/to/pharmacyStorage")
	if err != nil {
		return nil, fmt.Errorf("error creating pharmacy storage: %v\n", err)
	}

	pharmacyService := service.NewPharmacyService(pharmacyStorage)

	// initialize vc service
	vcStorage, err := leveldb.NewVCStorage("path/to/vcStorage")
	if err != nil {
		return nil, fmt.Errorf("error creating vc storage: %v\n", err)
	}

	holderAgent, err := aries.NewHolder("http://holder.endpoint")
	if err != nil {
		return nil, fmt.Errorf("error creating vc holder agent: %v\n", err)
	}

	issuerAgent, err := aries.NewIssuer("http://issuer.endpoint")
	if err != nil {
		return nil, fmt.Errorf("error creating vc issuer agent: %v\n", err)
	}

	verifierAgent, err := aries.NewVerifier("http://verifier.endpoint")
	if err != nil {
		return nil, fmt.Errorf("error creating vc verifier agent: %v\n", err)
	}

	wallet, err := aries.NewWallet("http://wallet.endpoint")
	if err != nil {
		return nil, fmt.Errorf("error creating vc wallet: %v\n", err)
	}

	vcService := service.NewVCService(vcStorage, issuerAgent, holderAgent, verifierAgent, wallet)

	// initialize rest hander
	h = handler.New(doctorService, patientService, pharmacyService, vcService)

	return
}
