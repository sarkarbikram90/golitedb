package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Record struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

var database map[string]Record

func main() {
	database = make(map[string]Record)
	http.HandleFunc("/record", handleRecord)
	http.HandleFunc("/record/query", handleQuery)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRecord(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getRecord(w, r)
	case http.MethodPost:
		createRecord(w, r)
	case http.MethodPut:
		updateRecord(w, r)
	case http.MethodDelete:
		deleteRecord(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func getRecord(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	record, found := database[id]
	if found {
		json.NewEncoder(w).Encode(record)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Record not found")
	}
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	var record Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request payload")
		return
	}

	database[record.ID] = record
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Record created successfully")
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, found := database[id]
	if found {
		var record Record
		err := json.NewDecoder(r.Body).Decode(&record)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request payload")
			return
		}

		database[id] = record
		fmt.Fprintf(w, "Record updated successfully")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Record not found")
	}
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, found := database[id]
	if found {
		delete(database, id)
		fmt.Fprintf(w, "Record deleted successfully")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Record not found")
	}
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	result := make([]Record, 0)

	for _, record := range database {
		if strings.Contains(record.Data, query) {
			result = append(result, record)
		}
	}

	json.NewEncoder(w).Encode(result)
}
