package data

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"
)

func TestNewUMBroadcast(t *testing.T) {
	want := &UMBroadcast{}
	got := NewUMBroadcast()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.NewUMBroadcast() returned: %v, wanted: %v", got, want)
	}
}

func TestUMBroadcastFromJSON(t *testing.T) {
	b := []byte(`{"block":"0x35d7858a13c4d76c9caac19e30ca7a7f13c26a0c494bb945561a343a9613acc7","height":13040101,"address":"0x5b8c84db6f40bf45","signature":"E1A035D7858A13C4D","timestamp":1638139521,"num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T22:49:51.387Z"}`)
	buf := bytes.NewReader(b)
	vt, _ := time.Parse(time.RFC3339, "2021-11-28T22:49:51.387Z")

	want := &UMBroadcast{
		Block:           "0x35d7858a13c4d76c9caac19e30ca7a7f13c26a0c494bb945561a343a9613acc7",
		Height:          13040101,
		Addr:            "0x5b8c84db6f40bf45",
		Signature:       "E1A035D7858A13C4D",
		Timestamp:       1638139521,
		NumPeers:        16,
		SufficientPeers: 16,
		CreatedAt:       vt,
	}
	got := NewUMBroadcast()
	if err := got.FromJSON(buf); err != nil {
		t.Fatalf("data.UMBroadcastFromJSON() returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.UMBroadcastFromJSON() returned: %v, wanted: %v", got, want)
	}
}

func TestUMBroadcastToJSON(t *testing.T) {
	vt, _ := time.Parse(time.RFC3339, "2021-11-28T22:49:51.387Z")

	um := &UMBroadcast{
		Block:           "0x35d7858a13c4d76c9caac19e30ca7a7f13c26a0c494bb945561a343a9613acc7",
		Height:          13040101,
		Addr:            "0x5b8c84db6f40bf45",
		Signature:       "E1A035D7858A13C4D",
		Timestamp:       1638139521,
		NumPeers:        16,
		SufficientPeers: 16,
		CreatedAt:       vt,
	}
	want := []byte(`{"block":"0x35d7858a13c4d76c9caac19e30ca7a7f13c26a0c494bb945561a343a9613acc7","height":13040101,"address":"0x5b8c84db6f40bf45","signature":"E1A035D7858A13C4D","timestamp":1638139521,"num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T22:49:51.387Z"}`)
	buf := &bytes.Buffer{}
	if err := um.ToJSON(buf); err != nil {
		t.Fatalf("data.UMBroadcastToJSON() returned error: %v", err)
	}
	got, err := io.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}
	got = got[:len(got)-1] // to remove new line character

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.UMBroadcastToJSON() returned: %s, wanted: %s", got, want)
	}
}

func TestUMBroadcastCreateUMBroadcast(t *testing.T) {}

func TestUMBroadcastUpdateUMBroadcasts(t *testing.T) {}

func TestUMBroadcastCreateUMBroadcastByAddr(t *testing.T) {}

func TestNewUMBroadcasts(t *testing.T) {
	want := &UMBroadcasts{}
	got := NewUMBroadcasts()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.NewUMBroadcasts() returned: %v, wanted: %v", got, want)
	}
}

func TestUMBroadcastsFromJSON(t *testing.T) {
	b := []byte(`[{"block":"0x35d7858a13c4d76c9caac19e30ca7a7f13c26a0c494bb945561a343a9613acc7","height":13040101,"address":"0x5b8c84db6f40bf45","signature":"E1A035D7858A13C4D","timestamp":1638139521,"num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T22:49:51.387Z"},{"block":"0xb6c7357e8372775fb81b57cdfac847945f1c4d8f13ce3ade6a4d9d3e053a1fca","height":13040201,"address":"0x5b8c84db6f40bf45","signature":"E1A035D7858A13C4D","timestamp":1638140156,"num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T23:00:26.989Z"}]`)
	buf := bytes.NewReader(b)
	vt0, _ := time.Parse(time.RFC3339, "2021-11-28T22:49:51.387Z")
	vt1, _ := time.Parse(time.RFC3339, "2021-11-28T23:00:26.989Z")

	want := &UMBroadcasts{
		{
			Block:           "0x35d7858a13c4d76c9caac19e30ca7a7f13c26a0c494bb945561a343a9613acc7",
			Height:          13040101,
			Addr:            "0x5b8c84db6f40bf45",
			Signature:       "E1A035D7858A13C4D",
			Timestamp:       1638139521,
			NumPeers:        16,
			SufficientPeers: 16,
			CreatedAt:       vt0,
		},
		{
			Block:           "0xb6c7357e8372775fb81b57cdfac847945f1c4d8f13ce3ade6a4d9d3e053a1fca",
			Height:          13040201,
			Addr:            "0x5b8c84db6f40bf45",
			Signature:       "E1A035D7858A13C4D",
			Timestamp:       1638140156,
			NumPeers:        16,
			SufficientPeers: 16,
			CreatedAt:       vt1,
		},
	}
	got := NewUMBroadcasts()
	if err := got.FromJSON(buf); err != nil {
		t.Fatalf("data.UMBroadcastFromJSON() returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.UMBroadcastFromJSON() returned: %v, wanted: %v", got, want)
	}
}

