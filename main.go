package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	serviceToken := req.Header.Get("serviceToken")
	if serviceToken != "secret" {
		http.Error(w, "Ah ah ah. You didn't say the magic word. Ah ah ah. Ah ah ah.", 401)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Println(string(body))
	fmt.Fprintf(w, "Hello Dennis, welcome back!\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("server is listening on 8080")
	http.ListenAndServe(":8080", nil)
}
