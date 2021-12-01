package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/edgestats/edgestats-server/data"
	"github.com/edgestats/edgestats-server/handlers"
	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	srvPort       = "8000"
	srvLogDir     = "./logs"
	srvAccessLogs = "./logs/access.log"
	srvErrorLogs  = "./logs/error.log"
)

func main() {
	// set server log dir
	if err := os.MkdirAll(srvLogDir, os.ModePerm); err != nil {
		log.Fatalf("Error starting logs: %s\n", err)
	}

	// set server log files
	accessLogFile, err := os.OpenFile(srvAccessLogs, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.FileMode(0664))
	if err != nil {
		log.Fatalf("Error starting logs: %s\n", err)
	}
	defer accessLogFile.Close()

	errorLogFile, err := os.OpenFile(srvErrorLogs, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.FileMode(0664))
	if err != nil {
		log.Fatalf("Error starting logs: %s\n", err)
	}
	defer errorLogFile.Close()

	l := log.New(errorLogFile, "edgestats ", log.LstdFlags)
	h := handlers.NewHandler(l)

	// close db session
	defer data.DB.Close()

	sm := mux.NewRouter()

	// broadcasts endpoints
	sm.HandleFunc("/stats/uptimes/broadcasts", h.CreateUMBroadcast).Methods(http.MethodPost)
	sm.HandleFunc("/stats/uptimes/broadcasts", h.GetUMBroadcasts).Methods(http.MethodGet) // select * query
	sm.HandleFunc("/stats/uptimes/broadcasts/{addr}/{min}", h.GetUMBroadcastsByAddrByRange).Methods(http.MethodGet)
	sm.HandleFunc("/stats/uptimes/broadcasts/{addr}/{min}/{max}", h.GetUMBroadcastsByAddrByRange).Methods(http.MethodGet)

	// peers endpoints
	sm.HandleFunc("/stats/uptimes/peers", h.CreateNumPeers).Methods(http.MethodPost)
	sm.HandleFunc("/stats/uptimes/peers", h.GetNumPeers).Methods(http.MethodGet)
	sm.HandleFunc("/stats/uptimes/peers/{addr}/{min}", h.GetNumPeersByAddrByRange).Methods(http.MethodGet)
	sm.HandleFunc("/stats/uptimes/peers/{addr}/{min}/{max}", h.GetNumPeersByAddrByRange).Methods(http.MethodGet)

	// clusters endpoints
	sm.HandleFunc("/clusters/uptimes/broadcasts/{addrs}", h.GetUMBroadcastsByCluster).Methods(http.MethodGet)
	sm.HandleFunc("/clusters/uptimes/peers/{addrs}", h.GetNumPeersByCluster).Methods(http.MethodGet)

	// blocks endpoints
	sm.HandleFunc("/stats/blocks", h.CreateBlock).Methods(http.MethodPost)
	sm.HandleFunc("/stats/blocks/{min}", h.GetBlocksByRange).Methods(http.MethodGet)
	sm.HandleFunc("/stats/blocks/{min}/{max}", h.GetBlocksByRange).Methods(http.MethodGet)
	sm.HandleFunc("/stats/blocks/misses/{addr}/{min}", h.GetMissedBlocksByAddrByRange).Methods(http.MethodGet)
	sm.HandleFunc("/stats/blocks/misses/{addr}/{min}/{max}", h.GetMissedBlocksByAddrByRange).Methods(http.MethodGet)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%v", srvPort),
		Handler:      gh.LoggingHandler(accessLogFile, h.MiddlewareAuthz(gh.RecoveryHandler()(sm))),
		ErrorLog:     l,
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 6 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		fmt.Printf("Initializing server on port: %s\n", srvPort)

		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			l.Fatalf("Error starting server: %s\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	sig := <-ch
	fmt.Printf("Recieved %s signal, shutting down...\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
