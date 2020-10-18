package common

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func MarshalJson(v interface{}) string {
	var json = jsoniter.ConfigFastest
	buf, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Json Marshall Error. : " + err.Error())
		panic(err)
	}
	return string(buf)
}
