package cmd

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"mail-service/ctrl"
	"net/http"

	"github.com/unrolled/render"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("User Sign In Handler Start!!")
	defer log.Println("User Sign In Handler End!!")
	ctx := context.Background()
	renderer := render.New()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		renderer.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	defer r.Body.Close()
	var body *ctrl.UserSignInRequestBody

	if err := json.Unmarshal(requestBody, &body); err != nil {
		renderer.JSON(w, http.StatusBadRequest, map[string]string{"error": "error unmarshalling request body"})
		return
	}

	flow := &ctrl.UserSignInFlow{
		RequestBody: *body,
		Context:     ctx,
	}

	result := flow.Run()

	if result.Error != nil {
		renderer.JSON(w, http.StatusBadRequest, map[string]string{"error": result.Error.Error()})
		return
	}

	renderer.JSON(w, http.StatusOK, result)
}
