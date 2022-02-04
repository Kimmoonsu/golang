package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
)

type Room struct {
	ID    int64 `gorm:"primaryKey"`
	Title string
}

func (r *Room) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{&r.Title: "title"}
}

func createRoom(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// binding 패지지 room 생성 요청 정보를 Room 타입 값으로 변환
	fmt.Println("???????????????????")
	r := new(Room)
	errs := binding.Bind(req, r)
	if errs.Handle(w) {
		fmt.Println("errorororororororo")
		return
	}
	fmt.Println("id : ", r.ID)
	fmt.Println("title : ", r.Title)

	db, err := GetDB()
	if err != nil {
		panic(err)
	}

	if err := db.Create(r); err != nil {
		renderer.JSON(w, http.StatusInternalServerError, err)
		return
	}

	// 처리 결과 반환
	renderer.JSON(w, http.StatusCreated, r)
}

func retrieveRooms(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Println("retrieveRooms")
	// renderer.JSON(w, http.StatusOK, rooms)
}
