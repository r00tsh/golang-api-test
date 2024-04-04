package main

import (
	"encoding/json"
	_ "main/docs"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// HealthCheck godoc
//
//	@Summary	health
//	@Schemes
//	@Description	do health
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Status
//	@Router			/health [get]
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(status)
}

func main() {
	http.HandleFunc("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	http.HandleFunc("/health", HealthHandler)

	http.ListenAndServe(":8080", nil)
}
