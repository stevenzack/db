package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/StevenZack/livedata"
)

func (db *DB) StringSecret(k, def string) *livedata.String {
	v := db.GetSecret(k)
	if v == "" {
		v = def
	}
	l := livedata.NewString(v)
	l.ObserveForever(func(s string) {
		secret := db.GetSecret(k)
		if db.log {
			fmt.Println("onchange ", secret, " -> ", s)
		}
		if secret == s {
			return
		}
		db.SetSecret(k, s)
	})
	return l
}

func (db *DB) Int64Secret(k string, def int64) *livedata.Int64 {
	is := db.GetSecret(k)
	var i int64
	if is == "" {
		i = def
	} else {
		var e error
		i, e = strconv.ParseInt(is, 10, 64)
		if e != nil {
			db.SetSecret(k, strconv.FormatInt(def, 10))
			i = def
		}
	}
	l := livedata.NewInt64(i)
	l.ObserveForever(func(v int64) {
		if strconv.FormatInt(v, 10) == db.GetSecret(k) {
			return
		}
		db.SetSecret(k, strconv.FormatInt(v, 10))
	})
	return l
}

func (db *DB) BoolSecret(k string, def bool) *livedata.Bool {
	bs := db.GetSecret(k)
	b := false
	if bs == "" {
		b = def
	} else {
		var e error
		b, e = strconv.ParseBool(bs)
		if e != nil {
			db.SetSecret(k, strconv.FormatBool(def))
			b = def
		}
	}
	l := livedata.NewBool(b)
	l.ObserveForever(func(v bool) {
		if strconv.FormatBool(v) == db.GetSecret(k) {
			return
		}
		db.SetSecret(k, strconv.FormatBool(v))
	})
	return l
}

func (db *DB) TimeSecret(k string, def time.Time) *livedata.Time {
	ts := db.GetSecret(k)
	t := time.Time{}
	if ts == "" {
		t = def
	} else {
		var e error
		t, e = time.Parse(time.RFC3339, ts)
		if e != nil {
			db.SetSecret(k, def.Format(time.RFC3339))
			t = def
		}
	}
	l := livedata.NewTime(t)
	l.ObserveForever(func(v time.Time) {
		if v.Format(time.RFC3339) == db.GetSecret(k) {
			return
		}
		db.SetSecret(k, v.Format(time.RFC3339))
	})
	return l
}
