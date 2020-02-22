package db

import (
	"fmt"
	"strconv"

	"github.com/StevenZack/livedata"
)

func (db *DB) StringSecret(k, def string) *livedata.LiveDataString {
	v := db.GetSecret(k)
	if v == "" {
		v = def
	}
	l := livedata.NewLiveDataString(v)
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

func (db *DB) IntSecret(k string, def int) *livedata.LiveDataInt {
	is := db.GetSecret(k)
	i := 0
	if is == "" {
		i = def
	} else {
		var e error
		i, e = strconv.Atoi(is)
		if e != nil {
			db.SetSecret(k, strconv.Itoa(def))
			i = def
		}
	}
	l := livedata.NewLiveDataInt(i)
	l.ObserveForever(func(v int) {
		if strconv.Itoa(v) == db.GetSecret(k) {
			return
		}
		db.SetSecret(k, strconv.FormatInt(int64(v), 10))
	})
	return l
}

func (db *DB) BoolSecret(k string, def bool) *livedata.LiveDataBool {
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
	l := livedata.NewLiveDataBool(b)
	l.ObserveForever(func(v bool) {
		if strconv.FormatBool(v) == db.GetSecret(k) {
			return
		}
		db.SetSecret(k, strconv.FormatBool(v))
	})
	return l
}
