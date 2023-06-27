package services

import (
	"errors"
	"go-architecture/dto"
)

// declare validate length
var (
	ErrFirstNameLength  = errors.New("patient first name minimum 4 length")
	ErrMiddleNameLength = errors.New("patient middle name minimum 3 length")
)

// declare Patient Service construct
type patientService struct {
}

// declare Patient Registration Service
func PatientRegistrationService() patientService {
	return patientService{}
}

// #region Business Logic Method

// implement Registration Service on handler
func (p patientService) RegistrationService(req dto.PatientRegistrationRequest) (err error) {
	if len(req.FirstName) < 4 {
		return ErrFirstNameLength
	}
	if len(req.MiddleName) < 3 {
		return ErrMiddleNameLength
	}

	return nil
}

// implement Patient List Service on handler
func (p patientService) ListService() (resp []dto.PatientListResponse, err error) {
	return []dto.PatientListResponse{{
		PatientId:  "P001",
		MedicalNo:  "002-000-01-99",
		FirstName:  "Fajar",
		MiddleName: "bin",
		LastName:   "Wahyudianto",
		Age:        23,
	}, {
		PatientId:  "P002",
		MedicalNo:  "002-000-01-98",
		FirstName:  "El",
		MiddleName: "bin",
		LastName:   "Wahyudi",
		Age:        13,
	}}, nil
}

// #endregion Business Logic Method
