package server

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Alumno struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	DNI    int    `json:"dni"`
	Edad   int    `json:"edad"`
}

var alumnos []Alumno

//Obtiene una lista de alumnos
func GetAlumnos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alumnos)
}

//Devuelve un alumno por id
func GetAlumno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range alumnos {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Alumno{})
}

//Crear un nuevo alumno
func CreateAlumno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var alumno Alumno
	_ = json.NewDecoder(r.Body).Decode(&alumno)
	alumno.ID = strconv.Itoa(rand.Intn(100000))
	alumnos = append(alumnos, alumno)
	json.NewEncoder(w).Encode(alumno)
}

//Actualiza un alumnos
func UpdateAlumno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range alumnos {
		if item.ID == params["id"] {
			alumnos = append(alumnos[:index], alumnos[index+1:]...)
			var alumno Alumno
			_ = json.NewDecoder(r.Body).Decode(&alumno)
			alumno.ID = params["id"]
			alumnos = append(alumnos, alumno)
			json.NewEncoder(w).Encode(alumno)
			return
		}
	}
}

//Elimina un alumno
func DeleteAlumno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range alumnos {
		if item.ID == params["id"] {
			alumnos = append(alumnos[:index], alumnos[index+1:]...)
			// mutamos todos los datos de alumnos, pasando todo menos el encontrado
			// una vez encontrado hacemos break para salir del ciclo
			break
		}
	}
}
