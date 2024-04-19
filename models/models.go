package data

import (
	"time"

	supabase "github.com/lengzuo/supa"
)

const dbTimeout = time.Second * 3

var db *supabase.Client

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *supabase.Client) Models {
	db = dbPool

	return Models{
		User: User{},
		Plan: Plan{},
	}
}

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type Models struct {
	User User
	Plan Plan
}

const (
	userTable     = "users"
	planTable     = "plans"
	userPlanTable = "user_plans"
)
