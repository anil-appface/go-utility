package store

import (
	"database/sql"
	"fmt"
	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

type User struct {
	ID        string `json:"id"`
	LoginID   string `json:"loginId"`
	FullName  string `json:"fullName"`
	IsEnabled bool   `json:"isEnabled"`
}

type UserRepo interface {
	UpsertUser(user *User) (*User, error)
	Delete(id string) error
	GetAll() ([]User, error)
}

type UserRepoSQL struct {
	db *sql.DB
}

func NewUserRepoSQL(db *sql.DB) UserRepo {
	return &UserRepoSQL{
		db: db,
	}
}

func (me *UserRepoSQL) UpsertUser(user *User) (*User, error) {

	var stmt *sql.Stmt
	var err error
	isInserting := false
	if user.ID != "" {
		stmt, err = me.db.Prepare("UPDATE user SET loginid = ?, fullname = ?, isenabled = ? WHERE id = ?")
	} else {
		isInserting = true
		uuid, _ := uuid.NewV4()
		user.ID = strings.ReplaceAll(uuid.String(), "-", "")
		stmt, err = me.db.Prepare("INSERT INTO user (id, loginid, fullname, isenabled) VALUES(?,?,?,?)")
	}
	if err != nil {
		fmt.Printf("error while preparing query for inserting/updating user record: %#v", err)
		return nil, err
	}
	fmt.Println(user.ID)
	if !isInserting {
		_, err = stmt.Exec(user.LoginID, user.FullName, user.IsEnabled, user.ID)
	} else {
		_, err = stmt.Exec(user.ID, user.LoginID, user.FullName, user.IsEnabled)
	}

	if err != nil {
		fmt.Printf("error while executing statement for inserting/updating user record: %#v", err)
		return nil, err
	}
	return user, nil
}

func (me *UserRepoSQL) Delete(id string) error {

	stmt, err := me.db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		fmt.Printf("error while preparing query for deleting user record: %#v", err)
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		fmt.Printf("error while executing statement for deleting user record: %#v", err)
		return err
	}
	return nil
}

func (me *UserRepoSQL) GetAll() ([]User, error) {

	allUsers := make([]User, 0)
	rows, err := me.db.Query("SELECT id, loginid, fullname, isenabled FROM user")
	if err != nil {
		fmt.Printf("error while getting all users %#v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, loginID, fullName string
		var isEnabled bool
		err = rows.Scan(&id, &loginID, &fullName, &isEnabled)
		if err != nil {
			// handle this error
			fmt.Printf("error while getting all user %#v", err)
			return nil, err
		}
		allUsers = append(allUsers, User{
			ID:        id,
			LoginID:   loginID,
			FullName:  fullName,
			IsEnabled: isEnabled,
		})
	}
	// get any error encountered during iteration
	err = rows.Err()
	return allUsers, err
}
