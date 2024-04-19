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

func SubscribePlanHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Starting user subscription to plan!!")
	defer log.Println("User Subscription End!!")

	ctx := context.Background()
	renderer := render.New()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		renderer.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	defer r.Body.Close()
	var body *ctrl.UserSignUpRequestBody

	if err := json.Unmarshal(requestBody, &body); err != nil {
		renderer.JSON(w, http.StatusBadRequest, map[string]string{"error": "error unmarshalling request body"})
		return
	}

	flow := &ctrl.UserSignUpFlow{
		RequestBody: *body,
		Context:     ctx,
	}

	result := flow.Run()

	renderer.JSON(w, http.StatusOK, result)
}
