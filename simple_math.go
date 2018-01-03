package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"./utilities"
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

	_, ok := r.Form["num"]
	if ok == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - request does not specify: num"))
		return
	}

	num_str := r.Form["num"][0]

	log.Printf("Receiving: %s num_str\n", num_str)
	num, err := strconv.Atoi(num_str)
	if err != nil {
		error_type, ok := err.(*strconv.NumError)

		if !ok {
			log.Printf("Failed to convert error to NumError")
		}

		if error_type.Err == strconv.ErrSyntax {
			fmt.Fprint(w, "Please provide a valid integer string")
		} else if error_type.Err == strconv.ErrRange {
			fmt.Fprint(w, "Please provide a smaller number")
		} else {
			fmt.Fprint(w, "Hmm..strange error, you discoverd something we did not cover")
		}
	} else {
		div, result, err := math_tools.IsPrime(num)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("500 = %s", err)))
			return
		}

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
