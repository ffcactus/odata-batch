package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var (
	client = http.Client{}
)

// OperationValue is the request DTO for the action supported by calculator service.
type OperationValue struct {
	Value float64 `json:"Value"`
}

// ResultValue is the response DTO for the action supported by calculator service.
type ResultValue struct {
	Result float64 `json:"Result"`
}

func add(value float64) {
	op(value, "http://localhost:3000/calculator/Actions/Calculator.Add")
}

func sub(value float64) {
	op(value, "http://localhost:3000/calculator/Actions/Calculator.Sub")
}

func mul(value float64) {
	op(value, "http://localhost:3000/calculator/Actions/Calculator.Mul")
}

func div(value float64) {
	op(value, "http://localhost:3000/calculator/Actions/Calculator.Div")
}

func reset() {
	op(0, "http://localhost:3000/calculator/Actions/Calculator.Reset")
}

func op(value float64, url string) {
	dto := &OperationValue{
		Value: value,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(dto)

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		log.Printf("NewRequest() failed, err = %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Get response failed, err = %v\n", err)
	}

	defer resp.Body.Close()
	result := &ResultValue{}

	if resp.StatusCode != http.StatusOK {
		log.Printf("HTTP status code = %d\n", resp.StatusCode)
	} else {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			log.Printf("Parse result failed, err = %v\n", err)
		}
		log.Printf("Result = %f\n", result.Result)
	}
}

func main() {
	reset()
	add(1)
	add(2)
}
