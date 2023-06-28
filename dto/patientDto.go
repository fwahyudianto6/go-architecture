package dto

import "go-architecture/entity"

//	declare Registration Patient [Request] construct
type PatientRegistrationRequest struct {
	PatientId  string `json:"PatientId"`
	MedicalNo  string `json:"MedicalNo"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName   string `json:"LastName"`
	IsVoid     bool   `json:"IsVoid"`
	Age        int    `json:"Age"`
}

//  declare Patient List data [Response] construct
type PatientListResponse struct {
	PatientId  string `json:"PatientId"`
	MedicalNo  string `json:"MedicalNo"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName   string `json:"LastName"`
	IsVoid     bool   `json:"IsVoid"`
	Age        int    `json:"Age"`
}

//  #region Method

// convert Object to Entity
func (p PatientRegistrationRequest) ParseToEntity() entity.Patient {
	return entity.Patient{
		PatientId:  p.PatientId,
		MedicalNo:  p.MedicalNo,
		FirstName:  p.FirstName,
		MiddleName: p.MiddleName,
		LastName:   p.LastName,
		IsVoid:     p.IsVoid,
		Age:        p.Age,
	}
}

// #endregion Method
