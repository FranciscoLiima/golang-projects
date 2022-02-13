package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/create", app.createDeckOfCards)
	
	router.HandlerFunc(http.MethodGet, "/v1/build", app.createDeckOfCardsByCards)

	router.HandlerFunc(http.MethodGet, "/v1/shuffle", app.shuffleDeck)

	router.HandlerFunc(http.MethodGet, "/v1/open", app.openDeckOfCards)

	router.HandlerFunc(http.MethodGet, "/v1/draw", app.drawCards)

	return router
}