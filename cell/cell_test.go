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

const xMargin = 0.0011317962
const yMargin = 0.0008983153

func TestNewCell(t *testing.T) {
	east := 127.107096
	north := 37.631880

	actual, _ := cell.NewCell(east, north)
	expected := &cell.Cell{
		LeftTop:     fmt.Sprintf("%.18f,%.18f", east-(xMargin/2), north+(yMargin/2)),
		RightTop:    fmt.Sprintf("%.18f,%.18f", east+(xMargin/2), north+(yMargin/2)),
		LeftBottom:  fmt.Sprintf("%.18f,%.18f", east-(xMargin/2), north-(yMargin/2)),
		RightBottom: fmt.Sprintf("%.18f,%.18f", east+(xMargin/2), north-(yMargin/2)),
		CenterX:     east,
		CenterY:     north,
		IsInRange:   true,
		CenterCity:  "경기도 구리시 갈매동",
		CreatedAt:   time.Time{}, // Zero Value
		UpdatedAt:   time.Time{}, // Zero Value
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
	testCell, err := cell.NewCell(127.107096, 37.631880)
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
