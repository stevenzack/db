package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/StevenZack/tools/cryptoToolkit"

	"github.com/StevenZack/tools/fileToolkit"

	"github.com/StevenZack/tools/strToolkit"
)

type DB struct {
	dir    string
	cypher string
	log    bool
}

func NewDB(dir, cypher string) (*DB, error) {
	e := os.MkdirAll(dir, 0755)
	if e != nil {
		return nil, e
	}
	return &DB{
		dir:    strToolkit.Getrpath(dir),
		cypher: cypher,
	}, nil
}

func MustNewDB(dir, cypher string) *DB {
	db, e := NewDB(dir, cypher)
	if e != nil {
		log.Fatal(e)
	}
	return db
}

func (d *DB) SetLog(b bool) {
	d.log = b
}

func (d *DB) SetVar(k, v string) {
	if d.log {
		fmt.Println("setVar", k, v)
	}
	path := d.dir + k
	fileToolkit.TruncateFile(path)
	f, e := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if e != nil {
		fmt.Println("open file error :", e)
		return
	}
	f.WriteString(v)
	f.Close()
}

func (d *DB) GetVar(k string) string {
	path := d.dir + k
	f, e := os.OpenFile(path, os.O_RDONLY, 0644)
	if e != nil {
		return ""
	}
	defer f.Close()
	b, e := ioutil.ReadAll(f)
	if e != nil {
		fmt.Println("read file error :", e)
		return ""
	}
	return string(b)
}

func (d *DB) SetSecret(k, v string) {
	b, e := cryptoToolkit.EncryptAES([]byte(v), []byte(d.cypher))
	if e != nil {
		log.Println(e)
		return
	}
	d.SetVar(k, string(b))
}

func (d *DB) GetSecret(k string) string {
	enc := d.GetVar(k)
	if enc == "" {
		return ""
	}
	dec, e := cryptoToolkit.DecryptAES([]byte(enc), []byte(d.cypher))
	if e != nil {
		fmt.Println("decrypt str error :", e)
		d.SetVar(k, "")
		return ""
	}
	return string(dec)
}
