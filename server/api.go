package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}
type Server interface {
	Router() http.Handler
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}
func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}
func New() Server {

	a := &api{}
	r := mux.NewRouter()
	enableCORS(r)
	//Empty Database
	alumnos = append(alumnos, Alumno{ID: "1", Nombre: "Omar Mendoza", DNI: 33429504, Edad: 80})
	alumnos = append(alumnos, Alumno{ID: "2", Nombre: "Roman Ramirez", DNI: 72511063, Edad: 57})
	alumnos = append(alumnos, Alumno{ID: "3", Nombre: "Roberto Gampi", DNI: 12318290, Edad: 35})
	alumnos = append(alumnos, Alumno{ID: "4", Nombre: "Tela Meto El Jueves", DNI: 34298424, Edad: 22})
	alumnos = append(alumnos, Alumno{ID: "5", Nombre: "Lo mismo para tu prima", DNI: 33407982, Edad: 19})
	//Rutas y Endpoints
	r.HandleFunc("/alumnos", GetAlumnos).Methods("GET")
	r.HandleFunc("/alumnos/{id}", GetAlumno).Methods("GET")
	r.HandleFunc("/alumnos", CreateAlumno).Methods("POST")
	r.HandleFunc("/alumnos/{id}", UpdateAlumno).Methods("PUT")
	r.HandleFunc("/alumnos/{id}", DeleteAlumno).Methods("DELETE")
	//Iniciar Servidor
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
