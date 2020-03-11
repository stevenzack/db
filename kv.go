package db

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/StevenZack/tools/strToolkit"

	"github.com/StevenZack/livedata"
)

func (db *DB) List(k string, def []interface{}) *livedata.List {
	vs := db.GetVar(k)
	l := []interface{}{}
	if vs == "" {
		l = def
	} else {
		e := json.Unmarshal([]byte(vs), &l)
		if e != nil {
			db.SetVar(k, strToolkit.JsonArray(def))
			l = def
		}
	}
	ld := livedata.NewList(l)
	ld.ObserveForever(func(list []interface{}) {
		db.SetVar(k, strToolkit.JsonArray(list))
	})
	return ld
}
func (db *DB) String(k, def string) *livedata.String {
	v := db.GetVar(k)
	if v == "" {
		v = def
	}
	l := livedata.NewString(v)
	l.ObserveForever(func(s string) {
		if db.GetVar(k) == s {
			return
		}
		db.SetVar(k, s)
	})
	return l
}

func (db *DB) Int(k string, def int) *livedata.Int {
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
	l := livedata.NewInt(i)
	l.ObserveForever(func(v int) {
		if strconv.Itoa(v) == db.GetVar(k) {
			return
		}
		db.SetVar(k, strconv.FormatInt(int64(v), 10))
	})
	return l
}

func (db *DB) Bool(k string, def bool) *livedata.Bool {
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
	l := livedata.NewBool(b)
	l.ObserveForever(func(v bool) {
		if strconv.FormatBool(v) == db.GetVar(k) {
			return
		}
		db.SetVar(k, strconv.FormatBool(v))
	})
	return l
}

func (db *DB) Time(k string, def time.Time) *livedata.Time {
	ts := db.GetVar(k)
	t := time.Time{}
	if ts == "" {
		t = def
	} else {
		var e error
		t, e = time.Parse(time.RFC3339, ts)
		if e != nil {
			db.SetVar(k, def.Format(time.RFC3339))
			t = def
		}
	}
	l := livedata.NewTime(t)
	l.ObserveForever(func(v time.Time) {
		if v.Format(time.RFC3339) == db.GetVar(k) {
			return
		}
		db.SetVar(k, v.Format(time.RFC3339))
	})
	return l
}
