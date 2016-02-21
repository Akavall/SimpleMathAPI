package main

import (
	"net/http"
	"text/template"
	"log"
	"fmt"
	"strconv"

	"github.com/Akavall/SimpleMathAPI/utilities"
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


func is_prime_wrapper(w http.ResponseWriter, r *http.Request) {
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
		result := math_tools.IsPrime(num)
		log.Printf("Writing result: %t\n", result)
		fmt.Fprint(w, strconv.FormatBool(result))
	}
}


func main() {

	http.HandleFunc("/simple_math", handler)
        http.HandleFunc("/is_prime", is_prime_wrapper)

	socket_address := "0.0.0.0:8088"
	log.Println(socket_address)
	
	error := http.ListenAndServe(socket_address, nil)
	if error != nil {
		log.Fatalln(error)
	}
}
