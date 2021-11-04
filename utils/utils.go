package utils

import (
	"encoding/json"
	"fmt"
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
	XStr, YStr := coordsArray[0], coordsArray[1]
	x, err := strconv.ParseFloat(XStr, 64)
	CheckErr(err)
	y, err := strconv.ParseFloat(YStr, 64)
	CheckErr(err)
	return x, y
}

func JoinFloatToCoord(x, y float64) string {
	coord := fmt.Sprintf("%.2f,%.2f", x, y)
	return coord
}
