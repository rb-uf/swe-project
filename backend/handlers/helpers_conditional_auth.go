//go:build !alt
// +build !alt

package handlers

import (
	"fmt"
	"net/http"
	"swe-project/backend/datamgr"
)

/*
 * Returns whether or not the operation should be performed
 * third operand is the level of permissions needed, so if the author can modify or not
 */

func CheckCookieAndPermissions(w http.ResponseWriter, r *http.Request, author_perm bool, author string, bypass bool) bool {

	c, _ := r.Cookie("rater-gator-cookie")

	// If cookie does not exist
	if c == nil {
		fmt.Println("Error, no cookie found")
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	present, user := VerifyCookie(*c)

	// If cookie was not issued return unauthorized
	if !present {
		fmt.Println("Error, request's cookie not found in cookiejar")
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	// Get user info from db
	var temp datamgr.User
	datamgr.DB.Find(&temp, "Name = ?", user)

	// If user is not an admin do not let them delete the subject
	if !bypass && (!temp.Admin || (author_perm && author != temp.Name)) {
		fmt.Println("Error, requester does not have permission to delete a subject")
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}

func ConfigureCookie(r *http.Request, v string) {
	r.Header.Set("Cookie", "Foo=Bar; ; ")

	r.AddCookie(&http.Cookie{
		Name:  "rater-gator-cookie",
		Value: v,
	})
}
