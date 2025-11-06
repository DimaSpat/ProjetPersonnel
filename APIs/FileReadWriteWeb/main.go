package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type WriteRequest struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

type JSONResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func main() {
	http.HandleFunc("/read", handleReadFile)
	http.HandleFunc("/write", handleWriteFile)

	port := "8080"
	log.Printf("Starting server on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Impossible to start the server : %v\n", err)
	}
}

func handleReadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}

	log.Printf("Reading file '%s'...", filename)

	content, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("File '%s' does not exist", filename)
			sendJSONError(w, "File does not exist", http.StatusNotFound)
		} else {
			log.Printf("Error reading file '%s': %v", filename, err)
			sendJSONError(w, "Error reading file", http.StatusInternalServerError)
		}
		return
	}

	responseData := map[string]string{
		"filename": filename,
		"content":  string(content),
	}
	sendJSONSuccess(w, "File read successfully", responseData)
}

func handleWriteFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		sendJSONError(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	var req WriteRequest
	if err := json.Unmarshal(body, &req); err != nil {
		log.Printf("Error unmarshalling request: %v", err)
		sendJSONError(w, "Error unmarshalling request", http.StatusInternalServerError)
		return
	}

	if req.Filename == "" || req.Content == "" {
		sendJSONError(w, "Filename and content are required", http.StatusBadRequest)
		return
	}

	log.Printf("Writing file '%s'...", req.Filename)

	err = os.WriteFile(req.Filename, []byte(req.Content), 0666)
	if err != nil {
		log.Printf("Error writing file: %v", err)
		sendJSONError(w, "Error writing file", http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("File '%s' successfully written", req.Filename)
	sendJSONSuccess(w, message, nil)
}

func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(JSONResponse{
		Success: false,
		Message: message,
	})
	if err != nil {
		return
	}
}

func sendJSONSuccess(w http.ResponseWriter, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(JSONResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
	if err != nil {
		return
	}
}
