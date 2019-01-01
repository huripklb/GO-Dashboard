package auth

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"

	dc "dbconn"

	"github.com/delivery-api/model"
)

// encryptPassword : encrypt password string with md5 (to be replaced with other encrypt method agreed)
func encryptPassword(password string) string {
	data := []byte(password)
	encryptedPass := fmt.Sprintf("%x", md5.Sum(data))

	return encryptedPass
}

// checkCredentials : match basic auth credetials sent with database data
func checkCredentials(username string, password string) bool {
	var users []model.Apiuser

	db, err := dc.PostgreConn()
	if err != nil {
		panic(err)
	}

	db.Where("username = ? AND password = ?", username, password).First(&users).RecordNotFound()

	if len(users) == 0 {
		log.Println("Invalid password or user not found")
		return false
	}

	return true
}

// BasicAuth : Go Mux Router Middleware for basic auth credential checking
// author : Huripto Sugandi
// created date : 5 Dec 2018
func BasicAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		username, password, authOK := r.BasicAuth()
		if authOK == false {
			http.Error(w, "Not authorized", 401)
			return
		}

		if checkCredentials(username, encryptPassword(password)) == false {
			http.Error(w, "Not authorized", 401)
			return
		}

		h.ServeHTTP(w, r)
	})
}
