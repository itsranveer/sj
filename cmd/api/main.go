package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itsranveer/sj/routes"
	"github.com/itsranveer/sj/worker"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting the Math API on 8080 ...")

	math := worker.NewMath()
	allRoutes := routes.NewRoutes(math)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/min", allRoutes.Min).Methods(http.MethodGet)
	router.HandleFunc("/max", allRoutes.Max).Methods(http.MethodGet)
	router.HandleFunc("/avg", allRoutes.Avg).Methods(http.MethodGet)
	router.HandleFunc("/median", allRoutes.Median).Methods(http.MethodGet)
	router.HandleFunc("/percentile", allRoutes.Percentile).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
