package entity

// declare Patient Entity
type Patient struct {
	PatientId  string `json:"PatientId"`
	MedicalNo  string `json:"MedicalNo"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName   string `json:"LastName"`
	IsVoid     bool   `json:"IsVoid"`
	Age        int    `json:"Age"`
}
