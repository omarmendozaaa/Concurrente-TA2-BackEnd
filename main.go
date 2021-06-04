package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//Rutas y Endpoints
	r.HandleFunc("/alumnos", GetAlumnos).Methods("GET")
	r.HandleFunc("/alumnos/{id}", GetAlumno).Methods("GET")
	r.HandleFunc("/alumnos", CreateAlumno).Methods("POST")
	r.HandleFunc("/alumnos/{id}", UpdateAlumno).Methods("PUT")
	r.HandleFunc("/alumnos/{id}", DeleteAlumno).Methods("DELETE")
	//Iniciar Servidor
	log.Fatal(http.ListenAndServe(":8080", r))
}
