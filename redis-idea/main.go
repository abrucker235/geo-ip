package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

//Country structure
type Country struct {
	Name  string   `json:"name"`
	ID    string   `json:"id"`
	CIDRS []string `json:"cidrs"`
}

type Block struct {
	id           string
	registeredID string
	cidrs        []string
}

func main() {
	countries := createCountries("GeoLite2-Country-Locations-en.csv")
	blocks := createBlocks("GeoLite2-Country-Blocks-IPv4.csv")

	//TODO: Add CIDRS from blocks to countries based on geoname_id/registered_geoname_id
}

func createCountries(path string) map[string]Country {
	csvFile, _ := os.Open(path)
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	countries := make(map[string]Country)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		countries[line[5]] = Country{Name: line[5], ID: line[0]}
	}
	return countries
}

func createBlocks(path string) map[string]Block {
	csvFile, _ := os.Open(path)
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	blocks := make(map[string]Block)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if block, ok := blocks[line[2]]; ok {
			block.cidrs = append(block.cidrs, line[0])
		} else {
			blocks[line[2]] = Block{id: line[1], registeredID: line[2], cidrs: []string{line[0]}}
		}

	}
	return blocks
}
