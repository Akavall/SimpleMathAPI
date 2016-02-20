package main

import (
	"net/http"
	"text/template"
	"log"
	"fmt"
	"strconv"
	"math"
)

func handler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func is_prime_inner(x int) bool {
	if x < 2 { return false }
	if x == 2 { return true }
	if x % 2 == 0 { 
		return false
	} else {
		limit := int(math.Sqrt(float64(x))) + 2
		for i := 3; i < limit; i += 2 {
			if x % i == 0 {
				return false
			}
		}
	}
	return true 
}

func is_prime(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	log.Println(r.Form)

	num_str := r.Form["num"][0]

	log.Println(num_str)
	num, err := strconv.Atoi(num_str)
	if err != nil {
		fmt.Fprint(w, "Please provide a valid integer string")
	} else {
		result := is_prime_inner(num)
		log.Printf("Writing result: %t\n", result)
		fmt.Fprint(w, strconv.FormatBool(result))
	}
}


func main() {

	http.HandleFunc("/simple_math", handler)
        http.HandleFunc("/is_prime", is_prime)

	socket_address := "0.0.0.0:8088"
	log.Println(socket_address)
	
	error := http.ListenAndServe(socket_address, nil)
	if error != nil {
		log.Fatalln(error)
	}
}
