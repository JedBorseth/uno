package uno

import (
	"fmt"
	"net/http"
)


func ws(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request:", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := `{"message": "Hello, WebSocket!"}`
	_, err := w.Write([]byte(response))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("Response sent successfully")
}