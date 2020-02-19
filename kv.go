package db

import (
	"github.com/StevenZack/livedata"
)

func String(k, v string) *livedata.LiveDataString {
	return livedata.NewLiveDataString(v)
}

func Int(i int) *livedata.LiveDataInt {
	return livedata.NewLiveDataInt(i)
}

func Bool(b bool) *livedata.LiveDataBool {
	return livedata.NewLiveDataBool(b)
}
