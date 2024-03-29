/*
 * As the rest of the handler files were setup with handling data and not users/accounts
 * I think it'll be more organized if we keep user handlers here
 */

package handlers

import (
	"log"
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
	user.Hash = hash
	user.Admin = false

	result := datamgr.DB.Create(&user)
	if result.Error != nil {
		log.Println("Error when creating user entry in database")
		WriteResponse(w, r, 400) // Error return code
	}

	log.Println("User created: ", user.Name)
	w.WriteHeader(http.StatusCreated)
}

/*
 * Expects to receive a username and password
 * This should receive the password and check it against the hash stored in the database
 */

func Login(w http.ResponseWriter, r *http.Request) {
	// Get login info off of the request
	temp := struct {
		Username string
		Password string
	}{}

	ReadRequest(w, r, &temp)

	// Check username and password in the database, if not correct return unauthorized
	var user datamgr.User
	datamgr.DB.Find(&user, "name = ?", temp.Username)
	err := bcrypt.CompareHashAndPassword(user.Hash, []byte(temp.Password))
	correct := (err == nil)

	if !correct {
		WriteResponse(w, r, http.StatusUnauthorized)
	}

	// Create cookie and push it to active cookie data structure
	// This should utilize http/cookiejar ultimatley I just don't know how it works yet
	bytes, _ := GenerateRandomBytes(32)
	cookie := http.Cookie{
		Name:   "rater-gator user cookie", // generic cookie name for users
		Value:  string(bytes),
		MaxAge: 0, // None specified
	}

	datamgr.CookieJar = append(datamgr.CookieJar, cookie)
	datamgr.CookieJarNames = append(datamgr.CookieJarNames, temp.Username)

	// Now with cookie made and added to the jar, just need to return the cookie in the response
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

/*
 * Should take in the cookie of a user and delete it from the active cookie jar. This way if they try to
 * delete/create anything when the rest of the handler functions look for the cookie it does not find it
 */

func Logout(w http.ResponseWriter, r *http.Request) {
	// Need to remove cookie from the cookiejar and then return
	// If cookie is not found in the cookie jar need to return false
	cookie, err := r.Cookie("rater-gator user cookie")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to find cookie in request body")
		return
	}

	// Get index of cookie to delete
	var index int
	for i, c := range datamgr.CookieJar {
		if c.Value == cookie.Value {
			index = i
		}
	}

	// Log request
	log.Println("Logout request received")

	// Removes the cookie from the cookie jar
	datamgr.CookieJar = append(datamgr.CookieJar[:index], datamgr.CookieJar[index+1:]...)
	datamgr.CookieJarNames = append(datamgr.CookieJarNames[:index], datamgr.CookieJarNames[index+1:]...)
	w.WriteHeader(http.StatusOK)
}

/*
 * GetUserStats:
 *		Takes in a user as input in a json body and returns various statistics about their post history/scores
 *	Sample request body:
 *	{
 *		User: string
 *  }
 *
 * Returns a JSON body of the following form:
 * {
 *		Posts: uint
 *		TotalScore: uint
 * }
 */

func GetUserStats(w http.ResponseWriter, r *http.Request) {
	// Decode request
	request := struct {
		User string
	}{}

	ReadRequest(w, r, &request)

	// Get list of reviews from db
	var reviews []datamgr.Review
	datamgr.DB.Find(&reviews, "Author = ?", request.User)

	// Get user stats by perfoming some form of analysis on the posts they are related to
	stats := struct {
		Posts      int
		TotalScore int
	}{}

	stats.Posts = len(reviews)
	stats.TotalScore = 0

	for _, review := range reviews {
		stats.TotalScore += int(review.Rating)
	}

	log.Println("Request for statistics about", request.User, "received.")
	WriteResponse(w, stats, 200)
}
