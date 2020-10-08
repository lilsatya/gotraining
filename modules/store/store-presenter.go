package store

import (
	"encoding/json"
	"gotraining/entities"
	"gotraining/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Create will create new data
func Create(w http.ResponseWriter, r *http.Request) {
	var s entities.Store
	handler := newHandler()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Payload")
		return
	}
	defer r.Body.Close()

	store, err := handler.create(s)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, store)
}

// List will fetch all data
func List(w http.ResponseWriter, r *http.Request) {
	handler := newHandler()

	stores, err := handler.list()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, stores)
}

// Read will fetch single data
func Read(w http.ResponseWriter, r *http.Request) {
	handler := newHandler()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	store, err := handler.read(id)

	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, store)
}

// Update will update existing data based on ID
func Update(w http.ResponseWriter, r *http.Request) {
	var s entities.Store
	handler := newHandler()
	decoder := json.NewDecoder(r.Body)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := decoder.Decode(&s); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Payload")
		return
	}
	defer r.Body.Close()

	store, err := handler.update(id, s)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, store)
}

// Delete will soft delete existing data based on ID
func Delete(w http.ResponseWriter, r *http.Request) {
	handler := newHandler()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	defer r.Body.Close()

	if err := handler.delete(id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}
