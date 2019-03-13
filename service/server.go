package main

import (
	"net/http"

	"github.com/oschwald/geoip2-golang"
)

var db *geoip2.Reader

func main() {
	db, _ = geoip2.Open("GeoLite2-Country.mmdb")
	defer db.Close()
	http.ListenAndServe(":8080", Router())
}
