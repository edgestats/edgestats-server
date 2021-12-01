package handlers

import (
	"net/http"

	"github.com/edgestats/edgestats-server/data"
	"github.com/gorilla/mux"
)

func (h *Handler) CreateNumPeers(w http.ResponseWriter, r *http.Request) {
	// get request body
	p2p := data.NewP2P()
	if err := p2p.FromJSON(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// update db collection
	if err := p2p.CreateNumPeers(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetNumPeers(w http.ResponseWriter, r *http.Request) {
	// get data from db
	p2p := data.NewP2Ps()
	if err := p2p.GetNumPeers(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// encode to json byte array
	if err := p2p.ToJSON(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetNumPeersByAddr(w http.ResponseWriter, r *http.Request) {
	// get path params
	pp := mux.Vars(r)

	// validate params
	if len(pp) != 1 {
		http.Error(w, "error with request params", http.StatusBadRequest)
		return
	}
	if err := isValidAddr(pp["addr"]); err != nil {
		// h.l.Println("error with address")
		http.Error(w, "error with address", http.StatusBadRequest)
		return
	}

	// get data from db
	p2p := data.NewP2Ps()
	if err := p2p.GetNumPeersByAddr(pp["addr"]); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// encode to json byte array
	if err := p2p.ToJSON(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetNumPeersByAddrByRange(w http.ResponseWriter, r *http.Request) {
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
	p2p := data.NewP2Ps()
	if err := p2p.GetNumPeersByAddrByRange(pp["addr"], pp["min"], pp["max"]); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// encode to json byte array
	if err := p2p.ToJSON(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetNumPeersByCluster(w http.ResponseWriter, r *http.Request) {
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
	p2p := data.NewP2Ps()
	if err := p2p.GetNumPeersByCluster(pp["addrs"]); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// encode to json byte array
	if err := p2p.ToJSON(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
