package handler

import (
	"encoding/json"
	"go-architecture/dto"
	"net/http"
)

// access Patient Service
type patientService interface {
	RegistrationService(req dto.PatientRegistrationRequest) (err error)
	ListService() (patients []dto.PatientListResponse, err error)
}

// declare Patient Handler construct
type patientHandler struct {
	patientSvc patientService
}

// declare Patient Registration Handler
func PatientRegistrationHandler(p_oService patientService) patientHandler {
	return patientHandler{
		patientSvc: p_oService,
	}
}

// #region Business Logic Method

// implement Registration handler
func (p patientHandler) Registration(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		rw.Header().Set("Content-Type", "application/json")
		var req dto.PatientRegistrationRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		err = p.patientSvc.RegistrationService(req)
		if err != nil {
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"success": true,
			"message": "Patient Registration Success",
		})

	} else {
		rw.Write([]byte("Methot No Allowed!"))
	}
}

// implement Patient List handler
func (p patientHandler) List(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rw.Header().Set("Content-Type", "application/json")
		patients, err := p.patientSvc.ListService()

		if err != nil {
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"success": true,
			"message": "Patient List success",
			"data":    patients,
		})
	} else {
		rw.Write([]byte("Methot No Allowed!"))
	}
}

// #endregion Business Logic Method
