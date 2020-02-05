package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/kanhaiya15/gra/cfg"
	"github.com/kanhaiya15/gra/cfg/dbs/kmysql"
	"github.com/kanhaiya15/gra/cfg/dbs/kredis"
	"github.com/kanhaiya15/gra/pkg/klog"
)

func main() {
	klog.NewLogger()
	cfg.NewConfig("./cfg", "pre-live")
	kredis.NewConfig()
	kmysql.NewConfig()
	cleanupHook()

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	s.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":3333", r))
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		klog.SLogger.Sync()
		kredis.Pool.Close()
		kmysql.Pool.Close()
		os.Exit(0)
	}()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
