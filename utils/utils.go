package utils

import (
	"encoding/json"
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalf("Return Err %s\n", err)
	}
}

func StringToJSON(res string) map[string]interface{} {
	resBody := res
	var data map[string]interface{}
	err := json.Unmarshal([]byte(resBody), &data)
	CheckErr(err)
	return data
}

func GetStatusCode(resBody string) int {
	jsonBody := StringToJSON(resBody)
	statusCode := jsonBody["status"].(map[string]interface{})["code"]
	floatStatusCode := statusCode.(float64)
	intStatusCode := int(floatStatusCode)
	return intStatusCode
}