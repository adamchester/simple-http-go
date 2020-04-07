package main

import (
	"github.com/nullseed/logruseq"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
)

func init() {
	log.AddHook(logruseq.NewSeqHook("http://localhost:5341"))
}

func main() {

	log.Info("hello server")

	n := negroni.New()
	mux := http.NewServeMux()
	n.UseHandler(mux)
	n.Use(negroni.HandlerFunc(MyMiddleware))

	mux.HandleFunc("/", HelloServer)
	mux.HandleFunc("/api", ApiServer)

	http.ListenAndServe(":8080", n)

	log.Info("Closing down.. bye!")
}
