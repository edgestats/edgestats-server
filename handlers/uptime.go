package handlers

import (
	"net/http"

	"github.com/edgestats/edgestats-server/data"
	"github.com/gorilla/mux"
)

func (h *Handler) CreateUMBroadcast(w http.ResponseWriter, r *http.Request) {
	// get request body
	um := data.NewUMBroadcast()
	if err := um.FromJSON(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// update db collection
	if err := um.CreateUMBroadcast(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetUMBroadcasts(w http.ResponseWriter, r *http.Request) {
	// get data from db
	um := data.NewUMBroadcasts()
	if err := um.GetUMBroadcasts(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// encode to json byte array
	if err := um.ToJSON(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUMBroadcastsByAddr(w http.ResponseWriter, r *http.Request) {
	// get path params
	pp := mux.Vars(r)

	// validate params
	if len(pp) != 1 {
		http.Error(w, "error with request params", http.StatusBadRequest)
		return
	}
	if err := isValidAddr(pp["addr"]); err != nil {
		http.Error(w, "error with address", http.StatusBadRequest)
		return
	}

	// get data from db
	um := data.NewUMBroadcasts()
	if err := um.GetUMBroadcastsByAddr(pp["addr"]); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// encode to json byte array
	if err := um.ToJSON(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUMBroadcastsByAddrByRange(w http.ResponseWriter, r *http.Request) {
	// get path params
	pp := mux.Vars(r)

	// validate params
	if len(pp) < 2 || len(pp) > 3 { // if !(2 <= len(pp) <= 3)
		http.Error(w, "error with request params", http.StatusBadRequest)
		return
	}
	if err := isValidAddr(pp["addr"]); err != nil {
		http.Error(w, "error with address", http.StatusBadRequest)
		return
	}
	if err := areValidTimes(pp["min"], pp["max"]); err != nil {
		http.Error(w, "error with times", http.StatusBadRequest)
		return
	}

	// get data from db
	um := data.NewUMBroadcasts()
	if err := um.GetUMBroadcastsByAddrByRange(pp["addr"], pp["min"], pp["max"]); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// encode to json byte array
	if err := um.ToJSON(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUMBroadcastsByCluster(w http.ResponseWriter, r *http.Request) {
	// get path params
	pp := mux.Vars(r)

	// validate params
	if len(pp) != 1 {
		http.Error(w, "error with request params", http.StatusBadRequest)
		return
	}
	if err := areValidAddrs(pp["addrs"]); err != nil {
		http.Error(w, "error with address", http.StatusBadRequest)
		return
	}

	// get data from db
	um := data.NewUMBroadcasts()
	if err := um.GetUMBroadcastsByCluster(pp["addrs"]); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// encode to json byte array
	if err := um.ToJSON(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
