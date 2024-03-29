package routes

import (
	"log"
	"net/http"

	"github.com/fincor-Blockchain/daemon-extract/utils/helper"

	"github.com/gorilla/mux"
	"github.com/fincor-Blockchain/daemon-extract/controllers"
)

//StartRoutes listening
func StartRoutes() {
	helper.ServerGenerateKey()
	router := mux.NewRouter().StrictSlash(true)

	dev := router.PathPrefix("/api/dev").Subrouter()

	dev.HandleFunc("/testing", controllers.Testing).Methods("POST")
	dev.HandleFunc("/generateCSVfromJSON", controllers.GenerateCSV).Methods("POST")
	dev.HandleFunc("/cleanData", controllers.CleanData).Methods("GET")
	dev.HandleFunc("/normalizeCSVData", controllers.NormalizeData).Methods("GET")
	dev.HandleFunc("/splitData", controllers.SplitAndShuffle).Methods("POST")
	//dev.HandleFunc("/train", controllers.GenerateCSV).Methods("GET")
	//dev.HandleFunc("/test", controllers.GenerateCSV).Methods("GET")
	//dev.HandleFunc("/validate", controllers.GenerateCSV).Methods("GET")
	dev.HandleFunc("/predict", controllers.Predict).Methods("POST")
	dev.HandleFunc("/trainComplete", controllers.Train).Methods("GET")
	dev.HandleFunc("/getKey", controllers.GettingPublicKey).Methods("GET")
	//dev.HandleFunc("/gettingData", controllers.GettingData).Methods("POST")

	log.Fatal(http.ListenAndServe(":4400", router))
}
