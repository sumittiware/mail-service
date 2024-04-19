package ctrl

import (
	"context"
	"log"
	"mail-service/config"
	data "mail-service/models"

	"github.com/lengzuo/supa/dto"
)

type UserSignUpRequestBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserSignUpResponseBody struct {
	ID           string `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int    `json:"expires_at"`
	Error        error  `json:"error"`
}

type UserSignUpFlow struct {
	RequestBody UserSignUpRequestBody
	Context     context.Context
}

func (f *UserSignUpFlow) Run() UserSignUpResponseBody {
	// validate the request body
	if err := f.validate(); err != nil {
		return UserSignUpResponseBody{
			Error: err,
		}
	}

	return f.do()

}

func (f *UserSignUpFlow) validate() error {
	// TODO : validate the request body
	return nil
}

func (f *UserSignUpFlow) do() UserSignUpResponseBody {
	response, err := config.ApplicationConfig.DB.Auth.SignUp(f.Context, dto.SignUpRequest{
		Email:    f.RequestBody.Email,
		Password: f.RequestBody.Password,
	})

	if err != nil {
		log.Fatalf("error signing up user: %v", err)
		return UserSignUpResponseBody{
			Error: err,
		}
	}

	// Create the user models based on this
	user := data.User{
		ID:        response.User.ID,
		Email:     f.RequestBody.Email,
		FirstName: f.RequestBody.FirstName,
		LastName:  f.RequestBody.LastName,
		IsAdmin:   false,
		Active:    true,
		CreatedAt: response.User.CreatedAt,
		UpdatedAt: response.User.UpdatedAt,
		Plan:      nil,
	}
	// create the user in our DB
	if err := user.Save(); err != nil {
		log.Fatalf("error saving user: %v", err)
		return UserSignUpResponseBody{
			Error: err,
		}
	}

	return UserSignUpResponseBody{
		ID:           response.User.ID,
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		ExpiresAt:    int(response.ExpiresAt),
		Error:        nil,
	}
}
