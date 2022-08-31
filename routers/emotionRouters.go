package routes

import (
	"moodly/controllers"

	"github.com/gorilla/mux"
)

var EmotionRouters = func(router *mux.Router) {
	router.HandleFunc("/emotions/{category}", controllers.GetRandomFeelingByCategory).Methods("GET")

}
