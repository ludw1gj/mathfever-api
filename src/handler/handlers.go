package handler

import (
	"encoding/json"
	"fmt"
	"gitlab.com/ludw1gj/mathfever-api/src/api"
	"log"
	"net/http"
	"strings"
)

type CalculationResponse struct {
	Content string `json:"answer"`
}

type ErrorJSONResponse struct {
	Error string `json:"error"`
}

func CategoriesHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if _, err := w.Write([]byte(api.GetCategoriesJson())); err != nil {
		log.Println(err)
	}
}

func CalculationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	calculationSlug := strings.Replace(r.URL.Path, "/calculation/", "", 1)
	calculation, err := api.FindMath(calculationSlug)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ErrorJSONResponse{"invalid slug"})
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&calculation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorJSONResponse{fmt.Sprintf("Malformed JSON request - %v", err.Error())})
		return
	}

	payload, err := calculation.ExecuteMath()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorJSONResponse{fmt.Sprintf("Could not execute math - %v", err.Error())})
		return
	}

	respBody, err := json.Marshal(CalculationResponse{Content: payload})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorJSONResponse{"failed to marshal payload"})
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(respBody); err != nil {
		log.Println(err)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ErrorJSONResponse{"invalid path - please request valid route"})
}
