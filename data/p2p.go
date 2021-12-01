package data

import (
	"encoding/json"
	"io"
	"time"
)

type P2P struct {
	Addr            string    `json:"address"`
	NumPeers        int8      `json:"num_peers"`
	SufficientPeers int8      `json:"sufficient_peers"`
	CreatedAt       time.Time `json:"created_at"`
}

func NewP2P() *P2P {
	return &P2P{}
}

func (p2p *P2P) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p2p)
}

func (p2p *P2P) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p2p)
}

func (p2p *P2P) CreateNumPeers() error {
	// write to stats current table
	if err := p2p.updateNumPeers(); err != nil {
		return err
	}

	// write to stats history table
	if err := p2p.createNumPeersByAddr(); err != nil {
		return err
	}

	return nil
}

func (p2p *P2P) updateNumPeers() error {
	// set key & value
	k := p2p.Addr
	v, err := json.Marshal(p2p)
	if err != nil {
		return err
	}

	// write data to db
	if err := writeData([]byte(statsUptimesPeers), []byte(k), []byte(v)); err != nil {
		return err
	}

	return nil
}

func (p2p *P2P) createNumPeersByAddr() error {
	// set key & value
	k := p2p.CreatedAt.Format(time.RFC3339)
	v, err := json.Marshal(p2p)
	if err != nil {
		return err
	}

	// write to db
	if err := writeNestedData([]byte(statsUptimesPeersByAddr), []byte(p2p.Addr), []byte(k), []byte(v)); err != nil {
		return err
	}

	return nil
}

type P2Ps []*P2P

func NewP2Ps() *P2Ps {
	return &P2Ps{}
}

func (p2p *P2Ps) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p2p)
}

func (p2p *P2Ps) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p2p)
}

func (p2pl *P2Ps) GetNumPeers() error {
	// read data from db
	buf, err := scanData([]byte(statsUptimesPeers))
	if err != nil {
		return err
	}

	// unmarshal data to struct
	return p2pl.unmarshalData(buf)
}

func (p2pl *P2Ps) GetNumPeersByAddr(addr string) error {
	// read data from db
	buf, err := scanNestedData([]byte(statsUptimesPeersByAddr), []byte(addr))
	if err != nil {
		return err
	}

	// unmarshal data to struct
	return p2pl.unmarshalData(buf)
}

func (p2pl *P2Ps) GetNumPeersByAddrByRange(addr, min, max string) error {
	// validate times
	if err := validateTimes(min, max); err != nil {
		return err
	}

	// read data from db
	buf, err := scanNestedDataByRange([]byte(statsUptimesPeersByAddr), []byte(addr), []byte(min), []byte(max))
	if err != nil {
		return err
	}

	// unmarshal data to struct
	return p2pl.unmarshalData(buf)
}

func (p2pl *P2Ps) GetNumPeersByCluster(addrs string) error {
	// split addrs string
	addrl := splitAddrs(addrs)

	// read data from db
	var buf [][]byte
	for _, addr := range addrl {
		b, err := readData([]byte(statsUptimesPeers), []byte(addr))
		if err != nil {
			return err // consider skip return and read as many as possible
		}

		buf = append(buf, b)
	}

	// unmarshal data to struct
	return p2pl.unmarshalData(buf)
}

func (p2pl *P2Ps) unmarshalData(buf [][]byte) error {
	for _, b := range buf {
		p2p := NewP2P()
		if err := json.Unmarshal(b, p2p); err != nil {
			return err
		}

		*p2pl = append(*p2pl, p2p)
	}

	return nil
}
