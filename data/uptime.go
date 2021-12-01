package data

import (
	"encoding/json"
	"io"
	"time"
)

type UMBroadcast struct {
	Block           string    `json:"block"`
	Height          int       `json:"height"`
	Addr            string    `json:"address"`
	Signature       string    `json:"signature"`
	Timestamp       int       `json:"timestamp"`
	NumPeers        int       `json:"num_peers"`
	SufficientPeers int       `json:"sufficient_peers"`
	CreatedAt       time.Time `json:"created_at"`
}

func NewUMBroadcast() *UMBroadcast {
	return &UMBroadcast{}
}

func (um *UMBroadcast) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(um)
}

func (um *UMBroadcast) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(um)
}

func (um *UMBroadcast) CreateUMBroadcast() error {
	// write to stats current table
	if err := um.updateUMBroadcasts(); err != nil {
		return err
	}

	// write to stats history table
	if err := um.createUMBroadcastsByAddr(); err != nil {
		return err
	}

	// write block to blocks table
	bk := NewBlock()
	bk.Height = um.Height
	if err := bk.CreateBlock(); err != nil {
		// return err // may return 400 error if block not yet in explorer
	}

	return nil
}

func (um *UMBroadcast) updateUMBroadcasts() error {
	// set key & value
	k := um.Addr
	v, err := json.Marshal(um)
	if err != nil {
		return err
	}

	// write data to db
	if err := writeData([]byte(statsUptimesBroadcasts), []byte(k), []byte(v)); err != nil {
		return err
	}

	return nil
}

func (um *UMBroadcast) createUMBroadcastsByAddr() error {
	// set key & value
	k := um.CreatedAt.Format(time.RFC3339)
	v, err := json.Marshal(um)
	if err != nil {
		return err
	}

	// write data to db
	if err := writeNestedData([]byte(statsUptimesBroadcatsByAddr), []byte(um.Addr), []byte(k), []byte(v)); err != nil {
		return err
	}

	return nil
}

type UMBroadcasts []*UMBroadcast

func NewUMBroadcasts() *UMBroadcasts {
	return &UMBroadcasts{}
}

func (um *UMBroadcasts) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(um)
}

func (um *UMBroadcasts) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(um)
}

func (uml *UMBroadcasts) GetUMBroadcasts() error {
	// read data from db
	buf, err := scanData([]byte(statsUptimesBroadcasts))
	if err != nil {
		return err
	}

	// unmarshal data to struct
	return uml.unmarshalData(buf)
}

func (uml *UMBroadcasts) GetUMBroadcastsByAddr(addr string) error {
	// read data from db
	buf, err := scanNestedData([]byte(statsUptimesBroadcatsByAddr), []byte(addr))
	if err != nil {
		return err
	}

	// unmarshal data to struct
	return uml.unmarshalData(buf)
}

func (uml *UMBroadcasts) GetUMBroadcastsByAddrByRange(addr, min, max string) error {
	// validate times
	if err := validateTimes(min, max); err != nil {
		return err
	}

	// read data from db
	buf, err := scanNestedDataByRange([]byte(statsUptimesBroadcatsByAddr), []byte(addr), []byte(min), []byte(max))
	if err != nil {
		return err
	}

	// unmarshal data to struct
	return uml.unmarshalData(buf)
}

func (uml *UMBroadcasts) GetUMBroadcastsByCluster(addrs string) error {
	// split addrs string
	addrl := splitAddrs(addrs)

	// read data from db
	var buf [][]byte
	for _, addr := range addrl {
		b, err := readData([]byte(statsUptimesBroadcasts), []byte(addr))
		if err != nil {
			return err // consider skip return and read as many as possible
		}

		buf = append(buf, b)
	}

	// unmarshal data to struct
	return uml.unmarshalData(buf)
}

func (uml *UMBroadcasts) unmarshalData(buf [][]byte) error {
	for _, b := range buf {
		um := NewUMBroadcast()
		if err := json.Unmarshal(b, um); err != nil {
			return err
		}

		*uml = append(*uml, um)
	}

	return nil
}
