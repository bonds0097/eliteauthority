package main

import (
	"net/http"
)

// Master Database connection from which others are spawned.
var dbMaster MongoDbConnection

// TODO: Tweak for dev/prod?
func setStaticHandlers() {
	http.Handle("/bower_components/", http.FileServer(http.Dir("/vagrant/client")))
	http.Handle("/styles/", http.FileServer(http.Dir("/vagrant/client/app")))
	http.Handle("/scripts/", http.FileServer(http.Dir("/vagrant/client/app")))
	http.Handle("/images/", http.FileServer(http.Dir("/vagrant/client/app")))
	http.Handle("/views/", http.FileServer(http.Dir("/vagrant/client/app")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/vagrant/client/app/index.html")
	})
}

func main() {
	dbMaster.start()
	defer dbMaster.stop()

	setApiHandlers()
	setStaticHandlers()
	http.ListenAndServe(":8080", nil)
}
