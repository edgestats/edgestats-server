package data

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"
)

func TestNewExplorerBlock(t *testing.T) {
	want := &explorerBlock{}
	got := newExplorerBlock()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.newExplorerBlock() returned: %v, wanted %v", got, want)
	}
}

func TestNewExplorerBlockFromJSON(t *testing.T) {
	b := []byte(`{"type":"block","body":{"epoch":"13111650","status":4,"height":13028501,"timestamp":"1638067374","hash":"0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9","parent_hash":"0x1865edbb30a19fb284b8672fd5841058c4afca1a21eb460e968497df22d85f5b","proposer":"0x80eab22e27d4b94511f5906484369b868d6552d2","state_hash":"0x8d6b91a8f27ed447c58e8ca9b5344b9a5fe4a7406153d1e4e7e17a70c10b264a","transactions_hash":"0x32f6d278bb7fbc2177a55a9dfb97e7794d46d7cd5e4fe92a02b5994bcbe4a17d","num_txs":1,"txs":[{"hash":"0x504712b63f6db33941a85aaccb52d425f9ca522ad79ad68dc7ec893ab758209b","type":0,"raw":null}],"total_deposited_guardian_stakes":"3.95948760391483100177603732e+26","total_voted_guardian_stakes":"3.83602073941373533569678392e+26"},"totalBlocksNumber":13034993}`)
	buf := bytes.NewReader(b)

	want := &explorerBlock{
		explorerBlockBody: explorerBlockBody{
			Epoch:     "13111650",
			Height:    13028501,
			Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
			Timestamp: "1638067374",
		},
	}
	got := newExplorerBlock()
	if err := got.FromJSON(buf); err != nil {
		t.Fatalf("data.explorerBlockFromJSON() returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.BlockFromJSON() returned: %v, wanted: %v", got, want)
	}
}

func TestNewBlock(t *testing.T) {
	want := &Block{}
	got := NewBlock()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.NewBlock() returned: %v, wanted: %v", got, want)
	}
}

func TestBlockFromJSON(t *testing.T) {
	b := []byte(`{"epoch":13111650,"height":13028501,"timestamp":1638067374,"created_at":"2021-11-28T02:42:54Z","hash":"0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9"}`)
	buf := bytes.NewReader(b)
	vt, _ := time.Parse(time.RFC3339, "2021-11-28T02:42:54Z")

	want := &Block{
		Epoch:     13111650,
		Height:    13028501,
		Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
		Timestamp: 1638067374,
		CreatedAt: vt,
	}
	got := NewBlock()
	if err := got.FromJSON(buf); err != nil {
		t.Fatalf("data.BlockFromJSON() returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.BlockFromJSON() returned: %v, wanted: %v", got, want)
	}
}

func TestBlockToJSON(t *testing.T) {
	vt, _ := time.Parse(time.RFC3339, "2021-11-28T02:42:54Z")

	bk := &Block{
		Epoch:     13111650,
		Height:    13028501,
		Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
		Timestamp: 1638067374,
		CreatedAt: vt,
	}

	want := []byte(`{"epoch":13111650,"height":13028501,"hash":"0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9","timestamp":1638067374,"created_at":"2021-11-28T02:42:54Z"}`)
	buf := &bytes.Buffer{}
	if err := bk.ToJSON(buf); err != nil {
		t.Fatalf("data.BlockToJSON() returned error: %v", err)
	}
	got, err := io.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}
	got = got[:len(got)-1] // to remove new line character

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.BlockToJSON() returned: %s, wanted: %s", got, want)
	}
}

func TestCreateBlock(t *testing.T) {}

func TestQueryExplorerBlock(t *testing.T) {
	// test valid block
	want := &explorerBlock{
		explorerBlockBody: explorerBlockBody{
			Epoch:     "13111650",
			Height:    13028501,
			Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
			Timestamp: "1638067374",
		},
	}
	got, err := queryExplorerBlock(13028501)
	if err != nil {
		t.Fatalf("data.queryExplorerBlock() returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.queryExplorerBlock() returned: %v, wanted: %v", got, want)
	}
}

