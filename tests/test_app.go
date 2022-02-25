package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const server = "http://0.0.0.0:8000/"

type Car struct {
	Make  string
	Model string
	Trim  string
	Color string
	Year  int
}

func getIndex() {
	resp, err := http.Get(server)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var bodyContent string = string(body)

	fmt.Println(bodyContent)
}

func getCars() {
	resp, err := http.Get(server + "cars")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	allCars := []Car{}
	_ = json.Unmarshal(body, &allCars)
	fmt.Println(allCars)
}

func getExactCarByMake() {
	resp, err := http.Get(server + "cars/4Runner")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fourRunnerCar := Car{}
	_ = json.Unmarshal(body, &fourRunnerCar)
	fmt.Println(fourRunnerCar)
}

func main() {
	getIndex()
	getCars()
	getExactCarByMake()
}