func TestUMBroadcastsToJSON(t *testing.T) {
	vt0, _ := time.Parse(time.RFC3339, "2021-11-28T22:49:51.387Z")
	vt1, _ := time.Parse(time.RFC3339, "2021-11-28T23:00:26.989Z")

	uml := &UMBroadcasts{
		{
			Block:           "0x35d7858a13c4d76c9caac19e30ca7a7f13c26a0c494bb945561a343a9613acc7",
			Height:          13040101,
			Addr:            "0x5b8c84db6f40bf45",
			Signature:       "E1A035D7858A13C4D",
			Timestamp:       1638139521,
			NumPeers:        16,
			SufficientPeers: 16,
			CreatedAt:       vt0,
		},
		{
			Block:           "0xb6c7357e8372775fb81b57cdfac847945f1c4d8f13ce3ade6a4d9d3e053a1fca",
			Height:          13040201,
			Addr:            "0x5b8c84db6f40bf45",
			Signature:       "E1A035D7858A13C4D",
			Timestamp:       1638140156,
			NumPeers:        16,
			SufficientPeers: 16,
			CreatedAt:       vt1,
		},
	}

	want := []byte(`[{"block":"0x35d7858a13c4d76c9caac19e30ca7a7f13c26a0c494bb945561a343a9613acc7","height":13040101,"address":"0x5b8c84db6f40bf45","signature":"E1A035D7858A13C4D","timestamp":1638139521,"num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T22:49:51.387Z"},{"block":"0xb6c7357e8372775fb81b57cdfac847945f1c4d8f13ce3ade6a4d9d3e053a1fca","height":13040201,"address":"0x5b8c84db6f40bf45","signature":"E1A035D7858A13C4D","timestamp":1638140156,"num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T23:00:26.989Z"}]`)
	buf := &bytes.Buffer{}
	if err := uml.ToJSON(buf); err != nil {
		t.Fatalf("data.UMBroadcastsToJSON() returned error: %v", err)
	}
	got, err := io.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}
	got = got[:len(got)-1] // to remove new line character

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.UMBroadcastsToJSON() returned: %s, wanted: %s", got, want)
	}
}

func TestUMBroadcastsGetUMBroadcasts(t *testing.T) {}

func TestUMBroadcastsGetUMBroadcastsByAddr(t *testing.T) {}

func TestUMBroadcastsGetUMBroadcastsByAddrByRange(t *testing.T) {}

func TestUMBroadcastsGetUMBroadcastsByCluster(t *testing.T) {}

func TestUMBroadcastsUnmarshalData(t *testing.T) {
	buf := [][]byte{
		[]byte(`{"block":"0x35d7858a13c4d76c9caac19e30ca7a7f13c26a0c494bb945561a343a9613acc7","height":13040101,"address":"0x5b8c84db6f40bf45","signature":"E1A035D7858A13C4D","timestamp":1638139521,"num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T22:49:51.387Z"}`),
		[]byte(`{"block":"0xb6c7357e8372775fb81b57cdfac847945f1c4d8f13ce3ade6a4d9d3e053a1fca","height":13040201,"address":"0x5b8c84db6f40bf45","signature":"E1A035D7858A13C4D","timestamp":1638140156,"num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T23:00:26.989Z"}`),
	}
	vt0, _ := time.Parse(time.RFC3339, "2021-11-28T22:49:51.387Z")
	vt1, _ := time.Parse(time.RFC3339, "2021-11-28T23:00:26.989Z")

	want := &UMBroadcasts{
		{
			Block:           "0x35d7858a13c4d76c9caac19e30ca7a7f13c26a0c494bb945561a343a9613acc7",
			Height:          13040101,
			Addr:            "0x5b8c84db6f40bf45",
			Signature:       "E1A035D7858A13C4D",
			Timestamp:       1638139521,
			NumPeers:        16,
			SufficientPeers: 16,
			CreatedAt:       vt0,
		},
		{
			Block:           "0xb6c7357e8372775fb81b57cdfac847945f1c4d8f13ce3ade6a4d9d3e053a1fca",
			Height:          13040201,
			Addr:            "0x5b8c84db6f40bf45",
			Signature:       "E1A035D7858A13C4D",
			Timestamp:       1638140156,
			NumPeers:        16,
			SufficientPeers: 16,
			CreatedAt:       vt1,
		},
	}

	got := NewUMBroadcasts()
	if err := got.unmarshalData(buf); err != nil {
		t.Fatalf("data.UMBroadcastsUnmarshalData() returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.UMBroadcastsUnmarshalData() returned: %v, wanted: %v", got, want)
	}

	// test error buf
	buf = [][]byte{
		[]byte(`{"block":"0x35d7858a13c4d76c9caac19e30ca7a7f13c26a0c494bb945561a343a9613acc7","height":13040101,"address":"0x5b8c84db6f40bf45","signature":"E1A035D7858A13C4D","timestamp":1638139521,"num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T22:49:51.387Z"}`),
		[]byte(`{"block":"0xb6c7357e8372775fb81b57cdfac847945f1c4d8f13ce3ade6a4d9d3e053a1fca","height":"13040201","address":"0x5b8c84db6f40bf45","signature":"E1A035D7858A13C4D","timestamp":1638140156,"num_peers":16,"sufficient_peers":16,"created_at":"2021-11-28T23:00:26.989Z"}`),
	}

	got = NewUMBroadcasts()
	if err := got.unmarshalData(buf); err == nil {
		t.Fatalf("data.UMBroadcastsUnmarshalData() returned: %v, wanted error: %v", got, err)
	}
}
