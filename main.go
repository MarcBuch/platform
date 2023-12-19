package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", clusterHandler)
	http.ListenAndServe(":8080", nil)
}

func clusterHandler(w http.ResponseWriter, r *http.Request) {
	reader := CueReader{}
	values, err := reader.ReadFile("test.cue")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	reader.values = values
	clusterConfig := mapConfig(reader)
	fmt.Println(clusterConfig)

	jsonData, err := json.Marshal(clusterConfig)
	if err != nil {
		fmt.Println(err) // TODO: Implement logging
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%s\n", jsonData)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
