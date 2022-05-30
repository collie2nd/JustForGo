package main

import (
	"encoding/json"
	"fmt"
)

type data struct {
	Id   string
	Name string
}

func main() {
	id := [5]string{"1", "2", "3", "4", "5"}
	name := [5]string{"A", "B", "C", "D", "E"}
	var parsedData []data

	for counter := range id {
		parsedData = append(parsedData, data{Name: name[counter], Id: id[counter]})
	}

	bytes, _ := json.Marshal(parsedData)
	fmt.Print(string(bytes))
}
