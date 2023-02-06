package api

import (
	"api-cv/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func (api *API) EducationList(w http.ResponseWriter, r *http.Request) {
	res, err := api.educationsRepo.ReadEducation()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	if len(res) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Education list not found!"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (api *API) AddEducation(w http.ResponseWriter, r *http.Request) {

	var education model.Education
	err := json.NewDecoder(r.Body).Decode(&education)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Bad Request"})
		return
	}

	err = api.educationsRepo.AddEducation(education)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "Education Added"})
}

func (api *API) DeleteEducation(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Bad Request"})
		return
	}

	err = api.educationsRepo.DeleteEducation(uint(idInt))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "Delete Success"})
}

func (api *API) UpdateEducation(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Bad Request"})
		return
	}

	var education model.Education
	err = json.NewDecoder(r.Body).Decode(&education)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Bad Request"})
		return
	}

	err = api.educationsRepo.UpdateEducation(uint(idInt), education)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "Update Success"})
}
