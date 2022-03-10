package naverapi

import (
	"fmt"
	_const "handlegeo/const"
	"handlegeo/utils"
	"io"
	"io/ioutil"
	"net/http"
)

// RequestAPI returns the Area Information data for the input coords as a JSON String
func RequestAPI(coords string) string {
	fmt.Println(coords)
	queryURL := queryURLWrapper(coords)
	res := getResponse(queryURL)
	resBody := resBodyToString(res.Body)
	defer res.Body.Close()
	return resBody
}

// queryURLWrapper retrieve coords as string type and it returns query url for communicate with Naver Maps API
// 군사보안 목적으로 법정동을 정상적으로 리턴하지 않는 경우 -> orders=admcode 파라미터를 추가하여 법정동을 요청하면 정상적으로 받아올 수 있음
func queryURLWrapper(coords string) string {
	//baseURL := fmt.Sprintf("https://naveropenapi.apigw.ntruss.com/map-reversegeocode/v2/gc?coords=%s&output=json", coords)
	baseURL := fmt.Sprintf("https://naveropenapi.apigw.ntruss.com/map-reversegeocode/v2/gc?orders=admcode&coords=%s&output=json", coords)
	return baseURL
}

// getResponse retrieve queryURL and this function interacts with the queryURL.
// And It returns Pointer of the http.Response
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

// resBodyToString retrieve http response body instance and It wrapping this http response body to "String" data
func resBodyToString(resBody io.ReadCloser) string {
	resBodyBytes, err := ioutil.ReadAll(resBody)
	utils.CheckErr(err)
	bodyString := string(resBodyBytes)
	return bodyString
}
