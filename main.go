package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	. "s3uploader/config"
	. "s3uploader/upload"
)

var config = Config{}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	//get file name from url requst
	vars := mux.Vars(r)
	fileName := vars["filename"]

	fmt.Println(r.Header, r.URL.Path, r.Form)

	hburl, err := TestUploadData(&config, r.Body, fileName, config.Expire)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Upload to s3 failed")
		return
	}
	baseResponse(w, http.StatusOK, hburl)
}

func baseResponse(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	w.Write([]byte(msg))
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	baseResponse(w, code, msg)
}

func init() {
	var configPath string
	flag.StringVar(&configPath, "c", "", "usage -c config")
	flag.Parse()
	config.GetConf(configPath)
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{filename}", uploadHandler).Methods("PUT")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	if err := http.ListenAndServe(config.Listen, loggedRouter); err != nil {
		log.Fatal(err)
	}
}
