/*
 * As the rest of the handler files were setup with handling data and not users/accounts
 * I think it'll be more organized if we keep user handlers here
 */

package handlers

import (
	"fmt"
	"net/http"
	"swe-project/backend/datamgr"

	"golang.org/x/crypto/bcrypt"
)

/*
 * CreateUser: This essentially creates the account within the database,
 * expected information within the request body should include a username,
 * and a password, and admin status will be manually controlled probably
 */

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// We expect a username and password to be sent to us
	temp := struct {
		Username string
		Password string
	}{}

	ReadRequest(w, r, &temp)

	// Generate the hash to be stored in the db with the user info
	// Uses constant cost for now, probably will change this in the future
	hash, _ := bcrypt.GenerateFromPassword([]byte(temp.Password), 14)

	var user datamgr.User
	user.Name = temp.Username
	user.Password = temp.Password
	user.Hash = hash
	user.Admin = false

	result := datamgr.DB.Create(&user)
	if result.Error != nil {
		fmt.Print("Error when creating user entry in database")
		WriteResponse(w, r, 400) // Error return code
	}

	fmt.Print("User created: ", user.Name)
	WriteResponse(w, r, http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// This should cook up the cookie and distribute it
	// Compares password sent to backend with the hash, if it fits, distribtute cookie
	// Establishes session
	// Otherwise return 401 unauthorized
}
