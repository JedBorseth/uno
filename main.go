package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"uno/api"
)

func main() {
	// Vercel Go applications do not require a main function.
	// This is just for development purposes.

 
	http.HandleFunc("/api", api.Handler)
	http.HandleFunc("/api/ws", api.WS)
	if os.Getenv("VERCEL_ENV") != "production" {
		cmd := exec.Command("pnpm", "run", "build")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to run npm build: %v\n", err)
		}
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.FileServer(http.Dir("frontend-build")).ServeHTTP(w, r)
	})
	
	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
	// Note: In production, Vercel will handle the routing and serving of the application.
}

