package data

import (
	"context"
	"log"
	"time"
)

// User is the structure which holds one user from the database.
type User struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Active    int       `json:"user_active"`
	IsAdmin   int       `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAll returns a slice of all users, sorted by last name
func GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := db.DB.From(userTable).Select("*")

	var users []*User

	if err := query.Execute(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetByEmail returns one user by email
func GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var users *[]User

	query := db.DB.From(userTable).Select("*").Eq("email", email)
	if err := query.Execute(ctx, &users); err != nil {
		log.Fatalf("Error getting user by email: %v", err)
		return nil, err
	}
	log.Println("User: ", users)
	return nil, nil
}

// GetOne returns one user by id
func GetOne(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var user *[]User
	query := db.DB.From(userTable).Select("*").Eq("user_id", id)

	if err := query.Execute(ctx, &user); err != nil {
		log.Fatalf("Error getting user by id: %v", err)
		return nil, err
	}
	log.Println("User: ", user)

	// var plan *[]Plan

	// query = db.DB.From(planTable).Select("*").Eq("user_id", user[0].UserID)

	// if err := query.Execute(ctx, &plan); err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (u *User) Save() error {
	ctx := context.Background()
	query := db.DB.From(userTable).Insert(u)

	if err := query.Execute(ctx, nil); err != nil {
		return err
	}

	return nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
// func (u *User) Update() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	stmt := `update users set
// 		email = $1,
// 		first_name = $2,
// 		last_name = $3,
// 		user_active = $4,
// 		updated_at = $5
// 		where id = $6`

// 	_, err := db.ExecContext(ctx, stmt,
// 		u.Email,
// 		u.FirstName,
// 		u.LastName,
// 		u.Active,
// 		time.Now(),
// 		u.ID,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// Delete deletes one user from the database, by User.ID
func (u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := db.DB.From(userTable).Delete().Eq("user_id", u.UserID)

	if err := query.Execute(ctx, nil); err != nil {
		return err
	}

	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
// func (u *User) Insert(user User) (int, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
// 	if err != nil {
// 		return 0, err
// 	}

// 	var newID int
// 	stmt := `insert into users (email, first_name, last_name, password, user_active, created_at, updated_at)
// 		values ($1, $2, $3, $4, $5, $6, $7) returning id`

// 	err = db.QueryRowContext(ctx, stmt,
// 		user.Email,
// 		user.FirstName,
// 		user.LastName,
// 		hashedPassword,
// 		user.Active,
// 		time.Now(),
// 		time.Now(),
// 	).Scan(&newID)

// 	if err != nil {
// 		return 0, err
// 	}

// 	return newID, nil
// }
