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

package domain

type Drug struct {
	DrugName           *string
	DrugNumber         *float32
	DrugType           *string
	RefillAvailability *bool

	TabletDrugInfo *struct {
		DaysOfMedication     *float32
		NumberOfDoses        *string
		NumberOfDrug         *float32
		TimingToTakeMedicine *string
	}

	OintmentDrugInfo *struct {
		AmountOfMedicine     *float32
		ApplicationFrequency *string
		Usage                *string
		UseArea              *string
	}
}

type Prescription struct {
	Drugs        *[]Drug
	HospitalInfo *struct {
		DoctorName               *string
		DoctorSignatureStamp     *string
		HospitalName             *string
		Location                 *string
		MedicalInstitutionNumber *float32
		Phone                    *string
		PrefectureNumber         *float32
		ScoreVoteNumber          *float32
	}
	IssuanceInfo *struct {
		ExpirationDate *string
		IssuanceDate   *string
	}
	PatientInfo *struct {
		Birthday *string
		Name     *string
		Sex      *string
	}
}
