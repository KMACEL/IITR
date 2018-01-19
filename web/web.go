package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//https://astaxie.gitbooks.io/build-web-application-with-golang/en/

/*
████████╗███████╗███████╗████████╗
╚══██╔══╝██╔════╝██╔════╝╚══██╔══╝
   ██║   █████╗  ███████╗   ██║
   ██║   ██╔══╝  ╚════██║   ██║
   ██║   ███████╗███████║   ██║
   ╚═╝   ╚══════╝╚══════╝   ╚═╝
*/

func main() {

	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)
	http.HandleFunc("/from", formGet)

	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func deviceStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, err := template.ParseFiles("/home/acel/go/src/github.com/KMACEL/IITR/web/devicereport.gtpl")
		fmt.Println("err : ", err)
		t.Execute(w, nil)
		fmt.Println("Hello Mert")
	} else {
		r.ParseForm()
		fmt.Println("Hello Mert")
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, err := template.ParseFiles("/home/acel/go/src/github.com/KMACEL/IITR/web/login.gtpl")
		fmt.Println("err : ", err)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		if len(r.Form["username"][0]) > 0 || len(r.Form["password"][0]) > 0 {
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])

			slice := []string{"apple", "pear", "bananaa"}
			for _, v := range slice {
				if v == r.Form.Get("fruit") {
					fmt.Println("Yes ", v)
				}
			}

			slice2 := []string{"1", "2"}

			for _, v2 := range slice2 {
				if v2 == r.Form.Get("gender") {
					fmt.Println("Yes ", v2)
				}
			}

			slice3 := []string{"football", "basketball", "tennis"}
			a, _ := r.Form["interest"], slice3
			fmt.Println("haha :", a)

			t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
			fmt.Printf("Go launched at %s\n", t.Local())

		} else {
			fmt.Fprintf(w, "Username & Password Null")

		}
	}
}

func formGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("methodxx:", r.Method) //get request method
	if r.Method == "GET" {
		fmt.Println("All :", r.Form.Get("age"))
		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			// error occurs when convert to number, it may not a number
			fmt.Println("Error :", err)
		}

		fmt.Println(getint)
	}

	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		//return false
	}

	// email
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}

	// english character
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
		///	return false
	}

}

/*func helloWeb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))

		if k=="haha"{
			fmt.Fprintf(w, "Zuhahah")
		}
	}

	fmt.Fprintf(w, "Acel")

	if r.URL.Path == "/abc" {
		fmt.Fprintf(w, "Mert")
	}
}*/
