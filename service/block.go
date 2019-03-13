package main

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
)

func block(writer http.ResponseWriter, request *http.Request) {
	bytes, _ := ioutil.ReadAll(request.Body)

	var input Request
	if err := json.Unmarshal(bytes, &input); err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ip := net.ParseIP(input.IP)
	response, err := db.Country(ip)
	if err != nil {
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

	if responseBytes, err := json.Marshal(resp); err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(responseBytes)
	}
}
