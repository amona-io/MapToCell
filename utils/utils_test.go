package utils_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"handlegeo/utils"
	"testing"
)

func TestStringToJSON(t *testing.T) {
	asserts := assert.New(t)
	targetJSON := "{\"area1\":{\"name\":\"서울특별시\",\"coords\":{\"center\":" +
		"{\"crs\":\"EPSG:4326\",\"x\":126.9783882,\"y\":37.5666103}},\"alias\":\"서울\"}}"
	expected := "map[area1:map[alias:서울 coords:map[center:map[crs:EPSG:4326 x:126.9783882 y:37.5666103]] name:서울특별시]]"
	actual := fmt.Sprintf("%v", utils.StringToJSON(targetJSON))
	asserts.Equal(expected, actual)
}

func TestStringToJSONFail(t *testing.T) {
	asserts := assert.New(t)
	targetJSON := "asdasdasd"
	asserts.Panicsf(func() {utils.StringToJSON(targetJSON) }, "유효하지 않은 JSON을 변환하는 함수는 실패해야합니다.")
}

func TestGetStatusCode(t *testing.T) {
	asserts := assert.New(t)
	expected := 0
	targetBody := "{\"status\":{\"code\":0,\"name\":\"ok\",\"message\":\"done\"}}"
	actual := utils.GetStatusCode(targetBody)
	asserts.Equal(expected, actual)
}

func TestGetStatusCodeFail(t *testing.T) {
	asserts := assert.New(t)
	expected := 0
	targetBody := "{\"status\":{\"code\":200,\"name\":\"ok\",\"message\":\"fail\"}}"
	actual := utils.GetStatusCode(targetBody)
	asserts.NotEqual(expected, actual)
}

func TestSplitCoorsToFloat(t *testing.T) {
	asserts := assert.New(t)
	coordString := "947800.00,1943400.00"
	actualLatitude, actualLongitude := utils.SplitCoorsToFloat(coordString)
	expectedLatitude := 947800.00
	expectedLongitude := 1943400.00
	expected := []float64 {expectedLatitude, expectedLongitude}
	actual := []float64 {actualLatitude, actualLongitude}
	asserts.Equal(expected, actual)
}