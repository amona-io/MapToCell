package naverapi_test

import (
	"github.com/google/go-cmp/cmp"
	"handlegeo/naverapi"
	"handlegeo/utils"
	"testing"
)

func TestRequestAPINoData(t *testing.T) {
	coords := "0,0"
	actual := naverapi.RequestAPI(coords)
	statusCode := utils.GetStatusCode(actual)
	expected := 3
	if !cmp.Equal(statusCode, expected) {
		t.Errorf("%v != %v", statusCode, expected)
	}
}

func TestRequestAPISuccess(t *testing.T) {
	coords := "950000.00,1950000.00"
	actual := naverapi.RequestAPI(coords)
	statusCode := utils.GetStatusCode(actual)
	expected := 0
	if !cmp.Equal(statusCode, expected) {
		t.Errorf("%v != %v", statusCode, expected)
	}
}

