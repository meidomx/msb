package utils

import "encoding/json"

func ObjToJson(obj interface{}) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return b
}
