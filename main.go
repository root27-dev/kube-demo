package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

		w.Write([]byte("Hello Kubernetes"))
	})

	fmt.Println("Server started at port 8080")

	http.ListenAndServe(":8080", nil)

}
