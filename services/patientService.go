package services

import (
	"errors"
	"go-architecture/dto"
	"go-architecture/entity"
	"log"
)

// declare validate length
var (
	ErrFirstNameLength  = errors.New("patient First Name minimum 4 length")
	ErrMiddleNameLength = errors.New("patient Middle Name minimum 3 length")
	ErrLastNameLength   = errors.New("patient Last Name minimum 3 length")
)

// access Patient Repository
type patientRepository interface {
	RegistrationRepo(patient entity.Patient) (p_strPatientId string, err error)
}

// declare Patient Service construct
type patientService struct {
	repo patientRepository
}

// declare Patient Registration Service
func PatientRegistrationService(p_oRepo patientRepository) patientService {
	return patientService{
		repo: p_oRepo,
	}
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

	id, err := p.repo.RegistrationRepo(req.ParseToEntity())
	if err != nil {
		return
	}

	log.Printf("Patient Registraton [%s] successfully!\n", id)

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
		IsVoid:     false,
		Age:        23,
	}, {
		PatientId:  "P002",
		MedicalNo:  "002-000-01-98",
		FirstName:  "El",
		MiddleName: "bin",
		LastName:   "Wahyudi",
		IsVoid:     false,
		Age:        13,
	}}, nil
}

// #endregion Business Logic Method
