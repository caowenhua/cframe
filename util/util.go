package util

import (
	"encoding/json"
	"time"
	"github.com/Centny/gwf/log"
)

func Timestamp(t time.Time) int64 {
	return t.Local().UnixNano() / 1e6
}

func Time(timestamp int64) time.Time {
	return time.Unix(0, timestamp*1e6)
}

func Now() int64 {
	return Timestamp(time.Now())
}

func ToJson(value interface{}) string {
	bs, _ := json.Marshal(value)
	return string(bs)
}

func JsonToValue(str string, value interface{}) error {
	log.D()
	return json.Unmarshal([]byte(str), value)
}
