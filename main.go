package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"

	"github.com/tjuqxy/web/api"
)

func main() {
	m := martini.Classic()

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET"},
	}))

	m.Use(martini.Static("./public"))

	m.Post("/api/sudo/calculate", api.SudoCalculate)

	m.RunOnAddr("0.0.0.0:80")
}
