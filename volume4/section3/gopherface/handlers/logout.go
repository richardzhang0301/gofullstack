package handlers

import (
	"net/http"

	"github.com/richardzhang0301/gofullstack/volume4/section3/gopherface/common/authenticate"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}
