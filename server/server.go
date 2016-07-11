package server

import (
	"log"
	"net/http"
	"os"
)

/*Start starts a web server */
func Start() {
	projdir := os.Getenv("GOPATH") + "/src/github.com/rprakashg/tageditor"
	pubdir := projdir + "/public/"
	certdir := projdir + "/server/certs"

	router := NewRouter()
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(pubdir))))
	http.Handle("/", router)
	err := http.ListenAndServeTLS(":10443", certdir+"/cert.pem", certdir+"/key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
