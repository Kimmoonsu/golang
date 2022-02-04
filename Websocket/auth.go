package main

import (
	"net/http"
	"strings"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/urfave/negroni"
)

func LoginRequired(ignore ...string) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		// ignore URL이면 다음 핸들러 실행
		for _, s := range ignore {
			if strings.HasPrefix(r.URL.Path, s) {
				next(w, r)
				return
			}
		}
		// CurrentUser 정보를 가져옴
		u := GetCurrentUser(r)

		// CurrentUser 정보가 유효하면 만료 시간을 갱신하고 다음 핸들러 실행
		if u != nil && u.Valid() {
			SetCurrentUser(r, u)
			next(w, r)
			return
		}

		// CurrentUser 정보가 유효하지 않으면 CurrentUser를 nil로 세팅
		SetCurrentUser(r, nil)

		// 로그인 후 이동할 URL을 세션에 저장(r)
		sessions.GetSession(r).Set("https://www.naver.com", r.URL.RequestURI())

		// 로그인 페이지로 리다이렉트
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
