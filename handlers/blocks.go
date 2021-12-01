package handlers

import (
	"net/http"

	"github.com/edgestats/edgestats-server/data"
	"github.com/gorilla/mux"
)

func (h *Handler) CreateBlock(w http.ResponseWriter, r *http.Request) {
	// get request body
	bk := data.NewBlock()
	if err := bk.FromJSON(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// update db collection
	if err := bk.CreateBlock(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetBlocks(w http.ResponseWriter, r *http.Request) {
	// get data from db
	bk := data.NewBlocks()
	if err := bk.GetBlocks(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// encode to json byte array
	if err := bk.ToJSON(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetBlocksByRange(w http.ResponseWriter, r *http.Request) {
	// get path params
	pp := mux.Vars(r)

	// validate params
	if len(pp) < 1 || len(pp) > 2 { // if !(1 <= len(pp) <= 2)
		http.Error(w, "error with request params", http.StatusBadRequest)
		return
	}
	if err := areValidTimes(pp["min"], pp["max"]); err != nil {
		http.Error(w, "error with times", http.StatusBadRequest)
		return
	}

	// get data from db
	bk := data.NewBlocks()
	if err := bk.GetBlocksByRange(pp["min"], pp["max"]); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// encode to json byte array
	if err := bk.ToJSON(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetMissedBlocksByAddrByRange(w http.ResponseWriter, r *http.Request) {
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
	bk := data.NewBlocks()
	if err := bk.GetMissedBlocksByAddrByRange(pp["addr"], pp["min"], pp["max"]); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// set http response headers
	w.Header().Set("Content-Type", "application/json")

	// encode to json byte array
	if err := bk.ToJSON(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
