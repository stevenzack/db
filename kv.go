package db

import (
	"strconv"

	"github.com/StevenZack/livedata"
)

func (db *DB) String(k, def string) *livedata.LiveDataString {
	v := db.GetVar(k)
	if v == "" {
		v = def
	}
	return livedata.NewLiveDataString(v)
}

func (db *DB) Int(k string, def int) *livedata.LiveDataInt {
	is := db.GetVar(k)
	i := 0
	if is == "" {
		i = def
	} else {
		var e error
		i, e = strconv.Atoi(is)
		if e != nil {
			db.SetVar(k, strconv.Itoa(def))
			i = def
		}
	}
	return livedata.NewLiveDataInt(i)
}

func (db *DB) Bool(k string, def bool) *livedata.LiveDataBool {
	bs := db.GetVar(k)
	b := false
	if bs == "" {
		b = def
	} else {
		var e error
		b, e = strconv.ParseBool(bs)
		if e != nil {
			db.SetVar(k, strconv.FormatBool(def))
			b = def
		}
	}
	return livedata.NewLiveDataBool(b)
}
