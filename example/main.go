package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID  string
	PWD string
}

type userHandler struct{}
type insertUserHandler struct{}

func (u *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := selectAll()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprint(w, string(data))
}

func (u *insertUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	insert(*user)
}

func main() {
	connection()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.Handle("/user", &userHandler{})
	mux.Handle("/insertUser", &insertUserHandler{})
	http.ListenAndServe(":3000", mux)
}
