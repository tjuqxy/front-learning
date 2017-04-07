package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"

	"github.com/tjuqxy/http-server/api"
)

func main() {
	m := martini.Classic()

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET"},
	}))

	m.Use(martini.Static("/Users/qinxingyao/qxy/page"))

	m.Post("/api/sudo/calculate", api.SudoCalculate)

	m.RunOnAddr("127.0.0.1:8080")
}