func TestBlockFromExplorerBlock(t *testing.T) {
	vt, _ := time.Parse(time.RFC3339, "2021-11-28T02:42:54Z")
	ek := &explorerBlock{
		explorerBlockBody: explorerBlockBody{
			Epoch:     "13111650",
			Height:    13028501,
			Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
			Timestamp: "1638067374",
		},
	}

	want := &Block{
		Epoch:     13111650,
		Height:    13028501,
		Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
		Timestamp: 1638067374,
		CreatedAt: vt,
	}
	got := NewBlock()
	if err := got.fromExplorerBlock(ek); err != nil {
		t.Fatalf("data.fromExplorerBlock() returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.fromExplorerBlock() returned: %v, wanted: %v", got, want)
	}

	// test epoch error
	ek = &explorerBlock{
		explorerBlockBody: explorerBlockBody{
			Epoch:     "bad string",
			Height:    13028501,
			Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
			Timestamp: "1638067374",
		},
	}
	got = NewBlock()
	if err := got.fromExplorerBlock(ek); err == nil {
		t.Fatalf("data.fromExplorerBlock() returned: %v, wanted error: %v", got, err)
	}

	// test timestamp error
	ek = &explorerBlock{
		explorerBlockBody: explorerBlockBody{
			Epoch:     "13111650",
			Height:    13028501,
			Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
			Timestamp: "bad string",
		},
	}
	got = NewBlock()
	if err := got.fromExplorerBlock(ek); err == nil {
		t.Fatalf("data.fromExplorerBlock() returned: %v, wanted error: %v", got, err)
	}
}

func TestBlockcreateBlock(t *testing.T) {}

func TestNewBlocks(t *testing.T) {
	want := &Blocks{}
	got := NewBlocks()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.NewBlocks() returned: %v, wanted: %v", got, want)
	}
}

func TestBlocksFromJSON(t *testing.T) {
	b := []byte(`[{"epoch":13111650,"height":13028501,"timestamp":1638067374,"created_at":"2021-11-28T02:42:54Z","hash":"0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9"},{"epoch":13111750,"height":13028601,"timestamp":1638067992,"created_at":"2021-11-28T02:53:12Z","hash":"0x27846dce2b14d5dead49ca1ecdd2420c1c9bd9c29b1b98394dcfc36738d47e1c"}]`)
	buf := bytes.NewReader(b)
	vt0, _ := time.Parse(time.RFC3339, "2021-11-28T02:42:54Z")
	vt1, _ := time.Parse(time.RFC3339, "2021-11-28T02:53:12Z")

	want := &Blocks{
		{
			Epoch:     13111650,
			Height:    13028501,
			Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
			Timestamp: 1638067374,
			CreatedAt: vt0,
		},
		{
			Epoch:     13111750,
			Height:    13028601,
			Hash:      "0x27846dce2b14d5dead49ca1ecdd2420c1c9bd9c29b1b98394dcfc36738d47e1c",
			Timestamp: 1638067992,
			CreatedAt: vt1,
		},
	}
	got := NewBlocks()
	if err := got.FromJSON(buf); err != nil {
		t.Fatalf("data.BlocksFromJSON() returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.BlocksFromJSON() returned: %v, wanted: %v", got, want)
	}
}

func TestBlocksToJSON(t *testing.T) {
	vt0, _ := time.Parse(time.RFC3339, "2021-11-28T02:42:54Z")
	vt1, _ := time.Parse(time.RFC3339, "2021-11-28T02:53:12Z")

	bkl := &Blocks{
		{
			Epoch:     13111650,
			Height:    13028501,
			Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
			Timestamp: 1638067374,
			CreatedAt: vt0,
		},
		{
			Epoch:     13111750,
			Height:    13028601,
			Hash:      "0x27846dce2b14d5dead49ca1ecdd2420c1c9bd9c29b1b98394dcfc36738d47e1c",
			Timestamp: 1638067992,
			CreatedAt: vt1,
		},
	}

	want := []byte(`[{"epoch":13111650,"height":13028501,"hash":"0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9","timestamp":1638067374,"created_at":"2021-11-28T02:42:54Z"},{"epoch":13111750,"height":13028601,"hash":"0x27846dce2b14d5dead49ca1ecdd2420c1c9bd9c29b1b98394dcfc36738d47e1c","timestamp":1638067992,"created_at":"2021-11-28T02:53:12Z"}]`)
	buf := &bytes.Buffer{}
	if err := bkl.ToJSON(buf); err != nil {
		t.Fatalf("data.BlocksToJSON() returned error: %v", err)
	}
	got, err := io.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}
	got = got[:len(got)-1] // to remove new line character

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.BlocksToJSON() returned: %s, wanted: %s", got, want)
	}
}

func TestBlocksGetBlocks(t *testing.T) {}

func TestBlocksGetBlocksByRange(t *testing.T) {}

