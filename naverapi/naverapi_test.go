package naverapi_test

import (
	"fmt"
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

func TestRequestAPISuccessUtmK(t *testing.T) {
	coords := "950000.00,1950000.00"
	actual := naverapi.RequestAPI(coords)
	statusCode := utils.GetStatusCode(actual)
	expected := 0
	if !cmp.Equal(statusCode, expected) {
		t.Errorf("%v != %v", statusCode, expected)
	}
}

func TestRequestAPISuccessEpsg4326(t *testing.T) {
	coords := "126.6860049135999,37.36519043699966"
	actual := naverapi.RequestAPI(coords)
	fmt.Println(actual)
	statusCode := utils.GetStatusCode(actual)
	expected := 0
	if !cmp.Equal(statusCode, expected) {
		t.Errorf("%v != %v", statusCode, expected)
	}
}
