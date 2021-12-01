package data

import (
	"bytes"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	bolt "go.etcd.io/bbolt"
)

const (
	statsUptimesBroadcasts      = "/stats/uptimes/broadcasts"
	statsUptimesBroadcatsByAddr = "/stats/uptimes/broadcasts/addrs"
	statsUptimesPeers           = "/stats/uptimes/peers"
	statsUptimesPeersByAddr     = "/stats/uptimes/peers/addrs"
	statsBlocks                 = "/stats/blocks"
)

var (
	dbPath = "./edgestats.db"
	DB     = NewBoltDB(dbPath) // autostart db // defer close in main
)

type BoltDB struct {
	db *bolt.DB
}

func NewBoltDB(fp string) *BoltDB {
	db, err := bolt.Open(fp, os.FileMode(0664), nil)
	if err != nil {
		log.Fatal(err)
	}

	return &BoltDB{db}
}

func (b *BoltDB) Path() string {
	return b.db.Path()
}

func (b *BoltDB) Close() {
	b.db.Close()
}

func readData(bkt []byte, key []byte) ([]byte, error) {
	var buf []byte

	db := DB.db
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bkt)
		buf = b.Get(key)
		return nil
	})

	return buf, err
}

func readLastData(bkt []byte) ([]byte, error) {
	var buf []byte

	db := DB.db
	// create bucket if not exists
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bkt)
		return err
	})
	if err != nil {
		return buf, err
	}

	// read data from db
	err = db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bkt).Cursor()
		_, buf = c.Last() // read last entry
		return nil
	})

	return buf, err
}

func scanData(bkt []byte) ([][]byte, error) {
	var buf [][]byte

	db := DB.db
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bkt).Cursor()

		// // scan in ascending order
		// for k, v := c.First(); k != nil; k, v = c.Next() {
		// 	// fmt.Printf("key: %s\nvalue: %s\n", k, v)
		// 	buf = append(buf, v)
		// }

		// scan in descending order
		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			buf = append(buf, v)
		}

		return nil
	})

	return buf, err
}

func writeData(bkt, key, val []byte) error {
	db := DB.db

	// create bucket if not exists
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bkt)
		return err
	})
	if err != nil {
		return err
	}

	// write data to db
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bkt)
		err := b.Put(key, val)
		return err
	})

	return err
}

func scanNestedData(bkt, nst []byte) ([][]byte, error) {
	var buf [][]byte

	db := DB.db
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bkt).Bucket(nst).Cursor()

		// // scan in ascending order
		// for k, v := c.First(); k != nil; k, v = c.Next() {
		//      // fmt.Printf("key: %s\nvalue: %s\n", k, v)
		//      buf = append(buf, v)
		// }

		// scan in descending order
		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			buf = append(buf, v)
		}

		return nil
	})

	return buf, err
}

func writeNestedData(bkt, nst, key, val []byte) error {
	db := DB.db

	// create nested bucket if not exists
	err := db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists(bkt)
		if err != nil {
			return err
		}

		_, err = root.CreateBucketIfNotExists(nst)
		return err
	})
	if err != nil {
		return err
	}

	// write data to db
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bkt).Bucket(nst)
		err := b.Put(key, val)
		return err
	})

	return err
}

func scanNestedDataByRange(bkt, nst, min, max []byte) ([][]byte, error) {
	var buf [][]byte

	if bytes.Equal(max, []byte("")) { // possible to read c.Last()
		max = []byte(time.Now().UTC().Format(time.RFC3339))
	}

	db := DB.db
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bkt).Bucket(nst).Cursor()

		// // scan in ascending order
		// for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() { // < if [min,max)
		//         buf = append(buf, v)
		// }

		// scan in descending order
		k, v := c.Seek(max)
		for k, v = c.Prev(); k != nil && bytes.Compare(k, min) >= 0; k, v = c.Prev() { // > if (min,max]
			buf = append(buf, v)
		}

		return nil
	})

	return buf, err
}

func scanDataByRange(bkt, min, max []byte) ([][]byte, error) {
	var buf [][]byte

	if bytes.Equal(max, []byte("")) { // possible to read c.Last()
		max = []byte(time.Now().UTC().Format(time.RFC3339))
	}

	db := DB.db
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bkt).Cursor()

		// // scan in ascending order
		// for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() { // < if [min,max)
		//         buf = append(buf, v)
		// }

		// scan in descending order
		k, v := c.Seek(max)
		for k, v = c.Prev(); k != nil && bytes.Compare(k, min) >= 0; k, v = c.Prev() { // > if (min,max]
			buf = append(buf, v)
		}

		return nil
	})

	return buf, err
}

func splitAddrs(addrs string) []string {
	// remove spaces
	addrs = strings.ReplaceAll(addrs, " ", "")

	// split addrs
	addrl := strings.Split(addrs, ",")

	// remove duplicates
	addrm := map[string]bool{}
	for _, v := range addrl {
		if addrm[v] {
			continue
		}
		addrm[v] = true
	}

	var l []string
	for k := range addrm {
		l = append(l, k)
	}

	// sort alphanumeric
	sort.Strings(l)

	return l
}

func validateTimes(min, max string) error {
	// parse min time
	_, err := time.Parse(time.RFC3339, min)
	if err != nil {
		return err
	}

	// parse max time
	_, err = time.Parse(time.RFC3339, max)
	if err != nil && max != "" {
		return err
	}

	return nil
}
