package jsonUtil

import (
	"encoding/json"
	"fmt"
)

func Encode (data interface{}) string {

	jsn, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Json Encoding Error: " + err.Error())
		return ""
	}

	jsonString := string(jsn)

	return jsonString
}


func Decode (jsonString string) (data interface{}) {

	err := json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		fmt.Println("Json Decoding Error: " + err.Error())
		return nil
	}

	return
}