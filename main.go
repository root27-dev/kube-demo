package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

		index := []byte(`<!DOCTYPE html>
				<html>
				<head>
					<title>My Kubernetes App</title>
				</head>
				<body>
					<p> Request from: ` + r.RemoteAddr + `</p>
					<p> Hostname: ` + r.Host + `</p>

				</body>
				</html>`)

		w.Write(index)
	})

	fmt.Println("Server started at port 8080")

	http.ListenAndServe(":8080", nil)

}
