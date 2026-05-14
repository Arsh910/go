package main

import "net/http"

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}

func handleError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "something wrong")
}
