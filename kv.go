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
	l := livedata.NewLiveDataString(v)
	l.ObserveForever(func(s string) {
		if db.GetVar(k) == s {
			return
		}
		db.SetVar(k, s)
	})
	return l
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
	l := livedata.NewLiveDataInt(i)
	l.ObserveForever(func(v int) {
		if strconv.Itoa(v) == db.GetVar(k) {
			return
		}
		db.SetVar(k, strconv.FormatInt(int64(v), 10))
	})
	return l
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
	l := livedata.NewLiveDataBool(b)
	l.ObserveForever(func(v bool) {
		if strconv.FormatBool(v) == db.GetVar(k) {
			return
		}
		db.SetVar(k, strconv.FormatBool(v))
	})
	return l
}
