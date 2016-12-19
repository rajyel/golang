// hello_world_webserver.go
package main

import (
	"fmt"
	"net/http"
	//"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	
	fmt.Fprintf(w, "<h1>Hello, %s </h1>", req.URL.Path[1:])
	
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe("localhost:8080", nil)

	//err := http.ListenAndServe("localhost:8080", nil)
	//if err != nil {
		//log.Fatal("ListenAndServe: ", err.Error())
	//}
	// http.ListenAndServe(":8080", http.HandlerFunc(HelloServer))
}
