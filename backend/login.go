package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type postLoginForm struct {
	Username string
	Password string
}

func (srp *sessionResolutionPoints) loginHandlerCA(w http.ResponseWriter,
	r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Login method not supported", http.StatusBadRequest)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Login method not supported", http.StatusBadRequest)
		return
	}
	Debugf("req: %s\n", string(b))
	var plf postLoginForm
	err = json.Unmarshal(b, &plf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Debugf("%s\n", b)
	username := plf.Username
	password := plf.Password
	Debugf("\tpass %s\n", password)
	Debugf("\tuser %s\n", username)
	storedPassword, ok := srp.users[username]

	if !ok {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	session, _ := srp.store.Get(r, "session.id")
	if storedPassword == password {
		Debugf("[INFO] user validated")
		session.Values["authenticatred"] = "true"
		session.Save(r, w)
	} else {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}
	w.Write([]byte("Login succesfully!"))
}

func (srp *sessionResolutionPoints) logoutHandlerCA(w http.ResponseWriter,
	r *http.Request) {
	session, _ := srp.store.Get(r, "session.id")
	session.Values["authenticated"] = false
	session.Save(r, w)
	w.Write([]byte("Logout Succesfull"))
}

func (srp *sessionResolutionPoints) healtcheckCA(w http.ResponseWriter, r *http.Request) {
	session, _ := srp.store.Get(r, "session.id")
	authenticated := session.Values["authenticated"]
	if authenticated != nil && authenticated != false {
		w.Write([]byte("Welcome!"))
		return
	}
	http.Error(w, "forbidden", http.StatusForbidden)
	return
}
