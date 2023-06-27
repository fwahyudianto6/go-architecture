package main

import (
	"fmt"
	"go-architecture/handler"
	"go-architecture/services"
	"log"
	"net/http"
	"runtime"
)

func main() {
	patientSvc := services.PatientRegistrationService()
	patientHandler := handler.PatientRegistrationHandler(patientSvc)

	http.HandleFunc("/patients/add", patientHandler.Registration)

	log.Println("Server Running at port 9090")
	http.ListenAndServe(":9090", nil)
	// Built()
}

func Built() {
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Printf("The application was built with the Go version: %s with Processor: %d CPU\n", runtime.Version(), runtime.NumCPU())
}
