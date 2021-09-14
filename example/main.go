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

	// err := json.NewDecoder(r.Body).Decode(user) // NewDecoder는 io.Reader를 인자로 받고 Body는 io.Reader를 포함하고 있다.
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest) // hear로 오류가 있다는 것을 알려줌
	// 	fmt.Fprint(w, "Bad Request: ", err)  // body에 에러를 알려줌
	// 	return
	// }

	data, _ := json.Marshal(user) // 인터페이스를 받아서 json형태로 바꿔주는 메소드(byte와 err를 리턴함)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprint(w, string(data))
}

func (u *insertUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)

	err := json.NewDecoder(r.Body).Decode(user) // NewDecoder는 io.Reader를 인자로 받고 Body는 io.Reader를 포함하고 있다.
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // hear로 오류가 있다는 것을 알려줌
		fmt.Fprint(w, "Bad Request: ", err)  // body에 에러를 알려줌
		return
	}
	insert(*user)
	// data, _ := json.Marshal(user) // 인터페이스를 받아서 json형태로 바꿔주는 메소드(byte와 err를 리턴함)
	// w.Header().Add("content-type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// fmt.Println("user : ", user.ID)
	// fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	// URL에서 argument를 뽑아냄
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "Hello %s!", name) // 그 값을 출력
}

func main() {
	connection()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/user", &userHandler{})
	mux.Handle("/insertUser", &insertUserHandler{})
	http.ListenAndServe(":3000", mux)
}
