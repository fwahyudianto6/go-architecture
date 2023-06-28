package repository

import (
	"database/sql"
	"go-architecture/entity"
)

// declare Patient Repository construct
type patientRepository struct {
	db *sql.DB
}

// declare Patient Registration Repository
func PatientRegistrationRepo(p_oDatabase *sql.DB) patientRepository {
	return patientRepository{db: p_oDatabase}
}

// implements Patient Registration Repository on Service
func (p patientRepository) RegistrationRepo(patients entity.Patient) (p_strPatientId string, err error) {
	query := `
        INSERT INTO patients (PatientId, MedicalNo, FirstName, MiddleName, LastName, IsVoid, Age)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING PatientId
    `
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(
		patients.PatientId, patients.MedicalNo,
		patients.FirstName, patients.MiddleName, patients.LastName,
		patients.IsVoid, patients.Age)
	err = row.Scan(&p_strPatientId)

	return
}