func TestBlocksUnmarshalData(t *testing.T) {
	buf := [][]byte{
		[]byte(`{"epoch":13111650,"height":13028501,"timestamp":1638067374,"created_at":"2021-11-28T02:42:54Z","hash":"0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9"}`),
		[]byte(`{"epoch":13111750,"height":13028601,"timestamp":1638067992,"created_at":"2021-11-28T02:53:12Z","hash":"0x27846dce2b14d5dead49ca1ecdd2420c1c9bd9c29b1b98394dcfc36738d47e1c"}`),
	}
	vt0, _ := time.Parse(time.RFC3339, "2021-11-28T02:42:54Z")
	vt1, _ := time.Parse(time.RFC3339, "2021-11-28T02:53:12Z")

	want := &Blocks{
		{
			Epoch:     13111650,
			Height:    13028501,
			Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
			Timestamp: 1638067374,
			CreatedAt: vt0,
		},
		{
			Epoch:     13111750,
			Height:    13028601,
			Hash:      "0x27846dce2b14d5dead49ca1ecdd2420c1c9bd9c29b1b98394dcfc36738d47e1c",
			Timestamp: 1638067992,
			CreatedAt: vt1,
		},
	}

	got := NewBlocks()
	if err := got.unmarshalData(buf); err != nil {
		t.Fatalf("data.BlocksUnmarshalData() returned error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.BlocksUnmarshalData() returned: %v, wanted: %v", got, want)
	}

	// test error buf
	buf = [][]byte{
		[]byte(`{"epoch":13111650,"height":13028501,"timestamp":1638067374,"created_at":"2021-11-28T02:42:54Z","hash":"0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9"}`),
		[]byte(`{"epoch":13111750,"height":"13028601","timestamp":1638067992,"created_at":"2021-11-28T02:53:12Z","hash":"0x27846dce2b14d5dead49ca1ecdd2420c1c9bd9c29b1b98394dcfc36738d47e1c"}`),
	}

	got = NewBlocks()
	if err := got.unmarshalData(buf); err == nil {
		t.Fatalf("data.BlocksUnmarshalData() returned: %v, wanted error: %v", got, err)
	}
}

func TestBlocksGetMissedBlocksByAddrByRange(t *testing.T) {}

func TestBlocksGetMissedBlocks(t *testing.T) {
	testUptimes := &UMBroadcasts{
		{
			Height:    12524001,
			Addr:      "0x1a2b3c",
			Timestamp: 1634919689,
		},
		{
			Height:    12523901,
			Addr:      "0x1a2b3c",
			Timestamp: 1634919069,
		},
	}
	testBlocks := &Blocks{
		{
			Epoch:     12345,
			Hash:      "0x32bd3a2c75696bc0faf8a2e0a59034984c6b86546850ed2f6270fda0feb3cdce",
			Height:    12524001,
			Timestamp: 1634919689,
			CreatedAt: time.Unix(1634919689, 0),
		},
		{
			Epoch:     12345,
			Hash:      "0x5bea24587cfe6509198cd05f26c173b0a68fa539c9a8b4a00e3186f8529157a0",
			Height:    12523901,
			Timestamp: 1634919069,
			CreatedAt: time.Unix(1634919069, 0),
		},
		{
			Epoch:     13111650,
			Height:    13028501,
			Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
			Timestamp: 1638067374,
			CreatedAt: time.Unix(1638067374, 0),
		},
		{
			Epoch:     13111750,
			Height:    13028601,
			Hash:      "0x27846dce2b14d5dead49ca1ecdd2420c1c9bd9c29b1b98394dcfc36738d47e1c",
			Timestamp: 1638067992,
			CreatedAt: time.Unix(1638067992, 0),
		},
	}

	want := &Blocks{
		{
			Epoch:     13111650,
			Height:    13028501,
			Hash:      "0x5b8c84db6f40bf45722e62f2d49cf7bb247e4131ad488f44cf65a20a911a18d9",
			Timestamp: 1638067374,
			CreatedAt: time.Unix(1638067374, 0),
		},
		{
			Epoch:     13111750,
			Height:    13028601,
			Hash:      "0x27846dce2b14d5dead49ca1ecdd2420c1c9bd9c29b1b98394dcfc36738d47e1c",
			Timestamp: 1638067992,
			CreatedAt: time.Unix(1638067992, 0),
		},
	}
	got := NewBlocks()
	got.getMissedBlocks(testBlocks, testUptimes)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("getMissedBlocks() got: %#v\n, wnt: %#v\n", got, want)
	}
}
