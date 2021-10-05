package utils

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		//log.Fatalf("Return Err %s\n", err)
		log.Panicf("Return Err %s\n", err)
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

func SplitCoorsToFloat(coord string) (float64, float64) {
	coordsArray := strings.Split(coord, ",")
	latitudeStr, longitudeStr := coordsArray[0], coordsArray[1]
	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	CheckErr(err)
	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	CheckErr(err)
	return latitude, longitude
}
