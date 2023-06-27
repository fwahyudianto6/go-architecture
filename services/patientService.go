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

// implement Registration Service on  handler
func (p patientService) RegistrationService(req dto.PatientRegistrationRequest) (err error) {
	if len(req.FirstName) < 4 {
		return ErrFirstNameLength
	}
	if len(req.MiddleName) < 3 {
		return ErrMiddleNameLength
	}

	return nil
}
