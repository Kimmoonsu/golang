package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
)

type Room struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func (r *Room) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{&r.Name: "name"}
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
	fmt.Println("name : ", r.Name)

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
	db, err := GetDB()
	if err != nil {
		fmt.Println("getdb err : ", err)
		panic(err)
	}
	var rooms []Room
	// 모든 room 정보 조회
	if err := db.Find(&rooms).Error; err != nil {
		// 오류 발생시 500 에러 리턴
		fmt.Println("select err : ", err)
		renderer.JSON(w, http.StatusInternalServerError, err)
		return
	}
	// room 조회 결과 리턴
	renderer.JSON(w, http.StatusOK, rooms)
}
