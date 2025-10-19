package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func index(writer http.ResponseWriter, request *http.Request) {
	var newCookieValue string
	var e error
	if cookie, err := request.Cookie("_cookie"); err != nil {
		newCookieValue = "0"
		e = err
	} else if numValue, err := strconv.Atoi(cookie.Value); err != nil {
		newCookieValue = "0"
		if numErr, ok := err.(*strconv.NumError); ok {
			e = numErr.Err
		} else {
			e = err
		}
	} else {
		numValue++
		newCookieValue = fmt.Sprintf("%d", numValue)
	}
	if e != nil {
		fmt.Printf("Error: %s", e.Error())
	}

	newCookie := http.Cookie{
		Name:     "_cookie",
		Value:    newCookieValue,
		HttpOnly: true,
		Secure:   true}
	http.SetCookie(writer, &newCookie)
	fmt.Fprintf(writer, "Protocol: %s\n", request.Proto)
	fmt.Fprintf(writer, "Hello, World: %s!\n", request.URL.Path[1:])
	fmt.Fprintf(writer, "Cookie Value: %s!\n", newCookieValue)

}

func main() {
	const certFile = "localhost+1.pem"
	const keyFile = "localhost+1-key.pem"
	const addr = ":8443"

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	fmt.Printf("Starting HTTPS server on %s...\n", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	err := server.ListenAndServeTLS(certFile, keyFile)

	if err != nil {
		log.Fatal("ListenAndServeTLS error:", err)
	}
}
