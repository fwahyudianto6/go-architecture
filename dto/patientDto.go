package dto

//	declare Registration Patient [Request] construct
type PatientRegistrationRequest struct {
	PatientId  string `json:"PatientId"`
	MedicalNo  string `json:"MedicalNo"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName   string `json:"LastName"`
	Age        int    `json:"Age"`
}
