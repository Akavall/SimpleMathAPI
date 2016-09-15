package main

import (
	"net/http"
	"text/template"
	"log"
	"fmt"
	"strconv"
	"os"

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

	num_str := r.Form["num"][0]

	log.Printf("Receiving: %s num_str\n", num_str)
	num, err := strconv.Atoi(num_str)
	if err != nil {
		fmt.Fprint(w, "Please provide a valid integer string")
	} else {
		div, result := math_tools.IsPrime(num)
		log.Printf("Writing result: %d, %t\n", div, result)
		var w_string string
		if result {
			w_string = fmt.Sprintf("%d is prime\n", num)
		} else if num < 0 {
			w_string = fmt.Sprintf("%d is not prime, negative numbers cannot be prime\n", num) 
		} else if num == 0 || num == 1 {
			w_string = fmt.Sprintf("%d is not prime by convention\n", num)
		} else {
			w_string = fmt.Sprintf("%d is not prime, smallest divisor: %d\n", num, div)
		}

		fmt.Fprint(w, w_string)
	}
}


func main() {

	// We need need to create a file manually
	// and change the permissions with: 
	// sudo chmod 666 logfile.txt
	f, err := os.OpenFile("/var/log/SimpleMathAPI/logfile.txt", os.O_RDWR|os.O_APPEND, 0660)

	if err != nil {
		fmt.Println("Could not open logfile.txt")
	}

	log.SetOutput(f)

	http.HandleFunc("/simple_math", handler)
        http.HandleFunc("/is_prime", is_prime_wrapper)

	socket_address := "0.0.0.0:8088"
	log.Println(socket_address)
	
	error := http.ListenAndServe(socket_address, nil)
	if error != nil {
		log.Fatalln(error)
	}
}
