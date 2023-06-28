package main

import (
	"fmt"
	"go-architecture/handler"
	"go-architecture/pkg/databases"
	"go-architecture/repository"
	"go-architecture/services"
	"log"
	"net/http"
	"runtime"
)

func main() {
	// init Connection Database
	db, err := databases.ConnectPostgres(databases.Postgres{
		Host:     "host",
		Port:     "5432",
		User:     "user",
		Password: "password",
		DBName:   "dbname",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatal(err)
	}

	if db == nil {
		log.Fatal(err)
	}

	log.Println("Database Postgree Connected")

	// init Patient Repository
	patientRepo := repository.PatientRegistrationRepo(db)

	// init Patient Service
	patientSvc := services.PatientRegistrationService(patientRepo)
	patientHandler := handler.PatientRegistrationHandler(patientSvc)

	http.HandleFunc("/patients/add", patientHandler.Registration)
	http.HandleFunc("/patients/list", patientHandler.List)

	log.Println("Server Running at port 9090")
	http.ListenAndServe(":9090", nil)
	// Built()
}

func Built() {
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Printf("The application was built with the Go version: %s with Processor: %d CPU\n", runtime.Version(), runtime.NumCPU())
}
