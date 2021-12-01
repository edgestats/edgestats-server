package data

import (
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func TestNewBoltDB(t *testing.T) {
	// test variables
	var tmp string
	var fp string

	// test new db
	tmp = t.TempDir()
	fp = filepath.Join(tmp, "test.db")

	got := NewBoltDB(fp)
	if got == nil {
		t.Fatalf("data.NewBoltDB() returned: %v", got)
	}
}

func TestBoltDBPath(t *testing.T) {
	// test variables
	tmp := t.TempDir()
	fp := filepath.Join(tmp, "test.db")
	db := NewBoltDB(fp)
	defer db.Close()

	got := db.Path()
	want := fp

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("BoltDBPath() returned: %v, wanted: %v", got, want)
	}
}

func TestBoltDBClose(t *testing.T) {}

func TestReadData(t *testing.T) {}

func TestReadLastData(t *testing.T) {}

func TestScanData(t *testing.T) {}

func TestWriteData(t *testing.T) {}

func TestScanNestedData(t *testing.T) {}

func TestWriteNestedData(t *testing.T) {}

func TestScanNestedDataByRange(t *testing.T) {}

func TestScanDataByRange(t *testing.T) {}

func TestSplitAddrs(t *testing.T) {
	l := "0xabc12345,0xdef67890,0xbcd45678,0xbcd45678,0xcdf28465"
	want := []string{"0xabc12345", "0xbcd45678", "0xcdf28465", "0xdef67890"}
	got := splitAddrs(l)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("data.splitAddrs() returned: %v, wanted: %v", got, want)
	}
}

func TestValidateTimes(t *testing.T) {
	var min string
	var max string
	var err error

	now := time.Now()

	// test no time error
	min = now.Format(time.RFC3339)
	max = now.Add(time.Hour * 2).Format(time.RFC3339)

	err = validateTimes(min, max)
	if err != nil {
		t.Fatalf("data.validateTimes() returned error: %v", err)
	}

	// test min time error
	min = now.Format(time.RFC1123)
	max = now.Add(time.Hour * 2).Format(time.RFC3339)

	err = validateTimes(min, max)
	if err == nil {
		t.Fatalf("data.validateTimes() returned: %v, wanted error", nil)
	}

	// test max time error
	min = now.Format(time.RFC3339)
	max = now.Add(time.Hour * 2).String()

	err = validateTimes(min, max)
	if err == nil {
		t.Fatalf("data.validateTimes() returned: %v, wanted error", nil)
	}
}
