package cmd

import (
	"encoding/json"
	"io"
	"log"
	"mail-service/ctrl"
	"net/http"

	"github.com/unrolled/render"
)

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting to get plans...")
	defer log.Println("Finished getting plans...")

	renderer := render.New()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		renderer.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	defer r.Body.Close()
	var body *ctrl.SendMailBody

	if err := json.Unmarshal(requestBody, &body); err != nil {
		renderer.JSON(w, http.StatusBadRequest, map[string]string{"error": "error unmarshalling request body"})
		return
	}

	flow := &ctrl.SendMailFlow{
		Body: *body,
	}

	response := flow.Run()

	if response.Error != nil {
		log.Fatalf("error sending mail : %v", response.Error)
		renderer.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	renderer.JSON(w, http.StatusOK, response)
}
