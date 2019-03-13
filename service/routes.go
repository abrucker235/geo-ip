package main

import (
	"net/http"
)

type route struct {
	name        string
	method      string
	pattern     string
	handlerFunc http.HandlerFunc
}

var routes = []route{
	route{
		"block",
		"POST",
		"/",
		block,
	},
}
