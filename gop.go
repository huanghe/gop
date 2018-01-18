package main

import (
	"fmt"
	"net/http"
	"log"
)
/**
	第一个go server
 */
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090",nil)
	if err!=nil {
		log.Fatal("ListenAndServer:",err)
	}
}
func sayHello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)

	fmt.Fprintf(w, "Hello Guest!")
}

