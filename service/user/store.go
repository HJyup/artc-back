package user

import (
	"artc-back/types"
	"database/sql"
	"fmt"
	"os/exec"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (store *Store) CreateUser(user types.User) error {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		return err
	}
	_, err = store.db.Exec("INSERT INTO users (id, first_name, last_name, email, password, speciality_id, location) VALUES (?, ?, ?, ?, ?, ?, ?)",
		uuid, user.FirstName, user.SecondName, user.Email, user.Password, user.SpecialityID, user.Location)
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) GetUserById(id string) (*types.User, error) {
	rows, err := store.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user == nil || user.PK == 0 {
		return nil, fmt.Errorf("error getting user by id: %s", id)
	}

	return convertToResponseUser(user), nil
}

func (store *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := store.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user == nil || user.PK == 0 {
		return nil, fmt.Errorf("error getting user by email: %s", email)
	}

	return convertToResponseUser(user), nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.PK,
		&user.ID,
		&user.Avatar,
		&user.FirstName,
		&user.SecondName,
		&user.Email,
		&user.Password,
		&user.SpecialityID,
		&user.Location,
		&user.IsAccepted,
		&user.IsReviewer,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func convertToResponseUser(user *types.User) *types.User {
	return &types.User{
		ID:           user.ID,
		Avatar:       user.Avatar,
		FirstName:    user.FirstName,
		SecondName:   user.SecondName,
		Email:        user.Email,
		Password:     user.Password,
		SpecialityID: user.SpecialityID,
		Location:     user.Location,
		IsAccepted:   user.IsAccepted,
		IsReviewer:   user.IsReviewer,
		CreatedAt:    user.CreatedAt,
	}
}
