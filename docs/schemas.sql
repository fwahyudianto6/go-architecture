-- DROP TABLE patients

CREATE TABLE patients (
    -- id SERIAL PRIMARY KEY,
    PatientId varchar(4),
    MedicalNo varchar(15),
    FirstName varchar(15),
    MiddleName varchar(15),
    LastName varchar(15),
    IsVoid boolean,
    Age int,
    CreateDate timestamptz DEFAULT now()
)

-- SELECT * FROM patients
