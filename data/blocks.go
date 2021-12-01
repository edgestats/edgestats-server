package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	explorerURL = "https://explorer.thetatoken.org:8443/api"
)

type explorerBlock struct {
	explorerBlockBody `json:"body"`
}

type explorerBlockBody struct {
	Epoch     string `json:"epoch"`
	Height    int    `json:"height"`
	Hash      string `json:"hash"`
	Timestamp string `json:"timestamp"`
}

func newExplorerBlock() *explorerBlock {
	return &explorerBlock{}
}

func (ek *explorerBlock) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(ek)
}

type Block struct {
	Epoch     int       `json:"epoch"`
	Height    int       `json:"height"`
	Hash      string    `json:"hash"`
	Timestamp int       `json:"timestamp"`
	CreatedAt time.Time `json:"created_at"`
}

func NewBlock() *Block {
	return &Block{}
}

func (bk *Block) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(bk)
}

func (bk *Block) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(bk)
}

func (bk *Block) unmarshalData(b []byte) error {
	return json.Unmarshal(b, bk)
}

func (bk *Block) CreateBlock() error {
	// check if create needed
	if ok := queryLastBlock(bk.Height); ok {
		return nil // go easy on explorer
	}

	// query explorer block
	ek, err := queryExplorerBlock(bk.Height)
	if err != nil {
		return err
	}

	// map explorerBlock to Block
	if err := bk.fromExplorerBlock(ek); err != nil {
		return err
	}

	// write to blocks table
	if err := bk.createBlock(); err != nil {
		return err
	}

	return nil
}

func queryLastBlock(h int) bool {
	// read data from db
	b, err := readLastData([]byte(statsBlocks))
	if err != nil {
		return false
	}

	// unmarshal data to struct
	bk := NewBlock()
	if err := bk.unmarshalData(b); err != nil {
		return false
	}

	// decide if create needed
	if bk.Height != h {
		return false
	}

	return true
}

func queryExplorerBlock(h int) (*explorerBlock, error) {
	url := fmt.Sprintf("%s/block/%s", explorerURL, strconv.Itoa(h))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// unmarshal data explorerBlock
	ek := newExplorerBlock()
	if err := ek.FromJSON(resp.Body); err != nil {
		return nil, err
	}

	return ek, nil
}

func (bk *Block) fromExplorerBlock(ek *explorerBlock) error {
	ve, err := strconv.Atoi(ek.Epoch)
	if err != nil {
		return err
	}

	vt, err := strconv.Atoi(ek.Timestamp)
	if err != nil {
		return err
	}

	bk.Epoch = ve
	bk.Height = ek.Height
	bk.Hash = ek.Hash
	bk.Timestamp = vt
	bk.CreatedAt = time.Unix(int64(vt), 0).UTC()

	return nil
}

func (bk *Block) createBlock() error {
	// set key & value
	k := bk.CreatedAt.Format(time.RFC3339)
	v, err := json.Marshal(bk)
	if err != nil {
		return err
	}

	// write data to db
	if err := writeData([]byte(statsBlocks), []byte(k), []byte(v)); err != nil {
		return err
	}

	return nil
}

type Blocks []*Block

func NewBlocks() *Blocks {
	return &Blocks{}
}

func (bk *Blocks) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(bk)
}

func (bk *Blocks) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(bk)
}

func (bkl *Blocks) GetBlocks() error {
	// read data from db
	buf, err := scanData([]byte(statsBlocks))
	if err != nil {
		return err
	}

	// unmarshal data to struct
	return bkl.unmarshalData(buf)
}

func (bkl *Blocks) GetBlocksByRange(min, max string) error {
	// validate times
	if err := validateTimes(min, max); err != nil {
		return err
	}

	// read data from db
	buf, err := scanDataByRange([]byte(statsBlocks), []byte(min), []byte(max))
	if err != nil {
		return err
	}

	// unmarshal data to struct
	return bkl.unmarshalData(buf)
}

func (bkl *Blocks) unmarshalData(buf [][]byte) error {
	for _, b := range buf {
		bk := NewBlock()
		if err := json.Unmarshal(b, bk); err != nil {
			return err
		}

		*bkl = append(*bkl, bk)
	}

	return nil
}

func (bkl *Blocks) GetMissedBlocksByAddrByRange(addr, min, max string) error {
	// get blocks by range // consider running in goroutine with channel
	vkl := NewBlocks()
	if err := vkl.GetBlocksByRange(min, max); err != nil {
		return err
	}

	// get broadcasts by range // consider running in goroutine with channel
	uml := NewUMBroadcasts()
	if err := uml.GetUMBroadcastsByAddrByRange(addr, min, max); err != nil {
		return err
	}

	// filter missed blocks
	bkl.getMissedBlocks(vkl, uml)

	return nil
}

func (bkl *Blocks) getMissedBlocks(vkl *Blocks, uml *UMBroadcasts) {
	// create map of broadcasts
	m := make(map[int]interface{})
	for _, v := range *uml {
		m[v.Height] = *v
	}

	// add to list if not in map
	for _, v := range *vkl {
		if m[v.Height] == nil {
			*bkl = append(*bkl, v)
		}
	}
}
