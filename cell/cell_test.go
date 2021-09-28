package cell_test

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"handlegeo/cell"
	database "handlegeo/db"
	"testing"
	"time"
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
		CreatedAt: time.Time{},		// Zero Value
		UpdatedAt: time.Time{},		// Zero Value
	}
	if !cmp.Equal(actual, expected) {
		t.Errorf("%v != %v", *actual, *expected)
	}
}

func TestNewCellFail(t *testing.T) {
	asserts := assert.New(t)
	east := 1.00
	north := 00.00
	_, err := cell.NewCell(east, north)
	asserts.Error(err, "비정상적임에도 처리되었습니다.")
}

// test db 생성해서 거기서 테스트 하는걸로 ?
func TestCreateAndDeleteCell(t *testing.T) {
	asserts := assert.New(t)
	testCell, err := cell.NewCell(1099400.00, 1823700.00)
	asserts.NoError(err, "테스트용 셀 인스턴스 생성에 실패했습니다.")

	db, err := database.Conn()
	asserts.NoError(err, "데이터베이스 연결에 실패했습니다.")

	err = testCell.Create(db)
	asserts.NoError(err, "데이터베이스에 셀 인스턴스를 입력할 수 없습니다.")

	lastCell := &cell.DBCell{}
	db.Model(&cell.DBCell{}).Last(&lastCell)
	err = lastCell.Delete(db)
	asserts.NoError(err, "데이터베이스에서 셀 인스턴스를 삭제할 수 없습니다.")
}

func TestGetCellData(t *testing.T) {
	asserts := assert.New(t)
	db, _ := database.Conn()
	result := map[string]interface{}{}
	db.Model(&cell.DBCell{}).First(&result)
	asserts.NotEmpty(result, "셀데이터를 데이터베이스에서 가져올 수 없습니다.")
}

/*
{0 949950.00,1950050.00 950050.00,1950050.00 949950.00,1949950.00 950050.00,1949950.00 950000.00,1950000.00 true 서울특별시 마포구 신수동 }
!=
{0 949950.00,1950050.00 950050.00,1950050.00 949950.00,1949950.00 950050.00,1949950.00 950000.00,1950000.00 true 서울특별시 마포구 신수동 }
*/
