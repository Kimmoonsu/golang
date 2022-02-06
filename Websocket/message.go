package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

const messageFetchSize = 10

type Message struct {
	ID        int
	RoomId    int
	Content   string
	CreatedAt time.Time
	User      *User
}

func (m *Message) create() error {
	db, _ := GetDB()
	// err != nil{}

	m.CreatedAt = time.Now()

	db.Create(m)
	// err != nil{
	// 	// return err
	// }
	// 처리 결과 반환
	return nil
}

func retrieveMessages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// db, err := GetDB()
	// err != nil {
	// 	panic(err)
	// }
	// 쿼리 매개변수로 전달된 limit 값 확인
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		// 정상적인 limit 값이 전달되지 않으면 limit를 messageFetchSize로 세팅
		limit = messageFetchSize
	}

	fmt.Println("limit : ", limit)

	// var message []Message
	// _id 역순으로 정렬하여 limit 수만큼 message 조회

}
