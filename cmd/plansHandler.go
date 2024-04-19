package cmd

import (
	"log"
	"mail-service/ctrl"
	"net/http"

	"github.com/unrolled/render"
)

func GetPlansHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting to get plans...")
	defer log.Println("Finished getting plans...")

	renderer := render.New()
	var flow *ctrl.GetPlansFlow

	response := flow.Run()

	if response.Error != nil {
		log.Fatalf("error getting plans: %v", response.Error)
		renderer.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	renderer.JSON(w, http.StatusOK, response)
}
