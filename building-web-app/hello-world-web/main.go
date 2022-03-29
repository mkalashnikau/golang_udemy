package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes: %d", n)) //fmt.Println("Number of bytes:" + n)) wouldn't work because n has non-string type
	})

	// Start web server
	// _ means that if there is no error it is ok
	_ = http.ListenAndServe(":8080", nil)
}
