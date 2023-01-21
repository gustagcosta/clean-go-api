package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gustagcosta/go-api/storage"
	"github.com/gustagcosta/go-api/types"
	"github.com/urfave/negroni"
)

type Server struct {
	port    string
	storage storage.Storage
}

func NewServer(port string, s storage.Storage) *Server {
	return &Server{
		port:    port,
		storage: s,
	}
}

func (s *Server) Start() error {
	router := mux.NewRouter()

	router.
		Methods("GET").
		Path("/dogs").
		HandlerFunc(s.handleGetDogs)

	router.
		Methods("POST").
		Path("/dogs").
		HandlerFunc(s.handleStoreDog)

	router.
		Methods("GET").
		Path("/dogs/{id}").
		HandlerFunc(s.handleGetDog)

	router.
		Methods("PUT").
		Path("/dogs").
		HandlerFunc(s.handleUpdateDog)

	router.
		Methods("DELETE").
		Path("/dogs/{id}").
		HandlerFunc(s.handleDeleteDog)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(ContentTypeJson))
	n.UseHandler(router)

	return http.ListenAndServe(s.port, n)
}

func (s *Server) handleGetDogs(w http.ResponseWriter, r *http.Request) {
	dogs, err := s.storage.GetDogs()

	if err != nil {
		negroni.NewLogger().ALogger.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	if len(*dogs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dogs)
}

func (s *Server) handleGetDog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if len(id) == 0 {
		http.Error(w, "id not provided", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "id is not valid", http.StatusBadRequest)
		return
	}

	dog, err := s.storage.GetDog(idInt)

	if err != nil {
		negroni.NewLogger().ALogger.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	if dog == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dog)
}

func (s *Server) handleStoreDog(w http.ResponseWriter, r *http.Request) {
	var req types.DogStoreRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		negroni.NewLogger().ALogger.Println(err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	id, err := s.storage.StoreDog(req.Name, req.Age)
	if err != nil {
		negroni.NewLogger().ALogger.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	idReturn := &types.IdReturn{
		ID: id,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(idReturn)
}

func (s *Server) handleDeleteDog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if len(id) == 0 {
		http.Error(w, "id not provided", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "id is not valid", http.StatusBadRequest)
		return
	}

	err = s.storage.DeleteDog(idInt)
	if err != nil {
		negroni.NewLogger().ALogger.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleUpdateDog(w http.ResponseWriter, r *http.Request) {
	var req types.Dog

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		negroni.NewLogger().ALogger.Println(err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := s.storage.UpdateDog(&req); err != nil {
		negroni.NewLogger().ALogger.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
