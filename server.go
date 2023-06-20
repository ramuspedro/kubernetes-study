package main

import ( 
		"net/http" 
		"os"
		"fmt"
		"io/ioutil"
		"log"
	)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configmap", ConfigMap)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	// w.Write([]byte("<h1>Hello Full Cycle!!!!!!!!!</h1>"))
	fmt.Fprintf(w, "Hello, I am %s. I'm %s.", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file", err)
	}
	fmt.Fprintf(w, "My family: ", string(data))
}