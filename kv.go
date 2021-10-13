package db

import (
	"strconv"
	"time"

	"github.com/StevenZack/livedata"
)

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

func (db *DB) Int32(k string, def int32) *livedata.Int32 {
	is := db.GetVar(k)
	var i int32
	if is == "" {
		i = def
	} else {
		i2, e := strconv.ParseInt(is, 10, 64)
		if e != nil {
			db.SetVar(k, strconv.FormatInt(int64(def), 10))
			i = def
		} else {
			i = int32(i2)
		}
	}
	l := livedata.NewInt32(i)
	l.ObserveForever(func(v int32) {
		if strconv.FormatInt(int64(v), 10) == db.GetVar(k) {
			return
		}
		db.SetVar(k, strconv.FormatInt(int64(v), 10))
	})
	return l
}

func (db *DB) Int64(k string, def int64) *livedata.Int64 {
	is := db.GetVar(k)
	var i int64
	if is == "" {
		i = def
	} else {
		var e error
		i, e = strconv.ParseInt(is, 10, 64)
		if e != nil {
			db.SetVar(k, strconv.FormatInt(def, 10))
			i = def
		}
	}
	l := livedata.NewInt64(i)
	l.ObserveForever(func(v int64) {
		if strconv.FormatInt(v, 10) == db.GetVar(k) {
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
