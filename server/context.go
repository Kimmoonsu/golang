package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Context struct {
	Params map[string]interface{}

	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

type HandlerFunc func(*Context)

func (c *Context) RenderJson(v interface{}) {
	// HTTP Status 를 StatusOK로 지정
	c.ResponseWriter.WriteHeader(http.StatusOK)
	// Content-Type을 application/json으로 지정
	c.ResponseWriter.Header().Set("Content-Type", "application/json; charset-utf-8")

	// v 값을 json으로 출력
	if err := json.NewEncoder(c.ResponseWriter).Encode(v); err != nil {
		// 에러 발생 시 RenderErr 메서드 호출
		c.RenderErr(http.StatusInternalServerError, err)
	}
}

func (c *Context) RenderXml(v interface{}) {
	// HTTP Status 를 StatusOK로 지정
	c.ResponseWriter.WriteHeader(http.StatusOK)
	// Content-Type을 application/json으로 지정
	c.ResponseWriter.Header().Set("Content-Type", "application/xml; charset-utf-8")

	// v 값을 xml로 출력
	if err := xml.NewEncoder(c.ResponseWriter).Encode(v); err != nil {
		// 에러 발생 시 RenderErr 메서드 호출
		c.RenderErr(http.StatusInternalServerError, err)
	}
}

func (c *Context) RenderErr(code int, err error) {
	if err != nil {
		if code > 0 {
			// 정상적인 code를 전달하면 HTTP Status를 해당 code로 지정
			http.Error(c.ResponseWriter, http.StatusText(code), code)
		} else {
			// 정상적인 code가 아니면 HTTP Status를 StatusInternalServerError로 지정
			defaultErr := http.StatusInternalServerError
			http.Error(c.ResponseWriter, http.StatusText(defaultErr), defaultErr)
		}
	}
}
