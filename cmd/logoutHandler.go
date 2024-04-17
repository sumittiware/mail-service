package cmd

import (
	"net/http"

	"github.com/unrolled/render"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	rendered := render.New()
	rendered.JSON(w, http.StatusOK, map[string]string{"message": "User logged out successfully"})
}
