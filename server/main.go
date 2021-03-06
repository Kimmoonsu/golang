package main

import (
	"fmt"
)

type User struct {
	Id        string
	AddressId string
}

func main() {
	s := NewServer()

	s.HandleFunc("GET", "/", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "welcome!")
	})

	s.HandleFunc("GET", "/about", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "about!")
	})

	s.HandleFunc("GET", "/users/:id", func(c *Context) {
		u := User{Id: c.Params["user_id"].(string)}
		c.RenderXml(u)
	})

	s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		u := User{
			Id:        c.Params["user_id"].(string),
			AddressId: c.Params["address_id"].(string),
		}
		c.RenderJson(u)
	})
	// s.HandleFunc("GET", "/users/:id", func(c *Context) {
	// 	if c.Params["id"] == "0" {
	// 		panic("id is zero")
	// 	}

	// 	fmt.Fprintf(c.ResponseWriter, "retrieve user %v\n", c.Params["id"])
	// })

	// s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
	// 	fmt.Fprintf(c.ResponseWriter, "retrieve user %v's address %v\n", c.Params["user_id"], c.Params["address_id"])
	// })

	// POST
	s.HandleFunc("POST", "/users", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})

	s.HandleFunc("POST", "/users/:user_id/addresses", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})

	s.Run(":5000")
}
