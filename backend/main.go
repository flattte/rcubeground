package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type sessionResolutionPoints struct {
	store *sessions.CookieStore
	users map[string]string
}

func init_router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	router.HandleFunc("/login_req", srp.loginHandlerCA).Methods("POST")
	router.HandleFunc("/logout", srp.logoutHandlerCA).Methods("GET")
	router.HandleFunc("/healthcheck", srp.healtcheckCA).Methods("GET")
	router.HandleFunc("/ws", serveWebsocket)

	spa := spaHandler{
		staticPath: "../Rubik-Cube/build",
		indexPath:  "/index.html",
	}
	router.PathPrefix("/").Handler(spa)
	return router
}

var srp = &sessionResolutionPoints{
	store: sessions.NewCookieStore([]byte("my_secret_key")),
	users: map[string]string{"user1@lmao.com": "password", "user2@lmao.com": "password2"},
}

func main() {
	router := init_router()

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	Debugf("Server running on %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
