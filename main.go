package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

func main() {
	log.Info("Starting server on port 8080")
	http.HandleFunc("/", clusterHandler)
	http.ListenAndServe(":8080", nil)
}

func clusterHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(fmt.Sprintf("Received %s request to %s from %s with %s", r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent()))
	reader := CueReader{}
	values, err := reader.ReadFile("test.cue")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	reader.values = values
	clusterConfig := mapConfig(reader)

	jsonData, err := json.Marshal(clusterConfig)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
