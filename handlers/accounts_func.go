package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// List returns all accounts.
func (s *Accounts) ListFunc(rw http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		limit = 0
	}

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		offset = 0
	}

	res, err := s.service.List(limit, offset)

	if err != nil {
		handleError(rw, s.log, err)
		return
	}

	handleJsonResponse(rw, http.StatusOK, res)
}

// Create creates a new account asynchronously.
// It returns a Job JSON representation.
func (s *Accounts) CreateFunc(rw http.ResponseWriter, r *http.Request) {
	// Decide whether to serve sync or async, default async
	sync := r.FormValue(SyncQueryParameter) != ""

	job, acc, err := s.service.Create(r.Context(), sync)

	if err != nil {
		handleError(rw, s.log, err)
		return
	}

	var res interface{}
	if sync {
		res = acc
	} else {
		res = job.ToJSONResponse()
	}

	handleJsonResponse(rw, http.StatusCreated, res)
}

// Details returns details regarding an account.
// It reads the address for the wanted account from URL.
// Account service is responsible for validating the address.
func (s *Accounts) DetailsFunc(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res, err := s.service.Details(vars["address"])

	if err != nil {
		handleError(rw, s.log, err)
		return
	}

	handleJsonResponse(rw, http.StatusOK, res)
}
