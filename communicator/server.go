package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	registerRoutes()
	http.ListenAndServe(":8080", nil)
}

func registerRoutes() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/completed", doneHandler)
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func doneHandler(w http.ResponseWriter, req *http.Request) {
	testRes := struct {
		Status     string   `json:"status"`
		IdeaID     string   `json:"idea_id"`
		Author     string   `json:"author"`
		Commenters []string `json:"commenters"`
	}{
		"Done",
		"12345",
		"Haydes",
		[]string{"Beep_Boop"},
	}
	res, err := json.Marshal(testRes)
	if err != nil {
		sendFailResponse(w)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
}

func sendFailResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}
