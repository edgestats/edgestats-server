package data

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"
)

func TestNewP2P(t *testing.T) {
	want := &P2P{}
	got := NewP2P()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.NewP2P() returned: %v, wanted: %v", got, want)
	}
}

func TestP2PFromJSON(t *testing.T) {
	b := []byte(`{"address":"0x5b8c84db6f40bf45","num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T22:49:51.387Z"}`)
	buf := bytes.NewReader(b)
	vt, _ := time.Parse(time.RFC3339, "2021-11-28T22:49:51.387Z")

	want := &P2P{
		Addr:            "0x5b8c84db6f40bf45",
		NumPeers:        16,
		SufficientPeers: 16,
		CreatedAt:       vt,
	}
	got := NewP2P()
	if err := got.FromJSON(buf); err != nil {
		t.Fatalf("data.P2PFromJSON() returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.P2PFromJSON() returned: %v, wanted: %v", got, want)
	}
}

func TestP2PToJSON(t *testing.T) {
	vt, _ := time.Parse(time.RFC3339, "2021-11-28T22:49:51.387Z")

	p2p := &P2P{
		Addr:            "0x5b8c84db6f40bf45",
		NumPeers:        16,
		SufficientPeers: 16,
		CreatedAt:       vt,
	}

	want := []byte(`{"address":"0x5b8c84db6f40bf45","num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T22:49:51.387Z"}`)
	buf := &bytes.Buffer{}
	if err := p2p.ToJSON(buf); err != nil {
		t.Fatalf("data.P2PToJSON() returned error: %v", err)
	}
	got, err := io.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}
	got = got[:len(got)-1] // to remove new line character

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.P2PToJSON() returned: %s, wanted: %s", got, want)
	}
}

func TestP2PCreateNumPeers(t *testing.T) {}

func TestP2PUpdateNumPeers(t *testing.T) {}

func TestP2PCreateNumPeersByAddr(t *testing.T) {}

func TestNewP2Ps(t *testing.T) {
	want := &P2Ps{}
	got := NewP2Ps()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.NewP2Ps() returned: %v, wanted: %v", got, want)
	}
}

func TestP2PsFromJSON(t *testing.T) {
	b := []byte(`[{"address":"0x5b8c84db6f40bf45","num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T22:49:51.387Z"},{"address":"0x5b8c84db6f40bf45","num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T23:00:26.989Z"}]`)
	buf := bytes.NewReader(b)
	vt0, _ := time.Parse(time.RFC3339, "2021-11-28T22:49:51.387Z")
	vt1, _ := time.Parse(time.RFC3339, "2021-11-28T23:00:26.989Z")

	want := &P2Ps{
		{
			Addr:            "0x5b8c84db6f40bf45",
			NumPeers:        16,
			SufficientPeers: 16,
			CreatedAt:       vt0,
		},
		{
			Addr:            "0x5b8c84db6f40bf45",
			NumPeers:        16,
			SufficientPeers: 16,
			CreatedAt:       vt1,
		},
	}
	got := NewP2Ps()
	if err := got.FromJSON(buf); err != nil {
		t.Fatalf("data.P2PsFromJSON() returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.P2PsFromJSON() returned: %v, wanted: %v", got, want)
	}
}

func TestP2PsToJSON(t *testing.T) {
	vt0, _ := time.Parse(time.RFC3339, "2021-11-28T22:49:51.387Z")
	vt1, _ := time.Parse(time.RFC3339, "2021-11-28T23:00:26.989Z")

	p2pl := &P2Ps{
		{
			Addr:            "0x5b8c84db6f40bf45",
			NumPeers:        16,
			SufficientPeers: 16,
			CreatedAt:       vt0,
		},
		{
			Addr:            "0x5b8c84db6f40bf45",
			NumPeers:        16,
			SufficientPeers: 16,
			CreatedAt:       vt1,
		},
	}

	want := []byte(`[{"address":"0x5b8c84db6f40bf45","num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T22:49:51.387Z"},{"address":"0x5b8c84db6f40bf45","num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T23:00:26.989Z"}]`)
	buf := &bytes.Buffer{}
	if err := p2pl.ToJSON(buf); err != nil {
		t.Fatalf("data.P2PsToJSON() returned error: %v", err)
	}
	got, err := io.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}
	got = got[:len(got)-1] // to remove new line character

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.P2PsToJSON() returned: %s, wanted: %s", got, want)
	}
}

func TestP2PsGetNumPeers(t *testing.T) {}

func TestP2PsGetNumPeersByAddr(t *testing.T) {}

func TestP2PsGetNumPeersByAddrByRange(t *testing.T) {}

func TestP2PsGetNumPeersByCluster(t *testing.T) {}

func TestP2PsUnmarshalData(t *testing.T) {}
