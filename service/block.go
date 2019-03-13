package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

func block(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var input Request

	err := decoder.Decode(&input)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ip := net.ParseIP(input.IP)
	response, err := db.Country(ip)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	var allow bool
	for _, country := range input.Countries {
		if country == response.Country.IsoCode {
			allow = true
		}
	}

	var resp Body
	if allow {
		resp = Body{Action: "ALLOW"}
	} else {
		resp = Body{Action: "BLOCK"}
	}

	if err := json.NewEncoder(writer).Encode(resp); err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		writer.Header().Set("Content-Type", "application/json")
	}
}
