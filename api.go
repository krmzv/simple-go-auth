package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/user/{id}", makeHTTPHandleFunc(s.handeGetUserByID))
	log.Println("JSON API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetUsers(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteUser(w, r)
	}
	return fmt.Errorf("Method not allowed %s", r.Method)
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.store.GetUsers()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, users)
}

func (s *APIServer) handeGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	//user := NewUser("Mihajlo", "J", "m@gmail.com")
	return WriteJSON(w, http.StatusOK, &User{})
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserRequest := new(CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&createUserRequest); err != nil {
		return err
	}
	user := NewUser(createUserRequest.FirstName, createUserRequest.LastName, createUserRequest.Email)
	if err := s.store.CreateUser(user); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}
