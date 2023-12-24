package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
}

type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)

	if err != nil {
		http.Error(w, "Error while Decoding", http.StatusBadRequest)
		return
	}

	if requestBody.Message == "" || requestBody.Message != "message" {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received message: %s\n", requestBody.Message)

	responseBody := ResponseBody{
		Status:  "success",
		Message: "Data successfully received",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseBody)
}

func main() {
	http.HandleFunc("/", handlePostRequest)
	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
