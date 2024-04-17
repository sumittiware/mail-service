package ctrl

import (
	"context"
	"log"
	"mail-service/config"
	data "mail-service/models"

	"github.com/lengzuo/supa/dto"
)

type UserSignInRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignInResponseBody struct {
	ID           string    `json:"id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    int       `json:"expires_at"`
	User         data.User `json:"user"`
	Error        error     `json:"error"`
}

type UserSignInFlow struct {
	RequestBody UserSignInRequestBody
	Context     context.Context
}

func (f *UserSignInFlow) Run() UserSignInResponseBody {
	// validate the request body
	if err := f.validate(); err != nil {
		return UserSignInResponseBody{
			Error: err,
		}
	}

	return f.do()

}

func (f *UserSignInFlow) validate() error {
	// TODO : validate the request body
	return nil
}

func (f *UserSignInFlow) do() UserSignInResponseBody {
	response, err := config.ApplicationConfig.DB.Auth.SignInWithPassword(f.Context, dto.SignInRequest{
		Email:    f.RequestBody.Email,
		Password: f.RequestBody.Password,
	})

	if err != nil {
		log.Fatalf("error signing up user: %v", err)
		return UserSignInResponseBody{
			Error: err,
		}
	}

	user, err := data.GetOne(response.User.ID)

	if err != nil {
		log.Fatalf("error getting user: %v", err)
		return UserSignInResponseBody{
			Error: err,
		}

	}

	return UserSignInResponseBody{
		ID:           response.User.ID,
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		ExpiresAt:    int(response.ExpiresAt),
		User:         *user,
		Error:        nil,
	}
}
