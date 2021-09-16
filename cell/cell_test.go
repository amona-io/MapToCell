package cell_test

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"handlegeo/cell"
	"testing"
)

func TestNewCell(t *testing.T) {
	east := 950000.00
	north := 1950000.00

	actual, _ := cell.NewCell(east, north)
	expected := &cell.Cell{
		LeftTop: fmt.Sprintf("%.2f,%.2f",east-50.00,north+50.00),
		RightTop: fmt.Sprintf("%.2f,%.2f",east+50.00,north+50.00),
		LeftBottom: fmt.Sprintf("%.2f,%.2f",east-50.00,north-50.00),
		RightBottom: fmt.Sprintf("%.2f,%.2f",east+50.00,north-50.00),
		Center: fmt.Sprintf("%.2f,%.2f", east,north),
		IsInRange: true,
		CenterCity: "서울특별시 마포구 신수동",
	}
	if !cmp.Equal(actual, expected) {
		t.Errorf("%v != %v", *actual, *expected)
	}
}

/*
{0 949950.00,1950050.00 950050.00,1950050.00 949950.00,1949950.00 950050.00,1949950.00 950000.00,1950000.00 true 서울특별시 마포구 신수동 }
!=
{0 949950.00,1950050.00 950050.00,1950050.00 949950.00,1949950.00 950050.00,1949950.00 950000.00,1950000.00 true 서울특별시 마포구 신수동 }
*/