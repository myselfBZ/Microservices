
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello world")
    })
    log.Fatal(http.ListenAndServe(":8080", mux))
}


