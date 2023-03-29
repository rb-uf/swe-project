//go:build alt
// +build alt

package handlers

import (
	"net/http"
)

/*
 * Returns whether or not the operation should be performed
 * third operand is the level of permissions needed, so if the author can modify or not
 */

func CheckCookieAndPermissions(w http.ResponseWriter, r *http.Request, author_perm bool, author string) bool {
	return true
}
