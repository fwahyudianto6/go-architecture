package handler

import (
	"encoding/json"
	"go-architecture/dto"
	"net/http"
)

// access Patient Service
type patientService interface {
	RegistrationService(req dto.PatientRegistrationRequest) (err error)
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
		rw.Write([]byte("METHOD NOT ALLOWED"))
	}
}
