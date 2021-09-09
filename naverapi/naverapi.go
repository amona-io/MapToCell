package naverapi

import (
	"encoding/json"
	"fmt"
	_const "handlegeo/const"
	"handlegeo/utils"
	"io"
	"io/ioutil"
	"net/http"
)

func queryURLWrapper(coords string) string {
	baseURL := fmt.Sprintf("https://naveropenapi.apigw.ntruss.com/map-reversegeocode/v2/gc?sourcecrs=nhn:2048&coords=%s&output=json", coords)	// UTM-K (동-서,북-남)
	//baseURL := fmt.Sprintf("https://naveropenapi.apigw.ntruss.com/map-reversegeocode/v2/gc?coords=%s&output=json", coords)
	return baseURL
}

func getResponse(queryURL string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", queryURL, nil)
	utils.CheckErr(err)

	req.Header.Add("X-NCP-APIGW-API-KEY-ID", _const.NCPKEYID)
	req.Header.Add("X-NCP-APIGW-API-KEY", _const.NCPKEY)
	res, err := client.Do(req)
	utils.CheckErr(err)
	return res
}

func resBodyToString(res io.ReadCloser) string {
	resBodyBytes, err := ioutil.ReadAll(res)
	utils.CheckErr(err)
	bodyString := string(resBodyBytes)
	return bodyString
}

func stringToJSON(res string) map[string]interface{} {
	resBody := res
	var data map[string]interface{}
	err := json.Unmarshal([]byte(resBody), &data)
	utils.CheckErr(err)
	return data
}

func RequestAPI(coords string) string {
	queryURL := queryURLWrapper(coords)
	res := getResponse(queryURL)
	resBody := resBodyToString(res.Body)
	return resBody
}

func GetStatusCode(resBody string) int {
	jsonBody := stringToJSON(resBody)
	statusCode := jsonBody["status"].(map[string]interface{})["code"]
	fmt.Println(statusCode)
	return 0
}